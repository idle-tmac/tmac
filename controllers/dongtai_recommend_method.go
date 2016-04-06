package controllers 
import (
	"fmt"
	"github.com/astaxie/beego"
	_"strings"
        _"tmac/common"
        "tmac/models"
        "tmac/lib"
        "encoding/json"
        _"strconv"
        _"github.com/astaxie/beego/cache"
        _"github.com/astaxie/beego/cache/redis" 
        _"log" 
        _"time"
	"os"
)

var imageDir string
var imageBase64Dir string
var baseDir string
var imageServer string
var platformServer string

func init() {	
	serverip := beego.AppConfig.String("serverip")
	image_server_port := beego.AppConfig.String("image_server_port")
	platform_server_port := beego.AppConfig.String("httpport")

	
	imageServer = "http://" + serverip + ":" + image_server_port;
	platformServer = "http://" + serverip + ":" + platform_server_port;
	var err error
	baseDir, err = os.Getwd()
	fmt.Println(baseDir)
	if err != nil {
    		fmt.Println("get current directory error=%s\n", err)
	}

	imageDir = fmt.Sprintf("%s/%s", baseDir, beego.AppConfig.String("imageDir"))
	imageBase64Dir = fmt.Sprintf("%s/%s", baseDir, beego.AppConfig.String("base64Dir"))
}

func (this *DongtaiController) Prepare() {
}

func (c *DongtaiController) ReqRecommendCells(){
	var ticket string
	module := c.Ctx.Input.Param(":module")
        flag := c.Ctx.Input.Param(":flag")
        num := c.Ctx.Input.Param(":num")
        ticket = c.Ctx.Input.Param(":ticket")
     	fmt.Println(module)
     	fmt.Println(num)
     	fmt.Println(flag)
     	fmt.Println(ticket)
	
	if ticket == "0" && flag == "0" {
		//prefix := lib.GetTime(-60 * 30);
		//imagePath := imageDir + "/" + module
		//prefixImages, _ := lib.ListFile(imagePath,prefix);
		//ticket = prefixImages[0]
		ticket = "1234567_1"
	}
		
	curImages, err := lib.GetNextImages(models.Module_images[module], ticket, flag, num);
	cells := []map[string]string{}
	var filename string
        for _, image := range curImages {
		cell := map[string]string{}
		filename = lib.GetFileName(image)
		cell["image_address"] = imageServer + "/" + beego.AppConfig.String("imageDir") + "/" + module + "/" + image
		cell["image_content"] = platformServer + "/dongtai/cell/" + module + "&" + image
		cell["ticket"] = filename
		cells = append(cells, cell)	
        }

        b, err := json.Marshal(cells)
        if err != nil {
                fmt.Println("error", err)
        }

	c.Ctx.WriteString(string(b));
}
func (c *DongtaiController) ReqRecommendCell() {
	module := c.Ctx.Input.Param(":module")
        imagename := c.Ctx.Input.Param(":imagename")

     	//fmt.Println(module)
	
	c.Data["Website"] = "campus basketball"
	c.Data["Email"] = "dreamer4@dream.com"
	imagePath := imageDir + "/" + module + "/" + imagename
	//fmt.Println(imagePath)
	filename := lib.GetFileName(imagename)
	encodeStr := lib.EncodeImageBase64(imagePath);
	imageBase64Path := imageBase64Dir + "/" + module + "/" + filename + ".base64"
	
	//fmt.Println(imageBase64Path)
	lib.WriteFile(imageBase64Path, encodeStr)
	//encodeStr := string(lib.ReadFile(imageBase64Path))
	c.Data["Testimage"] = encodeStr
	c.TplName = "1234567_1.tpl"
}
