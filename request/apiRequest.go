package request

import (
	"CallDigesto/models"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

const JSONPAGESIZE = 10

// APIRequest makes an API request to the specified URL using the specified HTTP method and authentication header.
// It returns a models.ResponseBody struct containing the API response body and an error (if any).
func APIRequest(url, method string, auth string, request models.ReadCsv) (models.ResponseBody, error) {
	// Create a new BodyRequest struct with the document ID and pagination settings for the initial API call.
	req := models.BodyRequest{
		Document: request.Document,
		Pages: models.Pagination{
			Size: JSONPAGESIZE,
		},
	}

	// Serialize the BodyRequest struct to JSON.
	jsonReq, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		return models.ResponseBody{}, err
	}

	// Create a new buffer with the JSON-encoded request body.
	reqBody := bytes.NewBuffer(jsonReq)

	// Make the API call and get the response.
	res, err := call(url, method, auth, reqBody)
	if err != nil {
		log.Println(err)
		return models.ResponseBody{}, errors.New(err.Error() + "  " + req.Document)
	}

	// Read the response body.
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return models.ResponseBody{}, err
	}

	// Unmarshal the response body into a ResponseBody struct.
	var response models.ResponseBody
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println(err)
		return models.ResponseBody{}, err
	}

	// If the API response has more pages of data, make additional API calls and append the results to the response.
	if response.Pagination.HasNextPage {
		lawsuits, err := callNextPage(url, method, auth, request.Document, response.Pagination.EndCursor)
		if err != nil {
			return models.ResponseBody{}, err
		}

		response.Lawsuits = append(response.Lawsuits, lawsuits...)

		return models.ResponseBody{
			Identification: response.Identification,
			Name:           response.Name,
			Pagination:     response.Pagination,
			Lawsuits:       response.Lawsuits,
		}, nil

	}

	return models.ResponseBody{
		Identification: response.Identification,
		Name:           response.Name,
		Pagination:     response.Pagination,
		Lawsuits:       response.Lawsuits,
	}, nil
}

// callNextPage returns a slice of models.Lawsuit structs containing the data from all pages of the API response.
func callNextPage(url string, method string, auth string, req string, cursor string) ([]models.Lawsuit, error) {
	var lawsuits []models.Lawsuit
	for {
		// Create a new BodyRequest struct with the document ID and updated pagination settings for the next API call.
		request := models.BodyRequest{
			Document: req,
			Pages: models.Pagination{
				Size:   JSONPAGESIZE,
				Cursor: cursor,
			},
		}
		// Serialize the BodyRequest struct to JSON.
		jsonReq, err := json.Marshal(request)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		// Create a new buffer with the JSON-encoded request body.
		reqBody := bytes.NewBuffer(jsonReq)

		// Call the API using the provided url, method, authorization, and request body.
		res, err := call(url, method, auth, reqBody)
		if err != nil {
			log.Println(err)
			return nil, errors.New(err.Error() + "  " + request.Document + "  " + request.Pages.Cursor)
		}

		// Read the response body.
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		// Unmarshal the response body into a models.ResponseBody struct.
		var response models.ResponseBody
		err = json.Unmarshal(body, &response)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		// Append the current response to the lawsuits slice.
		lawsuits = append(lawsuits, response.Lawsuits...)

		// If the API response indicates there are no more pages, break out of the loop.
		if !response.Pagination.HasNextPage {
			break
		}

		// Update the cursor for the next API call.
		request.Pages.Cursor = response.Pagination.EndCursor
	}

	return lawsuits, nil
}

// call sends an HTTP request to the specified URL using the specified method and request body, with the specified authorization header.
// It returns the HTTP response or an error if the request fails.
func call(url, method string, AUTH string, body io.Reader) (*http.Response, error) {
	// Create an HTTP client with a 10-second timeout.
	client := &http.Client{Timeout: time.Second * 10}

	// Create a new HTTP request with the specified method, URL, and request body.
	req, err := http.NewRequest(method, url+AUTH, body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Set the Content-Type and Authorization headers for the request.
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", AUTH)

	// Send the request and get the response.
	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// If the response status code is not OK, return an error with the status code.
	if response.StatusCode != http.StatusOK {
		return nil, errors.New(strconv.Itoa(response.StatusCode))
	}

	return response, nil
}
