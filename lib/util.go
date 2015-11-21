package lib

import (
	"reflect"
	_ "fmt"

)
func Call(m map[string]interface{}, name string, params ... interface{}) (result []reflect.Value) {
    
    f := reflect.ValueOf(m[name])
    in := make([]reflect.Value, len(params))
    for k, param := range params {
        in[k] = reflect.ValueOf(param)
    }
    result = f.Call(in)
    return

}
