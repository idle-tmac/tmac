package lib

import (
	"reflect"
	_"fmt"
	"tmac/common"
	"strconv" 
	"io/ioutil"
	"os"
	"strings"
	"path"
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

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func ListDir(dirPth string, prefix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	prefix = strings.ToUpper(prefix) //忽略后缀匹配的大小写 
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
        	if strings.HasPrefix(strings.ToUpper(fi.Name()), prefix) { //匹配文件
            		files = append(files, dirPth+PthSep+fi.Name())
        	}
    	}
    	return files, nil
}
func GetFileName(filePath string) (filename string) {
    	filenameWithSuffix := path.Base(filePath)
    	fileSuffix := path.Ext(filenameWithSuffix)
    	filename = strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	return
}
/*func EncodeImageBase64(imagePath string) (encodeStr string) {
	f, err := os.Open(imagePath)
 	fbytes := ioutil.ReadAll(f)
	base64.StdEncoding.Encode(base64Bytes, fbytes) //buff转成base64
	return string(base64Bytes)
}*/
