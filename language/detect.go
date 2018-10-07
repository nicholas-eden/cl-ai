package language

import (
	"bufio"
	"fmt"
	"github.com/algorithmiaio/algorithmia-go"
	"os"
	"sync"
)

func DetectLanguage(userInput []string, apiKey string) {
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
	input["text"] = userInput

	var client = algorithmia.NewClient(apiKey, "")
	algo, _ := client.Algo("api2ninja/LanguageDetector/0.2.0")
	resp, _ := algo.Pipe(input)
	response := resp.(*algorithmia.AlgoResponse)

	var highestScore float64
	var highestLang string
	for _, item := range response.Result.([]interface{}) {
		resultMap := item.(map[string]interface{})

		for lang, score := range resultMap {
			if score.(float64) > highestScore {
				highestScore = score.(float64)
				highestLang = lang
			}
		}

	}

	fmt.Printf("%v : %v\n", highestLang, userInput)

	wg.Done()
}
