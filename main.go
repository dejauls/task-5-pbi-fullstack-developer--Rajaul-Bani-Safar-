package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/dejauls/task-5-pbi-fullstack-developer--Rajaul-Bani-Safar-/database"
	"github.com/dejauls/task-5-pbi-fullstack-developer--Rajaul-Bani-Safar-/routes"
)

func main() {
	router := gin.Default()

	database.InitDB()
	database.MigrateDB()

	routes.SetupRoutes(router)

	port := 8080
	fmt.Printf("Server running on :%d\n", port)
	router.Run(fmt.Sprintf(":%d", port))
}
