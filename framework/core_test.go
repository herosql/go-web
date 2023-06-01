package framework

import (
	"net/http"
	"testing"
)

func TestCore(t *testing.T) {
	server := &http.Server{
		Handler: NewCore(),
		Addr:    ":8080",
	}
	server.ListenAndServe()
}
