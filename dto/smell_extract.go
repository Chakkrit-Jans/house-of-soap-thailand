package dto

// ========== Smell Extracts ==========
type SmellExtractsRequest struct {
	Name        string `form:"name"  binding:"required"`
	Description string `form:"description" binding:"required"`
}

type CreateOrUpdateSmellExtractsResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type SmellExtractsResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
