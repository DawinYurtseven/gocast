// Code generated by MockGen. DO NOT EDIT.
// Source: statistics.go

// Package mock_dao is a generated GoMock package.
package mock_dao

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dao "github.com/TUM-Dev/gocast/dao"
	model "github.com/TUM-Dev/gocast/model"
)

// MockStatisticsDao is a mock of StatisticsDao interface.
type MockStatisticsDao struct {
	ctrl     *gomock.Controller
	recorder *MockStatisticsDaoMockRecorder
}

// MockStatisticsDaoMockRecorder is the mock recorder for MockStatisticsDao.
type MockStatisticsDaoMockRecorder struct {
	mock *MockStatisticsDao
}

// NewMockStatisticsDao creates a new mock instance.
func NewMockStatisticsDao(ctrl *gomock.Controller) *MockStatisticsDao {
	mock := &MockStatisticsDao{ctrl: ctrl}
	mock.recorder = &MockStatisticsDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatisticsDao) EXPECT() *MockStatisticsDaoMockRecorder {
	return m.recorder
}

// AddStat mocks base method.
func (m *MockStatisticsDao) AddStat(stat model.Stat) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddStat", stat)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddStat indicates an expected call of AddStat.
func (mr *MockStatisticsDaoMockRecorder) AddStat(stat interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddStat", reflect.TypeOf((*MockStatisticsDao)(nil).AddStat), stat)
}

// GetCourseNumLiveViews mocks base method.
func (m *MockStatisticsDao) GetCourseNumLiveViews(courseID uint) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourseNumLiveViews", courseID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourseNumLiveViews indicates an expected call of GetCourseNumLiveViews.
func (mr *MockStatisticsDaoMockRecorder) GetCourseNumLiveViews(courseID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourseNumLiveViews", reflect.TypeOf((*MockStatisticsDao)(nil).GetCourseNumLiveViews), courseID)
}

// GetCourseNumStudents mocks base method.
func (m *MockStatisticsDao) GetCourseNumStudents(courseID uint) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourseNumStudents", courseID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourseNumStudents indicates an expected call of GetCourseNumStudents.
func (mr *MockStatisticsDaoMockRecorder) GetCourseNumStudents(courseID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourseNumStudents", reflect.TypeOf((*MockStatisticsDao)(nil).GetCourseNumStudents), courseID)
}

// GetCourseNumVodViews mocks base method.
func (m *MockStatisticsDao) GetCourseNumVodViews(courseID uint) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourseNumVodViews", courseID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourseNumVodViews indicates an expected call of GetCourseNumVodViews.
func (mr *MockStatisticsDaoMockRecorder) GetCourseNumVodViews(courseID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourseNumVodViews", reflect.TypeOf((*MockStatisticsDao)(nil).GetCourseNumVodViews), courseID)
}

// GetCourseNumVodViewsPerDay mocks base method.
func (m *MockStatisticsDao) GetCourseNumVodViewsPerDay(courseID uint) ([]dao.Stat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourseNumVodViewsPerDay", courseID)
	ret0, _ := ret[0].([]dao.Stat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourseNumVodViewsPerDay indicates an expected call of GetCourseNumVodViewsPerDay.
func (mr *MockStatisticsDaoMockRecorder) GetCourseNumVodViewsPerDay(courseID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourseNumVodViewsPerDay", reflect.TypeOf((*MockStatisticsDao)(nil).GetCourseNumVodViewsPerDay), courseID)
}

// GetCourseStatsHourly mocks base method.
func (m *MockStatisticsDao) GetCourseStatsHourly(courseID uint) ([]dao.Stat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourseStatsHourly", courseID)
	ret0, _ := ret[0].([]dao.Stat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourseStatsHourly indicates an expected call of GetCourseStatsHourly.
func (mr *MockStatisticsDaoMockRecorder) GetCourseStatsHourly(courseID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourseStatsHourly", reflect.TypeOf((*MockStatisticsDao)(nil).GetCourseStatsHourly), courseID)
}

// GetCourseStatsWeekdays mocks base method.
func (m *MockStatisticsDao) GetCourseStatsWeekdays(courseID uint) ([]dao.Stat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourseStatsWeekdays", courseID)
	ret0, _ := ret[0].([]dao.Stat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourseStatsWeekdays indicates an expected call of GetCourseStatsWeekdays.
func (mr *MockStatisticsDaoMockRecorder) GetCourseStatsWeekdays(courseID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourseStatsWeekdays", reflect.TypeOf((*MockStatisticsDao)(nil).GetCourseStatsWeekdays), courseID)
}

// GetStreamNumLiveViews mocks base method.
func (m *MockStatisticsDao) GetStreamNumLiveViews(streamID uint) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStreamNumLiveViews", streamID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStreamNumLiveViews indicates an expected call of GetStreamNumLiveViews.
func (mr *MockStatisticsDaoMockRecorder) GetStreamNumLiveViews(streamID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStreamNumLiveViews", reflect.TypeOf((*MockStatisticsDao)(nil).GetStreamNumLiveViews), streamID)
}

// GetStudentActivityCourseStats mocks base method.
func (m *MockStatisticsDao) GetStudentActivityCourseStats(courseID uint, live bool) ([]dao.Stat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudentActivityCourseStats", courseID, live)
	ret0, _ := ret[0].([]dao.Stat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudentActivityCourseStats indicates an expected call of GetStudentActivityCourseStats.
func (mr *MockStatisticsDaoMockRecorder) GetStudentActivityCourseStats(courseID, live interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudentActivityCourseStats", reflect.TypeOf((*MockStatisticsDao)(nil).GetStudentActivityCourseStats), courseID, live)
}
