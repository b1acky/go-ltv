package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"sg/solution/fileutils"
	"sg/solution/model"
	"sg/solution/results"
)

func main() {
	modelPtr := flag.String("model", "naive-linear", "model type")
	sourcePtr := flag.String("source", "./data/data.csv", "path to source file")
	aggregatePtr := flag.String("aggregate", "campaign", "aggregation key")

	flag.Parse()

	aggregate, err := results.NewAggregationKey(*aggregatePtr)
	if err != nil {
		fmt.Printf("%s", err)

		return
	}

	model, err := model.NewModel(*modelPtr)
	if err != nil {
		fmt.Printf("Cannot make model: %v", err)

		return
	}

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Cannot get working dir: %v", err)

		return
	}

	path := path.Join(currentDir, *sourcePtr)

	records, err := fileutils.Read(path)
	if err != nil {
		fmt.Printf("Cannot read file: %s", err)

		return
	}

	for _, record := range records {
		model.Train(record)
	}

	result := make(results.AggregationResult)
	settings := results.AggregationSettings{Day: 60, Key: aggregate}

	for _, record := range records {
		prediction := model.Predict(record, settings.Day)

		results.Aggregate(prediction, record, settings, result)
	}

	result.Print()
}
