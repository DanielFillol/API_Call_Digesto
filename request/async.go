package request

import (
	"CallDigesto/models"
	"fmt"
	"sync"
)

func AsyncAPIRequest(users []models.ReadCsv, numberOfWorkers int, url string, method string, auth string) ([]models.ResponseBody, error) {
	done := make(chan struct{})
	defer close(done)

	inputCh := StreamInputs(done, users)

	var wg sync.WaitGroup
	wg.Add(numberOfWorkers)

	resultCh := make(chan models.ResponseBody)

	for i := 0; i < numberOfWorkers; i++ {
		// spawn N worker goroutines, each is consuming a shared input channel.
		go func() {
			for input := range inputCh {
				bodyStr, err := APIRequest(url, method, auth, input)
				resultCh <- bodyStr
				if err != nil {
					fmt.Println(err)
				}
			}
			wg.Done()

		}()
	}

	// Wait all worker goroutines to finish. Happens if there's no error (no early return)
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	var results []models.ResponseBody
	for result := range resultCh {
		results = append(results, result)
	}

	return results, nil
}

func StreamInputs(done <-chan struct{}, inputs []models.ReadCsv) <-chan models.ReadCsv {
	inputCh := make(chan models.ReadCsv)
	go func() {
		defer close(inputCh)
		for _, input := range inputs {
			select {
			case inputCh <- input:
			case <-done:
				// in case done is closed prematurely (because error midway),
				// finish the loop (closing input channel)
				break
			}
		}
	}()
	return inputCh
}
