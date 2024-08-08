package model

type RoutePermission struct {
	Route string `json:"route"`
}

// NewRoutePermission create an instance of RoutePermission struct.
func NewRoutePermission(route string) *RoutePermission {
	return &RoutePermission{
		Route: route,
	}
}
