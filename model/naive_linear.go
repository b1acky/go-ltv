// Максимально наивная модель, которая прогнозирует выручку с помощью усреднения
// линейных функций, которые строятся по последним двум точкам каждой линейки.
package model

import (
	"sg/solution/datatypes"
)

type NaiveLinearModel struct {
	Model
	avgSlope float64
	count    int
}

func (m *NaiveLinearModel) Train(record datatypes.LtvRecord) {
	recordSlope := (record.Revenue[len(record.Revenue)-1] - record.Revenue[len(record.Revenue)-2])

	m.avgSlope = (recordSlope + (m.avgSlope * float64(m.count))) / (float64(m.count) + 1)
	m.count++
}

func (m *NaiveLinearModel) Predict(record datatypes.LtvRecord, day int) float64 {
	from := record.Revenue[len(record.Revenue)-1]
	days := day - len(record.Revenue)
	avgSlope := m.avgSlope

	return from + float64(days)*avgSlope
}
