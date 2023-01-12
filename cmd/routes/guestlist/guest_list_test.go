package guestlist

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"github.com/stretchr/testify/assert"

	"github.com/getground/tech-tasks/backend/cmd/common/db"
	"github.com/getground/tech-tasks/backend/cmd/common/models"
)

type mockDB struct {
	*gorm.DB
}

func TestAddGuest(t *testing.T) {
	// setup
	gin.SetMode(gin.TestMode)
	// create a mock database to use for testing
	r := gin.Default()
	h := db.Init("root:password@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	RegisterRoutes(r, h)
	ts := httptest.NewServer(r)
	defer ts.Close()

	type GuestRes struct {
		Name string
	}

	// create a table to use for testing
	table := models.Table{Capacity: 5}
	h.Create(&table)

	// define test cases
	tests := []struct {
		Name               string
		Table              int
		expectedStatus     int
		AccompanyingGuests int
		expectedResponse   GuestRes
	}{
		{
			Name:               "john",
			Table:              int(table.ID),
			expectedStatus:     http.StatusBadRequest,
			AccompanyingGuests: -1,
		},
		{
			Name:               "john",
			Table:              int(table.ID),
			expectedStatus:     http.StatusOK,
			AccompanyingGuests: 0,
			expectedResponse:   GuestRes{Name: "john"},
		},
	}

	// run tests
	for _, test := range tests {
		body := `{"table":` + strconv.Itoa(test.Table) + `,"accompanying_guests":` + strconv.Itoa(test.AccompanyingGuests) + `}`
		resp, _ := http.Post(fmt.Sprintf("%s/guest_list/%s", ts.URL, test.Name), "application/json", strings.NewReader(body))

		// check the status code
		assert.Equal(t, test.expectedStatus, resp.StatusCode)

		// check the response body
		if test.expectedStatus == http.StatusOK {
			assert.Equal(t, test.expectedResponse.Name, test.Name)
		}
	}
	h.Exec("DELETE FROM guests;")
	h.Exec("DELETE FROM tables;")
}
