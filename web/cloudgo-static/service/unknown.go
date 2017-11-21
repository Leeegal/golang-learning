package service

import (
	"fmt"
	"net/http"
)

func unknownHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "501 Not Implemented...")
}
