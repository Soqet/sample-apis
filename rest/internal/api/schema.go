package api

type ItemImport struct {
	Id       string `json:"id"`
	Info     string `json:"info"`
	ParentId string `json:"parentId"`
	Size     uint   `json:"size"`
}

type ImportRequest struct {
	Items []ItemImport `json:"items"`
}

type Item struct {
	Id       string `json:"id"`
	Info     string `json:"info"`
	ParentId string `json:"parentId"`
	Size     uint   `json:"size"`
	Children []Item `json:"children"`
}
