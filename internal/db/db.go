package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

func (db *DB) Init(filePath string) error {
	dbCon, err := sql.Open("sqlite3", filePath)
	if err != nil {
		return err
	}
	db.db = dbCon
	// db.db.Exec("DROP TABLE IF EXISTS items")
	db.db.Exec(`
		CREATE TABLE IF NOT EXISTS items ( 
			id TEXT UNIQUE,
			info TEXT,
			parentId TEXT,
			size NUMBER
		)`,
	)

	return nil
}

func (db *DB) AddItem(item Item) error {
	_, err := db.db.Exec("REPLACE INTO items ( id, info, parentId, size ) VALUES  ( $1, $2, $3, $4 )",
		item.Id, item.Info, item.ParentId, item.Size)
	return err
}

func (db *DB) DeleteItem(id string) error {
	_, err := db.db.Exec("DELETE FROM items WHERE id = $1", id)
	return err
}

func (db *DB) GetItem(id string) (Item, error) {
	rows, err := db.db.Query("SELECT id, info, parentId, size FROM items WHERE id = $1", id)
	if err != nil {
		return Item{}, err
	}
	defer rows.Close()
	if rows.Next() {
		var item Item
		err = rows.Scan(&item.Id, &item.Info, &item.ParentId, &item.Size)
		if err != nil {
			return Item{}, err
		}
		return item, nil
	}
	return Item{}, &DbError{message: "404"}
}

func (db *DB) GetChildren(parentId string) ([]Item, error) {
	rows, err := db.db.Query("SELECT id, info, parentId, size FROM items WHERE parentId = $1", parentId)
	if err != nil {
		return nil, err
	}
	children := []Item{}
	for rows.Next() {
		var item Item
		err = rows.Scan(&item.Id, &item.Info, &item.ParentId, &item.Size)
		if err != nil {
			continue
		}
		children = append(children, item)
	}
	return children, err
}

func (db *DB) DeleteAllChildren(node Item) error {
	children, err := db.GetChildren(node.Id)
	if err != nil {
		return err
	}
	for _, e := range children {
		db.DeleteAllChildren(e)
	}
	db.DeleteItem(node.Id)
	return nil
}
