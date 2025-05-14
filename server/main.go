package main

import (
	"fmt"
	"patient_portal/config"
	"patient_portal/middleware"
	"patient_portal/router"
)

func main() {
	cfg := config.Load()
	db := config.ConnectDB(cfg)
	defer db.Close()

	fmt.Println("Connected to the database")

	r := router.NewRouter(db, cfg)
	r.Use(middleware.CORSMiddleware())
	r.Run(":8080")
}
