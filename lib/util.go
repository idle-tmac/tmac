package lib

import (
	"reflect"
	"fmt"
	"io"
	"tmac/common"
	"strconv" 
	"io/ioutil"
	"os"
	"strings"
	"path"
	"encoding/base64"
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


//file and directory operate
func ListDir(dirPth string, prefix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	prefix = strings.ToUpper(prefix) 
	for _, fi := range dir {
		if fi.IsDir() { 
			continue
		}
        	if strings.HasPrefix(strings.ToUpper(fi.Name()), prefix) { 
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

func ReadFile(filePath string)([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}	
	return ioutil.ReadAll(f)
}
func WriteFile(filePath string, text string) {
	f, err := os.Create(filePath)
	defer f.Close()
	check(err)
	n, err1 := io.WriteString(f, text)
	check(err1)
	fmt.Println("write %d bytes", n)
}


//suanfa fuction 
func GetNextImages(imageRecommendDir string, ticket string) (files []string, err error){ 

	files, err = ListDir(imageRecommendDir, "1234567");
    	return files, nil
}
//other
func EncodeImageBase64(imagePath string) (encodeStr string) {
	ff, _ := ioutil.ReadFile(imagePath)
	fbytes := base64.StdEncoding.EncodeToString(ff)
	return string(fbytes)
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}



