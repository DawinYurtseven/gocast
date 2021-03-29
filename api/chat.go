package api

import (
	"TUM-Live/dao"
	"TUM-Live/model"
	"TUM-Live/tools"
	"context"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var m *melody.Melody
var stats = map[string]int{}
var statsLock = sync.RWMutex{}

func configGinChatRouter(router gin.IRoutes) {
	router.GET("/:vidId/ws", ChatStream)
	router.GET("/:vidId/stats", ChatStats)
	if m == nil {
		log.Printf("creating melody")
		m = melody.New()
	}
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		ctx, _ := s.Get("ctx") // get gin context

		user, uErr := tools.GetUser(ctx.(*gin.Context))
		student, sErr := tools.GetStudent(ctx.(*gin.Context))
		if uErr != nil && sErr != nil {
			// not allowed to send message
			return

		}
		var uid string
		if uErr == nil {
			uid = strconv.Itoa(int(user.ID))
		} else if sErr == nil {
			uid = student.ID
		}
		vID, err := strconv.Atoi(ctx.(*gin.Context).Param("vidId"))
		if err != nil {
			return
		}
		dao.AddMessage(string(msg), uid, uint(vID))
		_ = m.BroadcastFilter(msg, func(q *melody.Session) bool { // filter broadcasting to same lecture.
			return q.Request.URL.Path == s.Request.URL.Path
		})
	})
}

func CollectStats() {
	log.Printf("Collecting stats\n")
	for sID, numWatchers := range stats {
		log.Printf("Collecting stats for stream %v, viewers:%v\n", sID, numWatchers)
		stat := model.Stat{
			Time:    time.Now(),
			Viewers: numWatchers,
		}
		if s, err := dao.GetStreamByID(context.Background(), sID); err == nil {
			s.Stats = append(s.Stats, stat)
			if err = dao.SaveStream(&s); err != nil {
				log.Printf("Error saving stats: %v\n", err)
			}
		}
	}
}

func ChatStats(context *gin.Context) {
	/*u, uErr := tools.GetUser(context)
	if uErr != nil || u.Role != 1 {
		context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"msg": "stats can currently only be retrieved by admins"})
		return
	}*/
	vidId := context.Param("vidId")
	viewers, found := stats[vidId]
	if !found {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"msg": "stream not found"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"viewers": viewers})
}

func ChatStream(c *gin.Context) {
	ctxMap := make(map[string]interface{}, 1)
	println(c.Param("vidId") + "joined")
	statsLock.Lock()
	stats[c.Param("vidId")] = stats[c.Param("vidId")] + 1
	statsLock.Unlock()
	m.HandleClose(func(session *melody.Session, i int, s string) error {
		statsLock.Lock()
		stats[c.Param("vidId")] = stats[c.Param("vidId")] - 1
		if stats[c.Param("vidID")] == 0 {
			delete(stats, "vidID")
		}
		statsLock.Unlock()
		return nil
	})
	ctxMap["ctx"] = c
	_ = m.HandleRequestWithKeys(c.Writer, c.Request, ctxMap)
}
