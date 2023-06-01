package dto

// ========== Smell ==========
type SmellRequest struct {
	Name        string `form:"name"  binding:"required"`
	Description string `form:"description"`

	SmellExtractsID uint `form:"smellextractsId"`
	SmellTypeID     uint `form:"smelltypeId"`
	SmellGroupID    uint `form:"smellgroupId"`
}

type CreateOrUpdateSmellResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	SmellExtractsID uint `json:"smellextractsId"`
	SmellTypeID     uint `json:"smelltypeId"`
	SmellGroupID    uint `json:"smellgroupId"`
}

type SmellResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	SmellExtracts SmellExtractsResponse `json:"smellextracts"`
	SmellType     SmellTypeResponse     `json:"smelltype"`
	SmellGroup    SmellGroupResponse    `json:"smellgroup"`
}

type SmellOnlyResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
