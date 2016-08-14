package models
import (
	"fmt"
	_ "strings"
        _"tmac/common"
        "strconv"
        "log" 
)

func GetTeamInfo(teamid int) (cells map[string]string) {
	mydb := GetMysqlInstance()
	db := mydb.GetDb()
	qSql := fmt.Sprintf("select teamid, logoid, name from team where  teamid = ?")
	rows, err := db.Query(qSql, teamid)
	if err != nil {
		log.Println(err)
	}
        	
   	var logoid int
  	var name string
	cell := map[string]string{}
	for rows.Next() {
		err := rows.Scan(&teamid, &logoid, &name)
		if err != nil {
			log.Fatal(err)
		}
		cell["team_id"] = strconv.Itoa(teamid)
		cell["logoid"] = strconv.Itoa(logoid)
		cell["team_name"] = name
	}
	//if id == 0: 
	//	return "0";
	return cell	
}
