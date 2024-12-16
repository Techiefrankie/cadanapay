package model

import (
	"sort"
	"strconv"
)

type Persons struct {
	Data []Person `json:"data"`
}

type Person struct {
	Id         string `json:"id"`
	PersonName string `json:"personName"`
	Salary     struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"salary"`
}

func (p *Persons) Sort(asc bool) *Persons {
	sort.Slice(p.Data, func(i, j int) bool {
		valueI, _ := strconv.ParseFloat(p.Data[i].Salary.Value, 64)
		valueJ, _ := strconv.ParseFloat(p.Data[j].Salary.Value, 64)
		if asc {
			return valueI < valueJ // ascending order
		}
		return valueI > valueJ // descending order
	})

	return p
}

func (p *Persons) FilterByCurrency(currency string) *Persons {
	ps := make([]Person, 0, len(p.Data))

	for _, person := range p.Data {
		if person.Salary.Currency == currency {
			ps = append(ps, person)
		}
	}

	return &Persons{Data: ps}
}

func (p *Persons) GroupBySalaryCurrency() map[string][]Person {
	grouped := make(map[string][]Person)

	for _, person := range p.Data {
		currency := person.Salary.Currency
		grouped[currency] = append(grouped[currency], person)
	}

	return grouped
}
