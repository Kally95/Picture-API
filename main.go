package main

import (
	"net/http"
)

const Path = "/home/santito/Desktop/PictureAPI/images/"
const Ext = ".jpeg"

func main() {
	http.HandleFunc("/", WriteImage)
	http.ListenAndServe(":3000", nil)
}

// Consider a funcion that returns a slice of strings containing
// the names of all images, loop through that using the input value
// and extension variable to extraxt desired animal.
