package model

import (
	"fmt"
	"sg/solution/datatypes"
)

type Model interface {
	Train(records datatypes.LtvRecord)
	Predict(record datatypes.LtvRecord, day int) float64
}

func NewModel(model string) (Model, error) {
	switch model {
	case "naive-linear":
		return &NaiveLinearModel{}, nil
	case "naive-ensemble":
		return &NaiveEnsembleModel{byCountry: make(modelEnsemble), byCampaign: make(modelEnsemble), usersByCountry: make(usersCounter), usersByCampaign: make(usersCounter)}, nil
	default:
		return nil, fmt.Errorf("unknown model type '%s', available: 'naive-lienar', 'naive-ensemble'", model)
	}
}
