package router

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/maxicapodacqua/nearby/internal/config"
	"github.com/maxicapodacqua/nearby/internal/database/sqlite"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	path = "/nearby"
)

func startupDB(t *testing.T) *sql.DB {
	t.Setenv(config.SQLiteDNS, ":memory:")

	db := sqlite.Connect()
	if err := sqlite.StartSchema(db, "../../data/schema.sql"); err != nil {
		t.Errorf("Cannot start sqlite schema %v", err)
	}
	return db
}

func truncateTable(db *sql.DB) {
	db.Exec("DELETE FROM nearby_area_codes")
}
func TestNearby(t *testing.T) {
	db := startupDB(t)

	tests := []struct {
		name         string
		callPath     string
		startup      func()
		teardown     func()
		wantStatus   int
		wantError    error
		wantResponse interface{}
	}{
		{
			name:     "Search by area code and return values",
			callPath: path + "?area_code=111",
			startup: func() {
				db.Exec("INSERT INTO nearby_area_codes (area_code, nearby_area_code) VALUES (111, 222)")
				db.Exec("INSERT INTO nearby_area_codes (area_code, nearby_area_code) VALUES (111, 333)")
				db.Exec("INSERT INTO nearby_area_codes (area_code, nearby_area_code) VALUES (111, 444)")
			},
			wantStatus:   200,
			wantError:    nil,
			wantResponse: NewResponse(NearbyResponse{AreaCodes: []int{222, 333, 444}}),
		},
		{
			name:         "Returns empty array",
			callPath:     path + "?area_code=111",
			startup:      func() {},
			wantStatus:   200,
			wantError:    nil,
			wantResponse: NewResponse(NearbyResponse{AreaCodes: []int{}}),
		},
		{
			name:         "Invalid parameter sent",
			callPath:     path + "?area_code=not-integer",
			startup:      func() {},
			wantStatus:   422,
			wantError:    errors.New("strconv.Atoi: parsing"),
			wantResponse: NewErrorResponse(ErrInvalidAreaCode),
		},
		{
			name:     "Table doesn't exist",
			callPath: path + "?area_code=111",
			startup: func() {
				db.Exec("DROP TABLE nearby_area_codes")
			},
			teardown: func() {
				// Restarts db connection and schema for future tests executions
				if err := db.Close(); err != nil {
					t.Errorf("Something went wrong restaring database connection %v", err)
				}
				db = startupDB(t)
			},
			wantStatus:   500,
			wantError:    errors.New("SQL logic error: no such table"),
			wantResponse: NewResponse(NearbyResponse{AreaCodes: []int{}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.startup != nil {
				tt.startup()
			}
			_, gotHandler := Nearby(db)

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", tt.callPath, nil)

			gotErr := gotHandler(w, r)

			result := w.Result()
			gotStatusCode := result.StatusCode
			if gotStatusCode != tt.wantStatus {
				t.Errorf("Nearby() gotStatusCode = %v, want %v", gotStatusCode, tt.wantStatus)
			}

			if gotErr != nil && tt.wantError == nil {
				t.Errorf("Nearby() gotErr = %v, want %v", gotErr, tt.wantError)
			}
			if gotErr != nil && tt.wantError != nil && !strings.Contains(gotErr.Error(), tt.wantError.Error()) {
				t.Errorf("Nearby() gotErr = %v, want %v", gotErr, tt.wantError)
			}

			// if status is not a server error, we should expect a response body
			if tt.wantStatus < 500 {
				gotResponse, _ := io.ReadAll(result.Body)
				wantResponseJSON, _ := json.Marshal(tt.wantResponse)
				if string(gotResponse) != string(wantResponseJSON) {
					t.Errorf("Nearby() gotResponse = %v, want %v", gotResponse, wantResponseJSON)
				}
			}

			if tt.teardown != nil {
				tt.teardown()
			}
			truncateTable(db)
		})
	}
}
