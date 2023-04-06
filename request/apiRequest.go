package request

import (
	"CallDigesto/models"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func APIRequest(url, method string, auth string, request models.ReadCsv) (models.ResponseBody, error) {
	req := models.BodyRequest{
		Document: request.Document,
		Pages: models.Pagination{
			Size:   1,
			Cursor: "",
		},
	}

	jsonReq, err := json.Marshal(req) // Serialize struct to JSON
	if err != nil {
		return models.ResponseBody{}, err
	}

	reqBody := bytes.NewBuffer(jsonReq)
	res, err := call(url, method, auth, reqBody)
	if err != nil {
		return models.ResponseBody{}, errors.New(err.Error() + "  " + req.Document)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return models.ResponseBody{}, err
	}

	var response models.ResponseBody
	err = json.Unmarshal(body, &response)
	if err != nil {
		return models.ResponseBody{}, err
	}

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

func callNextPage(url string, method string, auth string, req string, cursor string) ([]models.Lawsuit, error) {
	var lawsuits []models.Lawsuit
	for {
		request := models.BodyRequest{
			Document: req,
			Pages: models.Pagination{
				Size:   500,
				Cursor: cursor,
			},
		}

		jsonReq, err := json.Marshal(request) // Serialize struct to JSON
		if err != nil {
			return nil, err
		}

		reqBody := bytes.NewBuffer(jsonReq)
		res, err := call(url, method, auth, reqBody)
		if err != nil {
			return nil, errors.New(err.Error() + "  " + request.Document)
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var response models.ResponseBody
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}

		lawsuits = append(lawsuits, response.Lawsuits...) // Append current response to allResponses

		if !response.Pagination.HasNextPage {
			break // Break the loop if hasNextPage is false
		}

		// Update the cursor for the next API call
		request.Pages.Cursor = response.Pagination.EndCursor
	}

	return lawsuits, nil
}

func call(url, method string, AUTH string, body io.Reader) (*http.Response, error) {
	client := &http.Client{Timeout: time.Second * 10}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json") // Set Content-Type header to application/json
	req.Header.Set("Authorization", AUTH)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(strconv.Itoa(response.StatusCode))
	}

	return response, nil
}
