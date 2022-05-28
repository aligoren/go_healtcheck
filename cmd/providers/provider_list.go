package providers

import (
	"sync"
)

var providerList []func() ProviderResult

func Init() {

	providerList = make([]func() ProviderResult, 0)
}

func (p Provider) AddProvider(result func() ProviderResult) {

	providerList = append(providerList, result)
}

func GetProviders() ProviderResultModel {

	providerResult := ProviderResultModel{
		Errors:  make(map[string]string),
		Results: make([]Provider, 0),
	}

	results := make(chan Provider, 0)

	mu := &sync.Mutex{}

	go func() {
		wg := sync.WaitGroup{}

		for _, provider := range providerList {

			wg.Add(1)

			go func() {
				defer wg.Done()

				result := provider()

				mu.Lock()
				results <- Provider{
					Name:   result.Name,
					Result: result,
				}

				if !result.Status {
					providerResult.Errors[result.Name] = result.Error
				}
				mu.Unlock()

			}()
			wg.Wait()
		}

		close(results)
	}()

	for result := range results {
		providerResult.Results = append(providerResult.Results, result)
	}

	return providerResult
}
