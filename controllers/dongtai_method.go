package controllers 
import (
	"fmt"
        _"tmac/common"
        _"tmac/models"
        _"tmac/lib"
        "encoding/json"
        _"strconv"
        _"github.com/astaxie/beego/cache"
        _"github.com/astaxie/beego/cache/redis" 
        _"log" 
        _"time"
)


//dongtai all function register
var dongtaiFuncs = map[string]interface{} {"all_cells":ReqAllCells}

func ReqAllCells(timeStamp string, c *DongtaiController) {
	homeInfo := map[string]string{} 
        homeInfo["all"] = timeStamp
        b, err := json.Marshal(homeInfo)
        if err != nil {
                fmt.Println("error", err)
        }
	c.Ctx.WriteString(string(b));
}
