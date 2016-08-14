package models
import (
	"fmt"
	_"strings"
        _"tmac/common"
        "strconv"
        "log" 
)

func GetMatchInfo(schoolid int, leagueid int) (cells []map[string]string) {
	mydb := GetMysqlInstance()
	db := mydb.GetDb()
	qSql := fmt.Sprintf("select teamid1,teamid2,match_time,match_address,status,matchid from `match` where  schoolid = ? and leagueid = ?")
	rows, err := db.Query(qSql, schoolid, leagueid)
	if err != nil {
		log.Println(err)
	}
	
	defer rows.Close()
	var teamid1 int
	var teamid2 int 
	var match_time string
	var match_address string
	var status int
	var matchid int
	cells = []map[string]string{}
	for rows.Next() {
		cell := map[string]string{}
		err := rows.Scan(&teamid1, &teamid2, &match_time, &match_address, &status, &matchid)
		if err != nil {
			log.Fatal(err)
		}
		cell["teamid1"] = strconv.Itoa(teamid1)
		cell["teamid2"] = strconv.Itoa(teamid2)
		cell["match_time"] = match_time
		fmt.Println(match_address)
		cell["match_address"] = match_address
		cell["status"] = strconv.Itoa(status)
		cell["leagueid"] = strconv.Itoa(leagueid)
		cell["schoolid"] = strconv.Itoa(schoolid)
		cell["matchid"] = strconv.Itoa(matchid)
		cells = append(cells, cell)
	}
	//if id == 0: 
	//	return "0";
	return cells	
}
