package providers

type ProviderFunc func() Provider

type ProviderResult struct {
	Name    string `json:"name"`
	Status  bool   `json:"status"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

type IProvider interface {
	Init()
	AddProvider(name string, result ProviderFunc)
}

type Provider struct {
	Name   string
	Result ProviderResult
}

type ProviderResultModel struct {
	Errors  map[string]string `json:"errors,omitempty"`
	Results []Provider        `json:"results"`
}
