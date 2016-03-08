package controllers 
import (
	"fmt"
	"strings"
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

var imageRecommendDir string
var baseDir string
var imageServer string
var platformServer string

func init() {
	imageServer = "http://115.159.101.179:8080";
	platformServer = "http://115.159.101.179:10086"
	var err error
	baseDir, err = os.Getwd()
	fmt.Println(baseDir)
	if err != nil {
    		fmt.Println("get current directory error=%s\n", err)
	}
	imageRecommendDir = fmt.Sprintf("%s/resource/dongtai/images/recommend", baseDir)
}

func ReqRecommendCells(timeStamp string, c *DongtaiController) {
        curImages, err := lib.ListDir(imageRecommendDir, timeStamp);
	cells := []map[string]string{}
        for _, image := range curImages {
		cell := map[string]string{}
	        imageAddress := fmt.Sprintf("%s", strings.TrimPrefix(image, baseDir))
		filename := lib.GetFileName(image)
		fmt.Println(filename);
		cell["image_address"] = imageServer + imageAddress
		cell["image_content"] = platformServer + "/dongtai/recommend&cell&" + filename
		cell["ticket"] = filename
		cells = append(cells, cell)	
        }
        b, err := json.Marshal(cells)
        if err != nil {
                fmt.Println("error", err)
        }
	c.Ctx.WriteString(string(b));
}
