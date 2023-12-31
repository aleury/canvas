package server_test

import (
	"net/http"
	"testing"

	"canvas/integrationtest"

	"github.com/matryer/is"
)

func TestServer_Start(t *testing.T) {
	integrationtest.SkipIfShort(t)

	t.Run("starts the server and listens for requests", func(t *testing.T) {
		is := is.New(t)

		cleanup := integrationtest.CreateServer()
		defer cleanup()

		resp, err := http.Get("http://localhost:8081/")
		is.NoErr(err)
		is.Equal(resp.StatusCode, http.StatusOK)
	})
}
