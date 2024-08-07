package dto

type CheckPermissionDto struct {
	ApiKey string `json:"apiKey" binding:"required"`
	Route  string `json:"route" binding:"required"`
}
