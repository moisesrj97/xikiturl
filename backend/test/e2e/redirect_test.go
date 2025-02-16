package e2e

import (
	"io"
	"net/http"
	"testing"
)
import (
	"backend/private/apps/api"
	"backend/test/e2e/utils"
)

func TestRedirect(t *testing.T) {
	go func() {
		api.StartServer()
	}()

	e2e_utils.WaitForHealthCheck("http://localhost:8080")

	want := "Hello, World!"

	response, err := http.Get("http://localhost:8080/hello")

	if err != nil {
		t.Errorf("Error making GET request: %v", err)
	}

	got, err := io.ReadAll(response.Body)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	if string(got) != want {
		t.Errorf("Got %s, want %s", got, want)
	}
}
