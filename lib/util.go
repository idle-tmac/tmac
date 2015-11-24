package lib

import (
	"reflect"
	_"fmt"
    "tmac/common"
    "strconv" 
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

func ComputeDtPushKey(dtid int64,timestamp int64) (key string) {
	datingid := strconv.FormatInt(dtid,10)
    timestampStr := strconv.FormatInt(timestamp,10)
    return strconv.Itoa(common.DATINGTALK) + datingid + timestampStr
}


