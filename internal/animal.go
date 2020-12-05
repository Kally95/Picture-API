package internal

import (
	"io/ioutil"
	"os"
	"strings"
)

const Path = "/home/santito/Desktop/PictureAPI/images/"
const Ext = ".jpeg"

// RetrievePicture obtains the specified animal from the
// "images" folder, where we are storing pictures in memory.
func RetrievePicture(a string) ([]byte, error) {
	str := strings.Title(strings.ToLower(a))
	if _, err := os.Stat(Path + str + Ext); os.IsNotExist(err) {
		return nil, err
	}
	image, err := ioutil.ReadFile(Path + str + Ext)
	if err != nil {
		return nil, err
	}
	return image, nil
}
