package latihan

import (
	"encoding/json"
	"net/http"
)

type student struct {
	ID    string
	Name  string
	Grade int
}

var data = []student{
	{"E001", "ethan", 21},
	{"W001", "wick", 22},
	{"B001", "bourne", 23},
	{"B002", "bond", 23},
}

func Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func User(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var id = r.FormValue("id")
		var result []byte
		var err error

		for _, each := range data {
			if each.ID == id {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}
		}

		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
