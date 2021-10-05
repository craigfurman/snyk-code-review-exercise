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

	var data api.NpmPackageVersion
	err = json.Unmarshal(body, &data)
	require.Nil(t, err)

	assert.Equal(t, "react", data.Name)
	assert.Equal(t, "16.13.0", data.Version)

	// TODO these assertions are rubbish, even for the intentionally-not-great PR.
	// Let's talk. We could hardcode a dependency tree for some package version,
	// e.g. react 16.13.0. These tests will break over time as new patch versions
	// of dependencies are released that satisfy the same constraints, but that's
	// almost a good thing for the exercise. Discussing what we'd want IRL is a
	// good interview topic.
	assert.Len(t, data.Dependencies, 3)
}
