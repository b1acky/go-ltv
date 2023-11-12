package datatypes

import (
	"fmt"
)

type LtvRecord struct {
	Campaign string
	Country  string
	Revenue  []float64
	Users    int
}

func (l LtvRecord) String() string {
	return fmt.Sprintf("campaign: %s, country: %s, users: %d, ltvs: %d, revenue: %v\n", l.Campaign, l.Country, l.Users, len(l.Revenue), l.Revenue)
}
