package check_test

import (
	"context"
	"http_healthchecker/internal/check"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func textOKHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func TestSTextCheckerSuccess(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(textOKHandler))
	client := http.Client{}

	checker := check.NewTextChecker(&client)

	err := checker.Check(ctx, srv.URL, map[string]interface{}{
		"expected_text": "ok",
	})
	assert.NoError(t, err)
}
