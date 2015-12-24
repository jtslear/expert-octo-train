package httpService

import (
	"fmt"
	"github.com/jtslear/expert-octo-train/myLogger"
	"net/http"
	"os"
)

// StartServer starts our web service
func StartServer() interface{} {
	if os.Getenv("OPENSHIFT_GO_IP") == "" {
		os.Setenv("OPENSHIFT_GO_IP", "127.0.0.1")
	}
	if os.Getenv("OPENSHIFT_GO_PORT") == "" {
		os.Setenv("OPENSHIFT_GO_PORT", "5000")
	}
	bind := fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))
	fmt.Printf("listening on %s...\n", bind)

	return http.ListenAndServe(bind, myLogger.Log(http.DefaultServeMux))
}
