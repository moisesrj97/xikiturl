package e2e

import (
	"backend/test/e2e/utils/prepare_db"
	"io"
	"net/http"
	"testing"
)

import (
	"backend/private/apps/api"
	"backend/test/e2e/utils/test_container"
	"backend/test/e2e/utils/wait_for_health_check"
)

func TestRedirect(t *testing.T) {
	testContainer := test_container.NewMySqlTestContainer()
	t.Cleanup(func() {
		defer testContainer.Stop()
	})

	connectionString, err := testContainer.Start()
	if err != nil {
		t.Fatalf("Error starting container: %v", err)
	}

	err = prepare_db.PrepareDb(connectionString)

	if err != nil {
		t.Fatalf("Error preparing db: %v", err)
	}

	go func() {
		api.StartServer()
	}()

	wait_for_health_check.WaitForHealthCheck("http://localhost:8080")

	want := "Hello, World!"

	response, err := http.Get("http://localhost:8080/hello")

	if err != nil {
		t.Errorf("Error making GET request: %v", err)
	}

	got, err := io.ReadAll(response.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Errorf("Error closing response body: %v", err)
		}
	}(response.Body)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	if string(got) != want {
		t.Errorf("Got %s, want %s", got, want)
	}
}
