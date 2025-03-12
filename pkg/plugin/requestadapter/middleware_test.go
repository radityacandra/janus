package requestadapter

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hellofresh/janus/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestAddHeader(t *testing.T) {
	config := Config{
		Mapping: map[string]string{
			"name": "name",
		},
	}
	req, err := http.NewRequest(http.MethodPost, "/", strings.NewReader(`{"name": "some name"}`))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	NewRequestAdapter(config)(http.HandlerFunc(test.Ping)).ServeHTTP(w, req)
}
