package dialect

import "reflect"

/*
不同数据库支持的数据类型也是有差异的，即使功能相同，在 SQL 语句的表达上也可能有差异。
ORM 框架往往需要兼容多种数据库，因此我们需要将差异的这一部分提取出来，每一种数据库分别
实现，实现最大程度的复用和解耦。这部分代码称之为 dialect。
*/

var dialectsMap = map[string]Dialect{}

//Dialect is an interface contains methods that a dialect has to implement
//DataTypeOf 将go的类型转换为该数据库的类型
//TableExistSQL 返回某个表是否存在的 SQL语句，参数为表名
type Dialect interface {
	DataTypeOf(typ reflect.Value) string
	TableExistSQL(tableName string) (string, []interface{})
}

//RegisterDialect registers a dialect to the global variable
func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

//GetDialect gets the dialect from global variable if it exists
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
