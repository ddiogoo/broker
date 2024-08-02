package model

type Key struct {
	Email  string `json:"email"`
	ApiKey string `json:"apiKey"`
	Routes string `json:"permissions"`
}

// NewKey create an instance of Key struct.
func NewKey(email, apiKey, routes string) *Key {
	return &Key{
		Email:  email,
		ApiKey: apiKey,
		Routes: routes,
	}
}
