package dto

// ========== Smell Type ==========
type SmellTypeRequest struct {
	Name        string `form:"name"  binding:"required"`
	Description string `form:"description" binding:"required"`
}

type CreateOrUpdateSmellTypeResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type SmellTypeResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
