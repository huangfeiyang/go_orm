/*
与数据库交互
封装Exec、QueryRow等方法，便于统一日志采集以及session会话复用
*/
package session

import (
	"database/sql"
	"geeorm/log"
	"strings"
)

type Session struct {
	db *sql.DB
	sql strings.Builder
	sqlVars []interface{}
}

func New(db *sql.DB) *Session {
	return &Session{db: db}
}

//每次执行数据库操作后都将清空sql以及占位符对应的数据，保证同一个会话可以复用
func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

//为什么要这样暴露？？？
func (s *Session) DB() *sql.DB {
	return s.db
}

//整合sql以及占位符对应数据
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

//执行sql
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

//获取一条记录
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

//获取纪录列表
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}