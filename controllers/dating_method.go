package controllers 
import (
	"fmt"
        _"tmac/common"
        _"tmac/lib"
        _"encoding/json"
        "strconv"
        _"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego"
        _"github.com/astaxie/beego/cache/redis" 
        _"log" 
        _"time"
)


//dating function register
var datingFuncs = map[string]interface{} {"push":PushDating}

func PushDating(msgPush string, flag string, c *DatingController) {
      
      fmt.Println(msgPush);
      fmt.Println(flag);
      /*mes := &common.DatingPushMsg{Datingid:int64(5), Uid:int64(123), Text:"hello"}
      msgPush1, _ := json.Marshal(mes); 
      msgPush = string(msgPush1)
	  fmt.Println(string(msgPush));
      
      var dtPushMsg common.DatingPushMsg;
      err := json.Unmarshal([]byte(msgPush), &dtPushMsg)
      if err != nil {
          return
      }
      fmt.Println(dtPushMsg);

      timestamp := time.Now().Unix()
      key := lib.ComputeDtPushKey(dtPushMsg.Datingid, timestamp)
      msg := &common.DatingMsg{Uid: int64(dtPushMsg.Uid), Time:timestamp, Text: dtPushMsg.Text}
      item := map[string] *common.DatingMsg {
           key:msg,
      }

      inputJson, err1 := json.Marshal(item);
      if err1 != nil {
           return
      }

      redis := models.GetRedisInstance()
      redis.Put("lhy",inputJson,10);*/
      inputJson := msgPush
      c.Ctx.WriteString(string(inputJson));

}
func PullDating(msgPUsh string, c *DatingController) {
	
    /*xxx := redis.Get("lhy");
      temp := fmt.Sprintf("%s\n",xxx)
      value := map[string]*common.DatingMsg{};
      err = json.Unmarshal([]byte(temp), &value)
      if err != nil {
          fmt.Println("Invalid request format:", temp)
      }
      fmt.Println(value["0111"].Uid)
      c.Ctx.WriteString(string(inputJson));
    */
}

func (c *DatingController) Test(){
	module := c.Ctx.Input.Param(":module")
	page, _ := c.GetInt("page")
	dd, _ := c.GetInt("dd")
	beego.Debug("Page", page)
     	fmt.Println(module)
     	fmt.Println(page)
	b := module + ":" + strconv.Itoa(page) + strconv.Itoa(dd)
      	c.Ctx.WriteString(b);
	
}
