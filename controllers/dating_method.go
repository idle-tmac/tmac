package controllers 
import (
	_"fmt"
)


//dating function register
var datingFuncs = map[string]interface{} {"push":PushDating}

func PushDating(msg string, c *DatingController) {
      
      c.Ctx.WriteString("1");	
}
