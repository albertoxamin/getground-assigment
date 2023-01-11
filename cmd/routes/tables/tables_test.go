package tables

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

func TestAddTable(t *testing.T) {
	// setup
	gin.SetMode(gin.TestMode)
	// create a mock database to use for testing
	r := gin.Default()
	h := db.Init("root:password@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	RegisterRoutes(r, h)
	ts := httptest.NewServer(r)
	defer ts.Close()

	// define test cases
	tests := []struct {
		capacity         int
		expectedStatus   int
		expectedResponse models.Table
	}{
		{
			capacity:       4,
			expectedStatus: http.StatusCreated,
			expectedResponse: models.Table{
				Capacity: 4,
			},
		},
		{
			capacity:         -1,
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: models.Table{},
		},
	}

	// run tests
	for _, test := range tests {
		body := `{"capacity":` + strconv.Itoa(test.capacity) + `}`
		resp, _ := http.Post(fmt.Sprintf("%s/tables", ts.URL), "application/json", strings.NewReader(body))

		// check the status code
		assert.Equal(t, test.expectedStatus, resp.StatusCode)

		// check the response body
		if test.expectedStatus == http.StatusOK {
			assert.Equal(t, test.expectedResponse.Capacity, test.capacity)
		}
	}
	h.Exec("DELETE FROM tables;")
}
