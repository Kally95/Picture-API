package transport

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kally95/pictureapi/internal"
)

// WriteImage returns an image to the user based on the animal
// they specify within the body of the request using the GET method.
func WriteImage(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	animal := string(data)
	image, err := internal.RetrievePicture(animal)
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
