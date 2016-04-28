
package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)

var mysqlBase *MysqlBase
  
type MysqlBase struct {
	DB *sql.DB
}
  
func GetMysqlInstance() (*MysqlBase){
	if mysqlBase == nil ||  mysqlBase.DB == nil {
		fmt.Println("haha")
		db, err := sql.Open("mysql", "root:admin@/tmac")
		if err != nil {
			log.Fatalf("Open database error: %s\n", err)
		}
	    	mysqlBase = &MysqlBase{DB:db}
   	}
   	return mysqlBase;
}

func (this *MysqlBase)insert(inSql string, args ...interface{}) {
	
	stmt, err := this.DB.Prepare(inSql)
	defer stmt.Close()
 
	if err != nil {
		log.Println(err)
		return
	}
	stmt.Exec(args...)
} 
func (this *MysqlBase)GetDb()(*sql.DB){
	return this.DB;
}
/*func (this *MysqlBase)get_data(qSql string, args ...interface{}){
	rows, err := this.DB.Query(qSql, args...)
	if err != nil {
		log.Println(err)
	}
	//defer rows.Close()
	return rows
} 
func main() {
	mydb := GetMysqlInstance()
	inSql := "INSERT INTO dongtai_news(newsid, imageid, title, create_time) VALUES(?, ?, ?, ?)"
	mydb.insert(inSql, 1, 1, "text", "2016-02-13 00:00:00") 

	qSql := "select id,title from dongtai_news where id = ?" 
	db := mydb.getDb()
	rows, err := db.Query(qSql, 1)
	if err != nil {
		log.Println(err)
	}
 
	defer rows.Close()
	var id int
	var title string
	for rows.Next() {
		err := rows.Scan(&id, &title)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, title)
	}
 
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}	
}*/
