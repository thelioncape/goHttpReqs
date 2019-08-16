package httpReqs

import (
	"io"
	"log"
	"net/http"
)

func PutRequest(url string, data io.Reader) *http.Response {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, data)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	return res
}
