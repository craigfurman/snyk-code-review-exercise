package api_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/snyk/snyk-code-review-exercise/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPackageHandler(t *testing.T) {
	handler := api.New()
	server := httptest.NewServer(handler)
	defer server.Close()

	resp, err := server.Client().Get(server.URL + "/package/react/16.13.0")
	require.Nil(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)

	var data api.PackageHandlerResponse
	err = json.Unmarshal(body, &data)
	require.Nil(t, err)

	assert.Equal(t, "react", data.Name)
	assert.Equal(t, "16.13.0", data.Version)
	assert.Equal(t, "^1.1.0", data.Dependencies["loose-envify"])
	assert.Equal(t, "^4.1.1", data.Dependencies["object-assign"])
	assert.Equal(t, "^15.6.2", data.Dependencies["prop-types"])
}
