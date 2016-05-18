
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

func (this *MysqlBase)Insert(inSql string, args ...interface{}) bool{
	
	stmt, err := this.DB.Prepare(inSql)
	defer stmt.Close()
 
	if err != nil {
		log.Println(err)
		return false
	}
	ret, _:= stmt.Exec(args...)
	if ret == nil {
		return false
        }
        return true
} 
func (this *MysqlBase)GetDb()(*sql.DB){
	return this.DB;
}

/*
func main() {
	mydb := GetMysqlInstance()
	inSql := "INSERT INTO dongtai_news(id, imageid, title, create_time) VALUES(?, ?, ?, ?)"
	ret := mydb.insert(inSql, 1, 1, "text", "2016-02-13 00:00:00") 
	fmt.Println(ret)
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
}

*/
