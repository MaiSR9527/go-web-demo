package util

import (
	"fmt"
	"reflect"
	"simple-iris-project/log"
)

func SetObjByJson(obj interface{}, data map[string]interface{}) error {
	for key, value := range data {

	}
}

func setField(obj interface{}, name string, value interface{}) error {
	structData := reflect.TypeOf(obj).Elem()

	fieldValue, err := structData.FieldByName(name)
	if !err {
		log.Log("error", "No such field: "+name)
		return fmt.Errorf("No such field %s", name)
	}
	t := fieldValue.Type

}
