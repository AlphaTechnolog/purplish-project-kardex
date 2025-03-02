package main

import (
	"log"

	"github.com/alphatechnolog/purplish-kardex/core"
	"github.com/alphatechnolog/purplish-kardex/database"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.OpenDBConnection()
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
		return
	}
	defer db.Close()

	r := gin.Default()

	core.CreateTransactionsRoutes(db, r.Group("/transactions/"))

	if err = r.Run(); err != nil {
		log.Fatal(err)
	}
}
