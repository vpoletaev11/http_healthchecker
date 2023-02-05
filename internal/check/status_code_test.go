package check_test

import (
	"context"
	"http_healthchecker/internal/check"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func statusCode200Handler(w http.ResponseWriter, r *http.Request) {

}

func TestSatusCodeCheckerSuccess(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(statusCode200Handler))
	client := http.Client{}

	checker := check.NewStatusCodeChecker(&client)

	err := checker.Check(ctx, srv.URL, map[string]interface{}{
		"expected_status_code": http.StatusOK,
	})
	assert.NoError(t, err)
}
