package main

import "github.com/Zekeriyyah/stellar-x/internal/database"

func main() {
	database.InitDB()
	database.Run()
}