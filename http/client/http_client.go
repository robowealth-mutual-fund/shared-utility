package httpclient

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// HTTPClient ...
type HTTPClient struct{}

// NewClient ...
func (h HTTPClient) NewClient() (*http.Client, error) {
	// Setup a http client
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: time.Duration(60) * time.Second,
	}
	return httpClient, nil
}

// Request ...
func (h HTTPClient) Request(req *http.Request) (*http.Response, error) {
	client, err := h.NewClient()

	if err != nil {
		errMsg := "Can not create HTTP client"
		log.Printf("Error HTTPClient.Request > %s: %s\n", errMsg, err.Error())

		// return nil, ue.NewInternalError(, errMsg)
		return nil, errors.New(errMsg)
	}

	response, err := client.Do(req)
	if err != nil {
		errMsg := "Can no make request"

		log.Printf("Error HTTPClient.Request > %s: %s\n", errMsg, err.Error())
		// return nil, ue.NewInternalError(, errMsg)
		return nil, errors.New(errMsg)
	}

	log.Printf("HTTPClient.Request > Request To [%s] %s\n", req.Method, req.URL)

	if response.StatusCode != 200 {
		responseBodyString := ""

		responseBody, err := ioutil.ReadAll(response.Body)
		if err == nil {
			responseBodyString = string(responseBody)
		}

		errMsg := fmt.Sprintf("Request Failure: StatusCode=%d ResponseBody=%s", response.StatusCode, responseBodyString)
		log.Println("Error HTTPClient.Request >", errMsg)

		response.Body = ioutil.NopCloser(bytes.NewBuffer(responseBody))
		// return nil, ue.NewInternalError(, errMsg)
		return response, errors.New(errMsg)
	}

	return response, nil
}

// New ...
func New() *HTTPClient {
	return &HTTPClient{}
}
