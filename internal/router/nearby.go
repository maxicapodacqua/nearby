package router

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type NearbyResponse struct {
	AreaCodes []int `json:"area_codes"`
}

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

		resp := NearbyResponse{AreaCodes: []int{}}

		for rows.Next() {
			var areaCode int
			if err := rows.Scan(&areaCode); err != nil {
				w.WriteHeader(500)
				return err
			}
			resp.AreaCodes = append(resp.AreaCodes, areaCode)
		}
		rMarshal, err := json.Marshal(NewResponse(resp))
		if err != nil {
			return err
		}

		if _, err = w.Write(rMarshal); err != nil {
			return err
		}
		return nil
	}
}
