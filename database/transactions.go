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

func GetTransactionHistory(d *sql.DB, itemID string, warehouseID string) ([]Transaction, error) {
	var transactions []Transaction

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
