package dto

// ========== Smell Group ==========
type SmellGroupRequest struct {
	Name        string `form:"name"  binding:"required"`
	Description string `form:"description" binding:"required"`
}

type CreateOrUpdateSmellGroupResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type SmellGroupResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
