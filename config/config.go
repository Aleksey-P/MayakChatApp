package config

import "os"

var MYSQL = map[string]interface{}{
	"USERNAME": os.Getenv("MYSQL_USERNAME"),
	"PASSWORD": os.Getenv("MYSQL_PASSWORD"),
	"DATABASE_NAME": os.Getenv("MYSQL_DATABASE_NAME"),
}
