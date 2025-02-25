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

	if warehouseID == "" || itemID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return nil
	}

	transactions, err := database.GetTransactionHistory(d, itemID, warehouseID)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{"transactions": transactions})

	return nil
}

func getItemStock(d *sql.DB, c *gin.Context) error {
	warehouseID := c.Query("warehouseID")
	itemID := c.Query("itemID")

	if warehouseID == "" || itemID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return nil
	}

	stock_entries, err := database.GetItemStock(d, itemID, warehouseID)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{"stock_entries": stock_entries})

	return nil
}

func getInventoryCost(d *sql.DB, c *gin.Context) error {
	return nil
}

func CreateTransactionsRoutes(d *sql.DB, r *gin.RouterGroup) {
	r.GET("/history", WrapError(WithDB(d, getTransactionsHistory)))
	r.GET("/stock", WrapError(WithDB(d, getItemStock)))
	r.GET("/inventory-cost", WrapError(WithDB(d, getInventoryCost)))
}
