package router

import (
	"database/sql"
	"log"
	"net/http"
)

func Nearby(db *sql.DB) (path string, handler HandlerFunc) {
	return "/nearby", func(w http.ResponseWriter, r *http.Request) error {

		// TODO implement endpoint when area code is not provided, maybe return all values?
		areaCode := r.URL.Query().Get("area_code")
		rows, err := db.Query("SELECT nearby_area_code FROM nearby_area_codes WHERE area_code=?", areaCode)
		if err != nil {
			w.WriteHeader(500)
			return err
		}
		defer rows.Close()
		for rows.Next() {
			var stub interface{}
			if err := rows.Scan(&stub); err != nil {
				w.WriteHeader(500)
				return err
			}
			log.Printf("%+v", stub)
		}
		return nil
	}
}
