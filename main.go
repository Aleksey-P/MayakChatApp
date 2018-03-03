package main

import (
	"fmt"
	"github.com/Aleksey-P/MayakChatApp/database"
)

func main() {
	fmt.Println("Start Server")
	if err := database.MySQL.Ping(); err != nil {
		panic(err)
	}	
	database.MySQL.Query("INSERT INTO user SET username = ?, password = ?", "test_1", "test")
	fmt.Println("Done")
}

