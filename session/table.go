package session

import (
	"fmt"
	"geeorm/log"
	"geeorm/schema"
	"reflect"
	"strings"
)

//go对象映射为数据库表
func (s *Session) Model(value interface{}) *Session {
	//会话中的表为空或者传入的对象名与表名冲突
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

//返回会话中的表
func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Error("Model is not set")
	}
	return s.refTable
}

//???是否只返回error
//创建表
func (s *Session) CreateTable() (bool, error) {
	table := s.RefTable()
	var columns []string
	for _, field := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}

	desc := strings.Join(columns, ",")
	if _, err := s.Raw(fmt.Sprintf("CREATE TABLE %s (%s);", s.refTable.Name, desc)).Exec(); err != nil {
		return false, err
	}
	
	return true, nil
}

//删除表
func (s *Session) DropTable() (bool, error) {
	if _, err := s.Raw(fmt.Sprintf("DROP TABLE IF EXISTS %s;", s.refTable.Name)).Exec(); err != nil {
		return false, err
	}

	return true, nil
}

//查询是否存在表
func (s *Session) HasTable() bool {
	sql, values := s.dialect.TableExistSQL(s.refTable.Name)
	row := s.Raw(sql, values...).QueryRow()
	var tmp string
	_ = row.Scan(&tmp)
	return tmp == s.refTable.Name
}