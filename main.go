package main

import (
	"log"

	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/pkg/routes"
	"github.com/notRaihan/GoLomba-BE-GDSC-Final-Project/platform/database"
)

func main() {
	database.InitDB()
	app := routes.New()

	log.Fatal(app.Listen(":3000"))
}
