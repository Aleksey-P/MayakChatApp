package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/Aleksey-P/MayakChatApp/config"
	"fmt"
)
var MySQL *sql.DB

func init() {
	dbConfig := config.MYSQL["USERNAME"].(string) + ":" + config.MYSQL["PASSWORD"].(string) + "@/" + config.MYSQL["DATABASE_NAME"].(string)
	conn, err := sql.Open("mysql", dbConfig)
	if err != nil {
		fmt.Println("ERROR CONNECTION")
	}
	MySQL = conn
}
