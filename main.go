package main

import (
	"pbi-task/database"
	"pbi-task/routes"
)

func main() {
	database.Connect("root:password123@tcp(127.0.0.1)/pbi_task?parseTime=true")
	database.Migrate()

	router := routes.SetUpRouter()
	router.Run(":9988")
}
