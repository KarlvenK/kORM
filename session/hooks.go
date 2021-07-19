package session

import (
	"github.com/KarlvenK/korm/log"
	"reflect"
)

//hooks constants
const (
	BeforeQuery  = "BeforeQuery"
	AfterQuery   = "AfterQuery"
	BeforeUpdate = "BeforeUpdate"
	AfterUpdate  = "AfterUpdate"
	BeforeDelete = "BeforeDelete"
	AfterDelete  = "AfterDelete"
	BeforeInsert = "BeforeInsert"
	AfterInsert  = "AfterInsert"
)

//CallMethod calls the registered hooks
func (s *Session) CallMethod(method string, value interface{}) {
	fm := reflect.ValueOf(s.RefTable().Model).MethodByName(method)
	if value != nil {
		fm = reflect.ValueOf(value).MethodByName(method)
	}

	param := []reflect.Value{reflect.ValueOf(s)}
	if fm.IsValid() {
		if v := fm.Call(param); len(v) > 0 {
			if err, ok := v[0].Interface().(error); ok {
				log.Error(err)
			}
		}
	}
	return
}

/*
钩子机制同样是通过反射来实现的，s.RefTable().Model 或 value
即当前会话正在操作的对象，使用 MethodByName 方法反射得到该对象的方法。

将 s *Session 作为入参调用。每一个钩子的入参类型均是 *Session
*/
