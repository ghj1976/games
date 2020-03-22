package tank90

import (
	"database/sql"
	"log"

	// sqlite 数据库
	_ "github.com/mattn/go-sqlite3"
)

// getMapByName 从数据库中加载地图数据
func getMapByName(mapid int) string {
	db, err := sql.Open("sqlite3", "./tank.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("select data from map where name = ?")
	if err != nil {
		log.Fatal(err)
	}
	// QueryRow执行一次查询，并期望返回最多一行结果（即Row）
	row := stmt.QueryRow(mapid)

	var mapData string
	err = row.Scan(&mapData)
	if err != nil {
		return ""
		log.Fatal(err)
	}
	return mapData
}
