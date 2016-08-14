package models
import (
	"fmt"
	"strings"
        _"tmac/common"
        "strconv"
        "log" 
)

func GetTicket(time string, module string) (ticket string){
	fmt.Println(time)
	fmt.Println(module)
	mydb := GetMysqlInstance()
	db := mydb.GetDb()
	qSql := fmt.Sprintf("select id, create_time from dongtai_news where type = ? and create_time >= ? limit 1")
	fmt.Println(qSql)
	rows, err := db.Query(qSql, module, time)
	if err != nil {
		log.Println(err)
	}
	
	defer rows.Close()
	var id int
	var create_time string
	for rows.Next() {
		err := rows.Scan(&id, &create_time)
		if err != nil {
			log.Fatal(err)
		}
	}
	if id == 0{
		return "0";
	}
	time1 := strings.Replace(create_time, "-", "", 2)
	time2 := strings.Replace(time1, ":", "", 2)
	time3 := strings.Replace(time2, " ", "", 1)
	ticket = time3 + "_" + strconv.Itoa(id) 
	return ticket;
}

func GetDongtaiInfo(nid int, module string, num int) (cells []map[string]string) {
	mydb := GetMysqlInstance()
	db := mydb.GetDb()
	qSql := fmt.Sprintf("select * from dongtai_news where type = ? and id > ? limit ?")
	rows, err := db.Query(qSql, module, nid, num)
	if err != nil {
		log.Println(err)
	}
	
	defer rows.Close()
	var id int
	var newslink string
        var imagename string
	var title string
	var Type string
	var create_time string
	for rows.Next() {
		cell := map[string]string{}
		err := rows.Scan(&id, &newslink, &imagename, &title, &Type, &create_time)
		if err != nil {
			log.Fatal(err)
		}
		cell["id"] = strconv.Itoa(id)
		cell["imagename"] = imagename
		cell["newslink"] = newslink
		cell["title"] = title
		cell["type"] = Type
		cell["create_time"] = create_time
		cells = append(cells, cell)
	}
	//if id == 0: 
	//	return "0";
	return cells	
}
