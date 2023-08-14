package models

import (
	"database/sql"
	"fmt"
	"time"
)

type ClipboardItem struct {
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}

// Create a custom ClipboardItemModel type which wraps the sql.DB connection pool.
type ClipboardItemModel struct {
	DB *sql.DB
}

func (m *ClipboardItemModel) DeleteAll() error {
	q := "DELETE FROM clipboard_items"

	_, err := m.DB.Exec(q)
	if err != nil {
		return err
	}
	return nil
}

func (m *ClipboardItemModel) DeleteByDateCreated(date string) error {
	q := fmt.Sprintf("DELETE FROM clipboard_items WHERE created_at = '%s'", date)
	_, err := m.DB.Exec(q)
	if err != nil {
		return err
	}
	return nil
}

func (m *ClipboardItemModel) Set(content string) error {

	timeNow := time.Now().Format(time.RFC3339)

	q, err := m.DB.Prepare("INSERT INTO clipboard_items VALUES (?,?)")
	if err != nil {
		fmt.Println("ERROR: could not create parametarised query: ", err)
	}
	fmt.Println("CREATE query", q)

	_, err = q.Exec(content, timeNow)
	if err != nil {
		return err
	}
	return nil
}

func (m *ClipboardItemModel) GetLatest() ([]ClipboardItem, error) {

	q := fmt.Sprintf("SELECT * FROM clipboard_items ORDER BY created_at DESC LIMIT 1")
	result, err := m.DB.Query(q)
	if err != nil {
		return nil, err
	}

	var items []ClipboardItem

	for result.Next() {
		var item ClipboardItem

		err := result.Scan(&item.Content, &item.CreatedAt)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}
	return items, nil
}

// Use a method on the custom ClipboardModel type to run the SQL query.
func (m *ClipboardItemModel) GetAll() ([]ClipboardItem, error) {
	rows, err := m.DB.Query("SELECT * FROM clipboard_items ORDER BY created_at DESC")
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
