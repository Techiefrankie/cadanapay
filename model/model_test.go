package model

import (
	"reflect"
	"testing"
)

func TestSortAscending(t *testing.T) {
	persons := &Persons{
		Data: []Person{
			{Id: "1", PersonName: "Alice", Salary: struct {
				Value    string "json:\"value\""
				Currency string "json:\"currency\""
			}{Value: "50000", Currency: "USD"}},
			{Id: "2", PersonName: "Bob", Salary: struct {
				Value    string "json:\"value\""
				Currency string "json:\"currency\""
			}{Value: "30000", Currency: "USD"}},
		},
	}

	expected := &Persons{
		Data: []Person{
			{Id: "2", PersonName: "Bob", Salary: struct {
				Value    string "json:\"value\""
				Currency string "json:\"currency\""
			}{Value: "30000", Currency: "USD"}},
			{Id: "1", PersonName: "Alice", Salary: struct {
				Value    string "json:\"value\""
				Currency string "json:\"currency\""
			}{Value: "50000", Currency: "USD"}},
		},
	}

	result := persons.Sort(true)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestSortDescending(t *testing.T) {
	persons := &Persons{
		Data: []Person{
			{Id: "1", PersonName: "Alice", Salary: struct {
				Value    string "json:\"value\""
				Currency string "json:\"currency\""
			}{Value: "50000", Currency: "USD"}},
			{Id: "2", PersonName: "Bob", Salary: struct {
				Value    string "json:\"value\""
				Currency string "json:\"currency\""
			}{Value: "30000", Currency: "USD"}},
		},
	}

	expected := &Persons{
		Data: []Person{
			{Id: "1", PersonName: "Alice", Salary: struct {
				Value    string "json:\"value\""
				Currency string "json:\"currency\""
			}{Value: "50000", Currency: "USD"}},
			{Id: "2", PersonName: "Bob", Salary: struct {
				Value    string "json:\"value\""
				Currency string "json:\"currency\""
			}{Value: "30000", Currency: "USD"}},
		},
	}

	result := persons.Sort(false)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestFilterByCurrency(t *testing.T) {
	persons := &Persons{
		Data: []Person{
			{Id: "1", PersonName: "Alice", Salary: struct {
				Value    string "json:\"value\""
				Currency string "json:\"currency\""
			}{Value: "50000", Currency: "USD"}},
			{Id: "2", PersonName: "Bob", Salary: struct {
				Value    string "json:\"value\""
				Currency string "json:\"currency\""
			}{Value: "30000", Currency: "EUR"}},
		},
	}

	expected := &Persons{
		Data: []Person{
			{Id: "1", PersonName: "Alice", Salary: struct {
				Value    string "json:\"value\""
				Currency string "json:\"currency\""
			}{Value: "50000", Currency: "USD"}},
		},
	}

	result := persons.FilterByCurrency("USD")
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestGroupBySalaryCurrency(t *testing.T) {
	persons := &Persons{
		Data: []Person{
			{Id: "1", PersonName: "Alice", Salary: struct {
				Value    string "json:\"value\""
				Currency string "json:\"currency\""
			}{Value: "50000", Currency: "USD"}},
			{Id: "2", PersonName: "Bob", Salary: struct {
				Value    string "json:\"value\""
				Currency string "json:\"currency\""
			}{Value: "30000", Currency: "USD"}},
			{Id: "3", PersonName: "Charlie", Salary: struct {
				Value    string "json:\"value\""
				Currency string "json:\"currency\""
			}{Value: "40000", Currency: "EUR"}},
		},
	}

	expected := map[string][]Person{
		"USD": {
			{Id: "1", PersonName: "Alice", Salary: struct {
				Value    string "json:\"value\""
				Currency string "json:\"currency\""
			}{Value: "50000", Currency: "USD"}},
			{Id: "2", PersonName: "Bob", Salary: struct {
				Value    string "json:\"value\""
				Currency string "json:\"currency\""
			}{Value: "30000", Currency: "USD"}},
		},
		"EUR": {
			{Id: "3", PersonName: "Charlie", Salary: struct {
				Value    string "json:\"value\""
				Currency string "json:\"currency\""
			}{Value: "40000", Currency: "EUR"}},
		},
	}

	result := persons.GroupBySalaryCurrency()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
