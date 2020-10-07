package helpers

import (
	"net/http"
)

func GetFileContentType(buffer []byte) (string, error) {
	// Use the net/http package's handy DetectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}
