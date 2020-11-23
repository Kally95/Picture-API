package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeImage(w http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	check(err)
	animal := string(d)
	image := animalPicture(animal)
	fmt.Print(image)
	w.Write(image)
}

func animalPicture(a string) []byte {
	str := strings.Title(strings.ToLower(a))
	image, err := ioutil.ReadFile("/home/santito/Desktop/PictureAPI/images/" + str + ".jpeg")
	check(err)
	return image
}

func main() {
	http.HandleFunc("/", writeImage)
	http.ListenAndServe(":3000", nil)
}
