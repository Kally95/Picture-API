package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const path = "/home/santito/Desktop/PictureAPI/images/"

func writeImage(w http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	check(err)
	animal := string(d)
	image := retrievePicture(animal)
	if image != nil {
		w.Header().Set("Content-Type", "image/jpeg")
		w.WriteHeader(200)
		w.Write(image)
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, `Sorry! We don't have the image you're looking for. Try "Cat"`)
	}
}

func retrievePicture(a string) []byte {
	str := strings.Title(strings.ToLower(a))
	if _, err := os.Stat(path + str + ".jpeg"); os.IsNotExist(err) {
		return nil
	}
	image, err := ioutil.ReadFile(path + str + ".jpeg")
	check(err)
	return image
}

func main() {
	http.HandleFunc("/", writeImage)
	http.ListenAndServe(":3000", nil)
}
