package controllers 
import (
	"fmt"
	_"github.com/astaxie/beego"
	_"strings"
        "tmac/common"
        "tmac/models"
        "tmac/lib"
        _"encoding/json"
        "strconv"
        _"github.com/astaxie/beego/cache"
        _"github.com/astaxie/beego/cache/redis" 
        "log" 
        "encoding/json" 
        _"time"
)
func init() {	
}

func (c *LoginController) LoginCheck(){
	phone := c.Ctx.Input.Param(":phone")
        passwd := c.Ctx.Input.Param(":passwd")
     	fmt.Println(phone)
     	fmt.Println(passwd)

	mydb := models.GetMysqlInstance()
	db := mydb.GetDb()
	qSql := "select passwd,schoolno from user where phone = ?"
	rows, err := db.Query(qSql, phone)
	if err != nil {
		log.Println(err)
	}
 
	defer rows.Close()
	var passwdb string
	var schoolno string
	var isUser bool
	isUser = false
	for rows.Next() {
		isUser = true
		err := rows.Scan(&passwdb, &schoolno)
		if err != nil {
			log.Fatal(err)
		}
	}
       
	cell := map[string]string{}
	cell["type"] = strconv.Itoa(common.USERNOEXIST)
	if isUser {
		cell["type"] = strconv.Itoa(common.USERWRONGPASSWD)
		if (lib.Base64Encode(passwd) == lib.Base64Encode(passwdb)) {
			cell["schoolid"] = schoolno
			//judge has other info or cell["type"] = strconv.Itoa(common.USERLESSINFO)
			cell["type"] = strconv.Itoa(common.USERPERFECT)
		}
	}
	
	b, err := json.Marshal(cell)
        if err != nil {
                fmt.Println("error", err)
        }

	c.Ctx.WriteString(string(b));
}

