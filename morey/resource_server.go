package main
import (
	_"os"
	"fmt"
	"net/http"
	_"path/filepath"
)
func main() {
    http.HandleFunc("/resource/dongtai/images/recommend/", GetResourceImage)
    http.ListenAndServe(":8080", nil)
}
 
func GetResourceImage(w http.ResponseWriter, r *http.Request) {
    path1 := r.URL.String()
    fmt.Println(path1);
    path := "/home/tmac/project/src/tmac" + path1;
    //resource/images/recommend/1234567_1.jpg"
    	fmt.Println(path);
	http.ServeFile(w, r, path)
}
