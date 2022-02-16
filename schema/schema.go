//go对象(结构体)与数据库表的转换

package schema

import (
	"geeorm/dialect"
	"go/ast"
	"reflect"
)

//字段约束
type Field struct {
	Name string
	Type string
	Tag string
	// Length int
	// DecimalPoint int
	// IfNull int
	// Description string
}

//对象映射数据库表，包含被映射的对象、表名、表所有字段、所有字段名、以及字段名和字段对应的map
type Schema struct {
	Model interface{}
	Name string
	Fields []*Field
	FieldNames []string
	FieldMap map[string]*Field
}

//入参都为指针，所以reflect需要用Indirect获取指针指向的内容
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	//获取dest对象的类型
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model: dest,
		Name: modelType.Name(),
		FieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		//判断struct是否为空以及Name是否以大写字母为首(表示是否私有变量)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeof(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("geeorm"); ok {
				field.Tag = v
			}

			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.FieldMap[p.Name] = field
		}
	}
	return schema
}

func (schema *Schema) GetField(name string) *Field {
	return schema.FieldMap[name]
}