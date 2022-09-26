package api

import (
	database "rest/internal/db"
)

func (item ItemImport) isCorrect() bool {
	if item.Id == item.ParentId {
		return false
	}
	if item.Size <= 0 {
		return false
	}
	if len(item.Info) > 255 {
		return false
	}
	return true
}

func (item *Item) castDb(dbNode database.Item) {
	item.Id = dbNode.Id
	item.Info = dbNode.Info
	item.ParentId = dbNode.ParentId
	item.Size = dbNode.Size
}

func (node *Item) getAllChildren(db *database.DB) error {
	children, err := db.GetChildren(node.Id)
	if err != nil {
		return err
	}
	for i, e := range children {
		var item Item
		item.castDb(e)
		node.Children = append(node.Children, item)
		node.Children[i].getAllChildren(db)
	}
	return nil
}
