package results

import (
	"fmt"
	"reflect"
	"sg/solution/datatypes"
	"strings"
)

type aggregationCounter struct {
	revenue float64
	users   int
}
type AggregationResult map[string]aggregationCounter

type AggregationKey int

type AggregationSettings struct {
	Key AggregationKey
	Day int
}

const (
	Campaign AggregationKey = iota
	Country
)

var (
	mapKeys = map[string]AggregationKey{
		"campaign": Campaign,
		"country":  Country,
	}
)

func NewAggregationKey(str string) (AggregationKey, error) {
	c, ok := mapKeys[strings.ToLower(str)]
	if !ok {
		return 0, fmt.Errorf("invalid aggregate '%s', available: %v", str, reflect.ValueOf(mapKeys).MapKeys())
	}

	return c, nil
}

func Aggregate(predicted float64, record datatypes.LtvRecord, settings AggregationSettings, result AggregationResult) {
	key := getRecordKey(record, settings.Key)

	counter := result[key]

	counter.revenue += predicted * float64(record.Users)
	counter.users += record.Users

	result[key] = counter
}

func getRecordKey(record datatypes.LtvRecord, key AggregationKey) string {
	switch key {
	case Campaign:
		return record.Campaign
	case Country:
		return record.Country
	default:
		panic(fmt.Sprintf("invalid aggregation key '%v'", key))
	}
}

func (r AggregationResult) Print() {
	for key, value := range r {
		fmt.Printf("%s: %f\n", key, value.revenue/float64(value.users))
	}
}
