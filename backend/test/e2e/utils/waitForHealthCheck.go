package e2e_utils

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func WaitForHealthCheck(baseUrl string) {
	for i := 0; i < 10; i++ {
		response, err := http.Get(fmt.Sprintf("%s/health", baseUrl))

		if err != nil {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		if response.StatusCode == http.StatusOK {
			return
		}
	}
	log.Fatal("Server did not start")
}
