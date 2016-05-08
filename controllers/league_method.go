package controllers 
import (
	"fmt"
	"github.com/astaxie/beego"
	_"strings"
        _"tmac/common"
        "tmac/models"
        _"tmac/lib"
        _"encoding/json"
        _"strconv"
        _"github.com/astaxie/beego/cache"
        _"github.com/astaxie/beego/cache/redis" 
        "log" 
        "encoding/json" 
        "strconv" 
        _"time"
	"os"
)
//var imageServer string
//var platformServer string
var leagueLogoDir string
var leaguePosterDir string
func init() {	
	//serverip := beego.AppConfig.String("serverip")
	//image_server_port := beego.AppConfig.String("image_server_port")
	//platform_server_port := beego.AppConfig.String("httpport")

	
	//imageServer = "http://" + serverip + ":" + image_server_port;
	//platformServer = "http://" + serverip + ":" + platform_server_port;
	var err error
	baseDir, err = os.Getwd()
	fmt.Println(baseDir)
	if err != nil {
    		fmt.Println("get current directory error=%s\n", err)
	}

	leagueLogoDir = fmt.Sprintf("%s/%s", baseDir, beego.AppConfig.String("leagueLogoDir"))
	leaguePosterDir = fmt.Sprintf("%s/%s", baseDir, beego.AppConfig.String("leaguePosterDir"))
}

func (this *LeagueController) Prepare() {
}

func (c *LeagueController) Test() {
	module := c.Ctx.Input.Param(":module")
        flag := c.Ctx.Input.Param(":bbb")
     	fmt.Println(module)
     	fmt.Println(flag)
	
}
func (c *LeagueController) ReqinSchoolLeagueCells(){
	schoolid := c.Ctx.Input.Param(":schoolid")
        num := c.Ctx.Input.Param(":num")
       	ticket := c.Ctx.Input.Param(":ticket")
     	fmt.Println(num)
     	fmt.Println(schoolid)
     	fmt.Println(ticket)
     	fmt.Println("----------\n")
	mydb := models.GetMysqlInstance()
	db := mydb.GetDb()
	qSql := "select * from league where schoolid = ? and id > ? limit " + num
	rows, err := db.Query(qSql, schoolid, ticket)
	if err != nil {
		log.Println(err)
	}
 
	defer rows.Close()
	var id int
	var leagueid int
	var logoid int
	var posterid int
	var name string
	var team_num int
	var team_fans int
	var league_type string
	var create_time string
	var start_time string
	var end_time string
	cells := []map[string]string{}
	for rows.Next() {
	 	cell := map[string]string{}
		err := rows.Scan(&id, &schoolid, &leagueid,  &logoid, &posterid, &name, &team_num, &team_fans, &league_type, &create_time, &start_time, &end_time)
		if err != nil {
			log.Fatal(err)
		}
		
		cell["league_id"] =  strconv.Itoa(leagueid)
		cell["league_logo_address"] = imageServer + "/" + beego.AppConfig.String("leagueLogoDir") + "/" + strconv.Itoa(logoid) + ".jpg"
		cell["league_poster_address"] = imageServer + "/" + beego.AppConfig.String("leaguePosterDir") + "/" + strconv.Itoa(posterid) + ".jpg"
		cell["league_name"] =  name 
		cell["league_team_num"] = strconv.Itoa(team_num)
		cell["league_type"] = league_type
		cell["league_start_date"] = start_time
		cell["league_end_date"] = end_time
		cell["ticket"] = strconv.Itoa(id)
		cells = append(cells, cell)	
		//fmt.Println(id, schoolid, leagueid)
	}
        
	b, err := json.Marshal(cells)
        if err != nil {
                fmt.Println("error", err)
        }

	c.Ctx.WriteString(string(b));
}

func (c *LeagueController) ReqinSchoolLeagueCell(){
	schoolid := c.Ctx.Input.Param(":schoolid")
       	leagueid := c.Ctx.Input.Param(":leagueid")
     	fmt.Println(schoolid)
	mydb := models.GetMysqlInstance()
	db := mydb.GetDb()
	qSql := "select team_fans from league where schoolid = ? and leagueid = ?"
	rows, err := db.Query(qSql, schoolid, leagueid)
	if err != nil {
		log.Println(err)
	}
 
	defer rows.Close()
	var team_fans int
	cell := map[string]string{}
	//cells := []map[string]string{}
	for rows.Next() {
		err := rows.Scan(&team_fans)
		if err != nil {
			log.Fatal(err)
		}
		
		cell["league_fans_num"]	= strconv.Itoa(team_fans)
	}
        
	b, err := json.Marshal(cell)
        if err != nil {
                fmt.Println("error", err)
        }

	c.Ctx.WriteString(string(b));
}
