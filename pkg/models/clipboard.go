package models

import (
	"database/sql"
)

type ClipboardItem struct {
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}

// Create a custom ClipboardItemModel type which wraps the sql.DB connection pool.
type ClipboardItemModel struct {
	DB *sql.DB
}

// Use a method on the custom ClipboardModel type to run the SQL query.
func (m ClipboardItemModel) All() ([]ClipboardItem, error) {
	rows, err := m.DB.Query("SELECT * FROM clipboard_items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []ClipboardItem

	for rows.Next() {
		var item ClipboardItem

		err := rows.Scan(&item.Content, &item.CreatedAt)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
