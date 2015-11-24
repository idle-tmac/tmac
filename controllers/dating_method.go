package controllers 
import (
	    "fmt"
        "tmac/common"
        "tmac/models"
        "tmac/lib"
        "encoding/json"
        _"strconv"
        _"github.com/astaxie/beego/cache"
        _"github.com/astaxie/beego/cache/redis" 
        _"log" 
        "time"
)


//dating function register
var datingFuncs = map[string]interface{} {"push":PushDating}

func PushDating(msgPush string, c *DatingController) {
      
      /*mes := &common.DatingPushMsg{Datingid:int64(5), Uid:int64(123), Text:"hello"}
      msgPush1, _ := json.Marshal(mes); 
      msgPush = string(msgPush1)
	  fmt.Println(string(msgPush));*/
      
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
      redis.Put("lhy",inputJson,10);
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
