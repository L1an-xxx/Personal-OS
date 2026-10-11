package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"

	"personalos/backend/internal/trade"
)

func main() {
	db, err := sql.Open("sqlite", "personalos.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tradeStore := trade.NewStore(db)
	if err := tradeStore.Migrate(); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	trade.NewHandler(tradeStore).Register(r)

	log.Fatal(r.Run(":8080"))
}
