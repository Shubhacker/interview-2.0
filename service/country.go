package service

import "context"

// service package will have access to all technologies, such as: DB, Cache, logger,  etc

type CurrencyInterface interface {
	GetCountry(ctx context.Context, name string) (Data, error)
	GetCache(name string) (Data, bool)
	SetCache(key string) (Data, error)
}

type Data struct {
	Name       string
	Capital    string
	Currency   string
	Population int
}

type Cache struct {
	CountryCache map[string]Data
}

func (c *Cache) GetCache(name string) (Data, bool) {
	data, ok := c.CountryCache[name]
	if ok {
		// data present, return from cache
		return data, true
	}

	return data, false
}

func (c *Cache) SetCache(key string) (Data, error) {
	externalData, err := ExternalAPI(key)
	if err != nil {
		return Data{}, err
	}

	for _, dt := range externalData {
		c.CountryCache[dt.Name] = dt
	}

	// verifying if data added correctly
	dt, present := c.GetCache(key)
	if !present {
		// something went wrong, retrying to add data, later can improve retry mechanism
		c.SetCache(key)
	}

	return dt, nil
}

// Later can use server timing to check performance improvement while using cache
func (c *Cache) GetCountry(ctx context.Context, name string) (Data, error) {

	data, dataPresent := c.GetCache(name)
	if dataPresent {
		// dataPresent, fetching record from cache
		return data, nil
	}

	// Adding data in cache
	dt, err := c.SetCache(name)
	if err != nil {
		return Data{}, err
	}
	return dt, nil
}
