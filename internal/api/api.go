package api

import (
	"github.com/gorilla/mux"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	database "rest/internal/db"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func handleImports(db *database.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req := ImportRequest{}
		bs := make([]byte, 1014)
		r.Body.Read(bs)
		err := json.Unmarshal(bs, &req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		for _, e := range req.Items {
			if !e.isCorrect() {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			err = db.AddItem(database.Item{
				Id:       e.Id,
				Info:     e.Info,
				ParentId: e.ParentId,
				Size:     e.Size,
			})
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
		w.WriteHeader(http.StatusOK)
	}
}

func handleDelete(db *database.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		dbNode, err := db.GetItem(id)
		if err != nil {
			if err.Error() == "404" {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
			return
		}
		db.DeleteAllChildren(dbNode)
		w.WriteHeader(http.StatusOK)
	}
}

func handleNodes(db *database.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		dbNode, err := db.GetItem(id)
		if err != nil {
			if err.Error() == "404" {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
			return
		}
		var item Item
		item.castDb(dbNode)
		item.getAllChildren(db)
		bytes, err := json.Marshal(item)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Write(bytes)
	}
}

func Init(r *mux.Router, db *database.DB) {
	r.HandleFunc("/imports", handleImports(db)).Methods(http.MethodPost)
	r.HandleFunc("/delete/{id:.+}", handleDelete(db)).Methods(http.MethodDelete)
	r.HandleFunc("/nodes/{id:.+}", handleNodes(db)).Methods(http.MethodGet)
}
