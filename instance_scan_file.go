package clamavrestsdk

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
)

func (i *clamAVRestInstance) ScanFile(fileContent []byte) (bool, error) {
	var httpProtocol string = "http"
	if i.apiIsHttps {
		httpProtocol = "https"
	}

	address := fmt.Sprintf("%s://%s:%d/scan-document", httpProtocol, i.apiAddress, i.apiPort)

	// Prepare multipart form
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, err := w.CreateFormFile("file", "upload.bin")
	if err != nil {
		return false, err
	}
	if _, err := fw.Write(fileContent); err != nil {
		return false, err
	}
	w.Close()

	request, requestErr := http.NewRequest(http.MethodPost, address, &b)
	if requestErr != nil {
		return false, requestErr
	}
	request.Header.Set("Content-Type", w.FormDataContentType())

	response, responseErr := http.DefaultClient.Do(request)
	if responseErr != nil {
		return false, responseErr
	}

	defer response.Body.Close()

	isVirusFree := response.StatusCode != http.StatusOK

	return isVirusFree, nil
}
