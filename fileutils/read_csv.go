package fileutils

import (
	"encoding/csv"
	"fmt"
	"os"
	"sg/solution/datatypes"
	"strconv"
)

func readCsv(path string) ([]datatypes.LtvRecord, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %v", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("unable to read csv file: %v", err)
	}

	var ltvRecords []datatypes.LtvRecord
	for i, record := range records {
		// skip header
		if i == 0 {
			continue
		}

		ltvRecord := datatypes.LtvRecord{Users: 1}
		for j, field := range record {
			switch {
			case j == 1:
				ltvRecord.Campaign = field
			case j == 2:
				ltvRecord.Country = field
			case j > 2:
				floatVal, err := strconv.ParseFloat(field, 64)
				if err != nil {
					return nil, fmt.Errorf("unable to parse float from row %d, field#%d %v", i, j, field)
				}

				if floatVal > 0 {
					ltvRecord.Revenue = append(ltvRecord.Revenue, floatVal)
				}
			}
		}

		ltvRecords = append(ltvRecords, ltvRecord)
	}

	return ltvRecords, nil
}
