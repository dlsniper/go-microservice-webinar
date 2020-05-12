package handler

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

//Home handles the home page
func Home(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		visitorID := 0
		err := db.QueryRow(
			"INSERT INTO visitors(user_agent, datetime) VALUES ($1, now()) RETURNING id",
			r.UserAgent(),
		).Scan(&visitorID)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Internal error"))
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, fmt.Sprintf("{\"status\":  200, \"message\": \"Hello visitor %d!\"}", visitorID))
	}
}
