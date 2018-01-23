package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	textPtr := flag.String("text", "", "Text to parse. (Required)")
	metricPtr := flag.String("metric", "chars", "Metric {chars|words|lines};. (Required)")
	uniquePtr := flag.Bool("unique", false, "Measure unique values of a metric.")
	flag.Parse()

	if *textPtr == "" {
		fmt.Println("no text supplied. I'm exiting.")
		os.Exit(1)
	}
	switch *metricPtr {
	case "words":
		{
			WordCount(*textPtr)
		}

	}

	fmt.Printf("=== textPtr: %s, metricPtr: %s, uniquePtr: %t\n", *textPtr, *metricPtr, *uniquePtr)
}

//functions
func WordCount(s string) int {

	words := strings.Fields(s)
	wordCountMap := make(map[string]int)

	for _, word := range words {
		wordCountMap[word]++
	}
	fmt.Println("%d words found ", len(words)) //wordCountMap
	return len(words)
}
