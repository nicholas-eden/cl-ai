package sentiment

import (
	"bufio"
	"fmt"
	"github.com/algorithmiaio/algorithmia-go"
	"os"
	"sync"
)

func Execute(userInput []string, apiKey string) {
	var wg sync.WaitGroup
	if len(userInput) > 0 {

		for _, row := range userInput {
			wg.Add(1)
			go handleOne(row, &wg, apiKey)
		}

	} else {

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			row := scanner.Text()

			wg.Add(1)
			go handleOne(row, &wg, apiKey)
		}

	}

	wg.Wait()

}

func handleOne(userInput string, wg *sync.WaitGroup, apiKey string) {
	input := make(map[string]interface{})
	input["document"] = userInput

	var client = algorithmia.NewClient(apiKey, "")
	algo, _ := client.Algo("nlp/SentimentAnalysis/1.0.5")
	resp, _ := algo.Pipe(input)
	response := resp.(*algorithmia.AlgoResponse)

	for _, item := range response.Result.([]interface{}) {
		resultMap := item.(map[string]interface{})
		document := resultMap["document"]
		score := resultMap["sentiment"]

		fmt.Printf("%v : %v\n", score, document)
	}

	wg.Done()
}
