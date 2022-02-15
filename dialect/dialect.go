//使用dialect隔离不同类型的数据库，便于扩展
//使用reflect获取任意struct对象的名称和字段，映射成数据库中的表
//数据库表创建、删除

package dialect

import "reflect"

var dialectsMap = map[string]Dialect{}

type Dialect interface {
	DataTypeof(typ reflect.Value) string//将go语言的类型转换为该数据库的数据类型
	TableExistSQL(tableName string) (string, []interface{})//返回某个表是否存在的sql语句，参数是表名
}

//注册不同数据库
func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

//获取不同数据库
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}