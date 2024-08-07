package model

type Key struct {
	Email  string `json:"email"`
	ApiKey string `json:"apiKey"`
	Route  string `json:"route"`
}

// NewKey create an instance of Key struct.
func NewKey(email, apiKey, route string) *Key {
	return &Key{
		Email:  email,
		ApiKey: apiKey,
		Route:  route,
	}
}
