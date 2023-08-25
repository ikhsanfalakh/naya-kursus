package latihan

import (
	"0702/tugas"
	"encoding/json"
	"fmt"
	"net/http"
)

func UsersNew(url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := tugas.GetUsers(url)
		if err != nil {
			http.Error(w, "Error getting users", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

func UserNew(url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodGet {
			id := r.URL.Query().Get("id")
			users, err := tugas.GetUsers(url)
			if err != nil {
				http.Error(w, "Error getting users", http.StatusInternalServerError)
				return
			}

			for _, user := range users {
				if fmt.Sprintf("%d", user.ID) == id {
					json.NewEncoder(w).Encode(user)
					return
				}
			}

			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Invalid request method", http.StatusBadRequest)
	}
}
