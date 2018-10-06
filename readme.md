

## Sentiment Analysis
Based on the [Sentiment Analysis Model](https://algorithmia.com/algorithms/nlp/SentimentAnalysis).

Takes input from either stdin or as args and returns a sentiment score for each value.  Multiple values can be passed in on new lines through stdin.

Values are sent through the model asynchronously and may not be in the same order they were submitted. 

Using stdin:
```
cat sentiment-test.txt | ./clai sentiment | sort
```
```
0.8074 : Cats are pretty awesome
0.6369 : I love Golang
-0.5859 : Wasps are a bunch of assholes
```

Using arguments:
```
./clai sentiment "I like cake" "You can pry it out of my cold dead hands" "Tillamook icecream is the best in the country"
```
```
-0.6486 : You can pry it out of my cold dead hands
0.6369 : Tillamook icecream is the best in the country
0.3612 : I like cake
```




## Config

Create a config file at ~/.clai.yaml

Example:
```
algorithmia:
  key: "yourkeyhere"

```

| Key | Description |
| --- | --- |
| `algorithmia.key` | Generate one at https://algorithmia.com | 