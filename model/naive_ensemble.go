// "Ансмабль" наивных моделей - учится на отдельных линейках по компании и стране, результат
// отдаёт в виде средневзвешенного значения двух предсказаний, где вес - доля кол-ва юзеров
// в текущей линейке по сравнению с общим кол-вом юзеров во всех линейках

package model

import (
	"sg/solution/datatypes"
)

type modelEnsemble map[string]*Model
type usersCounter map[string]int

type NaiveEnsembleModel struct {
	Model
	byCountry       modelEnsemble
	byCampaign      modelEnsemble
	usersByCountry  usersCounter
	usersByCampaign usersCounter
	totalUsers      int
}

func (m *NaiveEnsembleModel) Train(record datatypes.LtvRecord) {
	country := record.Country
	campaign := record.Campaign

	_, exists := m.byCountry[country]
	if !exists {
		countryModel, err := NewModel("naive-linear")
		if err != nil {
			panic(err)
		}

		m.byCountry[country] = &countryModel
	}

	_, exists = m.byCampaign[campaign]
	if !exists {
		campaignModel, err := NewModel("naive-linear")
		if err != nil {
			panic(err)
		}

		m.byCampaign[campaign] = &campaignModel
	}

	(*m.byCampaign[campaign]).Train(record)
	(*m.byCountry[country]).Train(record)

	m.totalUsers += record.Users
	m.usersByCampaign[campaign] += record.Users
	m.usersByCountry[country] += record.Users
}

func (m *NaiveEnsembleModel) Predict(record datatypes.LtvRecord, day int) float64 {
	countryModel := m.byCountry[record.Country]
	campaignModel := m.byCampaign[record.Campaign]

	predictByCountry := (*countryModel).Predict(record, day)
	predictByCampaign := (*campaignModel).Predict(record, day)

	countryWeight := float64(m.usersByCountry[record.Country]) / float64(m.totalUsers)
	campaignWeight := float64(m.usersByCampaign[record.Campaign]) / float64(m.totalUsers)

	return (predictByCampaign*campaignWeight + predictByCountry*countryWeight) / (campaignWeight + countryWeight)
}
