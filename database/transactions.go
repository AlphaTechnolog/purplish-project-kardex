package database

import "database/sql"

type Transaction struct {
	ID          string `json:"id"`
	ItemID      string `json:"item_id"`
	WarehouseID string `json:"warehouse_id"`
	Kind        string `json:"kind"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Cost        int    `json:"cost"`
	CreatedAt   string `json:"created_at"`
}

type StockEntry struct {
	ItemID      string `json:"item_id"`
	WarehouseID string `json:"warehouse_id"`
	Stock       int    `json:"stock"`
	TotalCost   int    `json:"total_cost"`
}

func GetTransactionHistory(d *sql.DB, itemID string, warehouseID string) ([]Transaction, error) {
	var transactions []Transaction = []Transaction{}

	sql := `
		SELECT
			t.id,
			t.item_id,
			t.warehouse_id,
			t.kind,
			t.amount,
			t.description,
			t.cost,
			t.created_at
		FROM
			transactions t
		WHERE
			t.item_id = ?
			AND t.warehouse_id = ?
	`

	rows, err := d.Query(sql, itemID, warehouseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction Transaction
		err = rows.Scan(
			&transaction.ID,
			&transaction.ItemID,
			&transaction.WarehouseID,
			&transaction.Kind,
			&transaction.Amount,
			&transaction.Description,
			&transaction.Cost,
			&transaction.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func GetItemStock(d *sql.DB, itemID string, warehouseID string) ([]StockEntry, error) {
	var entries []StockEntry = []StockEntry{}

	sql := `
		SELECT
			s.item_id,
			s.warehouse_id,
			s.stock,
			s.total_cost
		FROM
			stock_view s
		WHERE
			s.item_id = ?
			AND s.warehouse_id = ?;
	`

	rows, err := d.Query(sql, itemID, warehouseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var entry StockEntry
		err = rows.Scan(&entry.ItemID, &entry.WarehouseID, &entry.Stock, &entry.TotalCost)
		if err != nil {
			return nil, err
		}

		entries = append(entries, entry)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}
