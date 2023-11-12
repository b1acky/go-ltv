package fileutils

import (
	"fmt"
	"path/filepath"
	"sg/solution/datatypes"
)

func Read(path string) ([]datatypes.LtvRecord, error) {
	switch filepath.Ext(path) {
	case ".csv":
		return readCsv(path)
	case ".json":
		return readJson(path)
	default:
		return nil, fmt.Errorf("only .json and .csv file types are supported")
	}
}
