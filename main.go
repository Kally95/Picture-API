package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const path = "/home/santito/Desktop/PictureAPI/images/"
const ext = ".jpeg"

// WriteImage returns an image to the user based on the animal
// they specify within the body of the request using the GET method.
func WriteImage(w http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	animal := string(d)
	image, err := RetrievePicture(animal)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `Sorry! We don't have the image you're looking for, try "Cat"`)
	} else {
		w.Header().Set("Content-Type", "image/jpeg")
		w.WriteHeader(http.StatusOK)
		w.Write(image)
	}
}

// RetrievePicture obtains the specified animal from the
// "images" folder, where we are storing pictures in memory.
func RetrievePicture(a string) ([]byte, error) {
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
	http.HandleFunc("/", WriteImage)
	http.ListenAndServe(":3000", nil)
}
