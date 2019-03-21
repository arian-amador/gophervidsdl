package gophervidsdl

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

// ValidatePath helper to validate if a path exists
func ValidatePath(p string) error {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return err
	}

	return nil
}

// ReadJSON process json video file and return array of Vidoes
func ReadJSON(in string) []byte {
	jsonFile, err := os.Open(in)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue
}

// Sanitize is used to remove special chars and trim a string
func Sanitize(s string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9\\s]+")
	if err != nil {
		log.Fatal(err)
	}

	s = strings.Trim(s, " ")
	s = reg.ReplaceAllString(s, "")
	s = strings.Replace(s, " ", "-", -1)
	s = strings.ToLower(s)

	return s
}