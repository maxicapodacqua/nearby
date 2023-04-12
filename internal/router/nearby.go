package router

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type NearbyResponse struct {
	AreaCodes []int `json:"area_codes"`
}

var ErrInvalidAreaCode = errors.New("invalid value for `area_code`, value must be an integer")

// Nearby
// @Summary Returns a list of nearby area codes based on the input
// @Produce json
// @Tags			nearby
// @Accept			json
// @Param area_code query string true "Area code to search for nearby values"
// @Success 200 {object} GeneralResponse[NearbyResponse] "List of nearby area codes"
// @Failure	422	{object} GeneralResponse[string] "Invalid area code sent"
// @Failure	500	{object} GeneralResponse[string] "Server error, likely database is down"
// @Router			/nearby [get]
func Nearby(db *sql.DB) (path string, handler HandlerFunc) {
	return "/nearby", func(w http.ResponseWriter, r *http.Request) error {

		qsInput := r.URL.Query().Get("area_code")
		areaCode, err := strconv.Atoi(qsInput)
		if err != nil {
			w.WriteHeader(422) // Invalid parameter sent
			rMarshal, _ := json.Marshal(NewErrorResponse(ErrInvalidAreaCode))
			w.Write(rMarshal)
			return err
		}

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
