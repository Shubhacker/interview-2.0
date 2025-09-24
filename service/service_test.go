package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var i CurrencyInterface

func defineRequired() {
	var d Cache
	d.CountryCache = map[string]Data{}
	i = &d
}

func TestInitialCacheCount(t *testing.T) {
	defineRequired()
	_, present := i.GetCache("India")
	// No data present, returned false
	assert.Equal(t, false, present)

	// Added data in cache
	i.SetCache("India")

	_, present = i.GetCache("India")
	// data present, returned true
	assert.Equal(t, true, present)
}

// Happy path for complete flow
func TestCache(t *testing.T) {

	defineRequired()
	data, err := i.GetCountry(context.Background(), "India")
	assert.Nil(t, err)

	assert.Equal(t, "India", data.Name)
	assert.Equal(t, "New Delhi", data.Capital)
	assert.Equal(t, "Indian rupee", data.Currency)
	assert.Equal(t, 1380004385, data.Population)
}
