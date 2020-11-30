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
const ext = ".jpeg"

func writeImage(w http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	check(err)
	animal := string(d)
	image, err := retrievePicture(animal)
	if err != nil {
		w.Header().Set("Content-Type", "image/jpeg")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `Sorry! We don't have the image you're looking for, try "Cat"`)
	} else {
		w.Header().Set("Content-Type", "image/jpeg")
		w.WriteHeader(http.StatusOK)
		w.Write(image)
	}

}

func retrievePicture(a string) ([]byte, error) {
	str := strings.Title(strings.ToLower(a))
	if _, err := os.Stat(path + str + ext); os.IsNotExist(err) {
		return nil, err
	}
	image, err := ioutil.ReadFile(path + str + ext)
	if err != nil {
		return nil, err
	}
	return image, nil
}

func main() {
	http.HandleFunc("/", writeImage)
	http.ListenAndServe(":3000", nil)
}
