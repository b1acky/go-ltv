package fileutils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sg/solution/datatypes"
)

type jsonRecord struct {
	CampaignId string  `json:"CampaignId"`
	Country    string  `json:"Country"`
	Ltv1       float64 `json:"Ltv1"`
	Ltv2       float64 `json:"Ltv2"`
	Ltv3       float64 `json:"Ltv3"`
	Ltv4       float64 `json:"Ltv4"`
	Ltv5       float64 `json:"Ltv5"`
	Ltv6       float64 `json:"Ltv6"`
	Ltv7       float64 `json:"Ltv7"`
	Users      int     `json:"Users"`
}

func readJson(path string) ([]datatypes.LtvRecord, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %v", err)
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %v", err)
	}

	var records []jsonRecord

	err = json.Unmarshal(bytes, &records)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal JSON: %v", err)
	}

	var ltvRecords []datatypes.LtvRecord
	for _, record := range records {
		ltvRecords = append(ltvRecords, convert(record))
	}

	return ltvRecords, nil
}

func convert(record jsonRecord) datatypes.LtvRecord {
	var ltv datatypes.LtvRecord

	ltv.Campaign = record.CampaignId
	ltv.Country = record.Country
	ltv.Users = record.Users
	ltv.Revenue = []float64{record.Ltv1, record.Ltv2, record.Ltv3, record.Ltv4, record.Ltv5, record.Ltv6, record.Ltv7}

	return ltv
}
