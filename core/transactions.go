package core

import (
	"database/sql"
	"net/http"

	"github.com/alphatechnolog/purplish-kardex/database"
	"github.com/gin-gonic/gin"
)

func getTransactionsHistory(d *sql.DB, c *gin.Context) error {
	warehouseID := c.Query("warehouseID")
	itemID := c.Query("itemID")

	transactions, err := database.GetTransactionHistory(d, itemID, warehouseID)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{"transactions": transactions})

	return nil
}

func CreateTransactionsRoutes(d *sql.DB, r *gin.RouterGroup) {
	r.GET("/history", WrapError(WithDB(d, getTransactionsHistory)))
}
