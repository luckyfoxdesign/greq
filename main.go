package greq

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func GetHTMLSource(websiteURL, proxyAddress string, headers map[string][]string) ([]byte, error) {

	var errorString string

	proxyURL, err := url.Parse(proxyAddress)

	if err != nil {
		errorString = fmt.Errorf("error when url.parse(proxyAddress): %w", err).Error()
		return []byte{}, errors.New(errorString)
	}

	siteURL, err := url.Parse(websiteURL)
	if err != nil {
		errorString = fmt.Errorf("error when url.parse(websiteURL): %w", err).Error()
		return []byte{}, errors.New(errorString)
	}

	transport := &http.Transport{
		Proxy:                 http.ProxyURL(proxyURL),
		ResponseHeaderTimeout: time.Minute,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   time.Millisecond * 500,
	}

	requestToSite, err := http.NewRequest("GET", siteURL.String(), nil)
	if err != nil {
		errorString = fmt.Errorf("error when http.NewRequest(siteURL): %w", err).Error()
		return []byte{}, errors.New(errorString)
	}

	requestToSite.Header = headers

	requestResponse, err := client.Do(requestToSite)
	if err != nil {
		errorString = fmt.Errorf("Error when client.do(requestToSite): %w", err).Error()
		return []byte{}, errors.New(errorString)
	}
	defer requestResponse.Body.Close()

	readDataInBytes, err := ioutil.ReadAll(requestResponse.Body)
	if err != nil {
		errorString = fmt.Errorf("Error when ioutil.ReadAll(): %w", err).Error()
		return []byte{}, errors.New(errorString)
	}

	return readDataInBytes, err
}
