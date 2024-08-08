package model

type Key struct {
	Email  string            `json:"email"`
	ApiKey string            `json:"apiKey"`
	Routes []RoutePermission `json:"routes"`
}

// NewKey create an instance of Key struct.
func NewKey(email, apiKey string, routes []RoutePermission) *Key {
	return &Key{
		Email:  email,
		ApiKey: apiKey,
		Routes: routes,
	}
}
