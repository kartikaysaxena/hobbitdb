package api

import (
	"strconv"

	"github.com/kartikayaxena/hobbitdb/hobbit"
)

type FilterMap struct {
	filters map[string]hobbit.Map
}

func NewFilterMap() *FilterMap {
	filters := make(map[string]hobbit.Map)
	filters[hobbit.FilterTypeEQ] = hobbit.Map{}
	return &FilterMap{
		filters: filters,
	}
}

func (f *FilterMap) Get(filterType string) hobbit.Map {
	val, ok := f.filters[filterType]
	if !ok {
		return hobbit.Map{}
	}
	return val
}

func (f *FilterMap) Add(filterType, k string, v string) {
	if _, ok := f.filters[filterType]; !ok {
		return
	}
	f.filters[filterType][k] = ensureCorrectTypeFromString(v)
}

func ensureCorrectTypeFromString(v string) any {
	switch {
	case v == "true":
		return true
	case v == "false":
		return false
	case isInteger(v):
		val, _ := strconv.Atoi(v)
		return val
	case isFloat(v):
		val, _ := strconv.ParseFloat(v, 64)
		return val
	default:
		return v
	}
}

func isFloat(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func isInteger(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
