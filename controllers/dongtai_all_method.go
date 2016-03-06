package controllers 
import (
	"fmt"
        _"tmac/common"
        _"tmac/models"
        "tmac/lib"
        "encoding/json"
        _"strconv"
        _"github.com/astaxie/beego/cache"
        _"github.com/astaxie/beego/cache/redis" 
        _"log" 
        _"time"
	"os"
)

var imageAllDir string

func init() {
	var currentDir string
	currentDir, err := os.Getwd()
	fmt.Println(currentDir)
	if err != nil {
    		fmt.Println("get current directory error=%s\n", err)
	}
	imageAllDir = fmt.Sprintf("%s/resource/images", currentDir)
}

func ReqAllCells(timeStamp string, c *DongtaiController) {
        curImages, err := lib.ListDir(imageAllDir, timeStamp);
	cells := []map[string]string{}
        for index, image := range curImages {
		cell := map[string]string{}
            	fmt.Println("Index = ", index, "Value = ", image)
		cell["image_address"] = "http://115.159.101.179" + image
		cell["image_content"] = "http://115.159.101.179:10086/dongtai/all&cell&20160203081509_1"
	        cells = append(cells, cell)	
        }
        b, err := json.Marshal(cells)
        if err != nil {
                fmt.Println("error", err)
        }
	c.Ctx.WriteString(string(b));
}
