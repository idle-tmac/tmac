package controllers 
import (
	"fmt"
	"github.com/astaxie/beego"
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
var imageBase64RecommendDir string
var baseDir string
var imageServer string
var platformServer string

func init() {	
	serverip := beego.AppConfig.String("serverip")
	image_server_port := beego.AppConfig.String("image_server_port")
	platform_server_port := beego.AppConfig.String("platform_server_port")
	imageServer = "http://" + serverip + ":" + image_server_port;
	platformServer = "http://" + serverip + ":" + platform_server_port;
	var err error
	baseDir, err = os.Getwd()
	fmt.Println(baseDir)
	if err != nil {
    		fmt.Println("get current directory error=%s\n", err)
	}
	imageRecommendDir = fmt.Sprintf("%s/resource/dongtai/images/recommend", baseDir)
	imageBase64RecommendDir = fmt.Sprintf("%s/resource/dongtai/imagesbase64/recommend", baseDir)
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
func ReqRecommendCell(ticket string, c *DongtaiController) {

	c.Data["Website"] = "campus basketball"
	c.Data["Email"] = "dreamer4@dream.com"
	imagePath := imageRecommendDir + "/" + ticket + ".jpg"
	encodeStr := lib.EncodeImageBase64(imagePath);
	imageBase64Path := imageBase64RecommendDir + "/" + ticket + ".base64"
	lib.WriteFile(imageBase64Path, encodeStr)
	c.Data["Testimage"] = encodeStr
	c.TplName = ticket + ".tpl"
}
