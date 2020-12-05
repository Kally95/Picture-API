package main

import (
	"net/http"

	"github.com/kally95/pictureapi/transport"
)

func main() {
	http.HandleFunc("/", transport.WriteImage)
	http.ListenAndServe(":3000", nil)
}
