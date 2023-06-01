package dto

type AppConfigRequest struct {
	ProductionCostPerItem float32 `form:"productitemcost" binding:"required"`
	ProductionCostPerUnit float32 `form:"productunitcost" binding:"required"`
	PercentageProfit      float32 `form:"profit" binding:"required"`
}

type CreateOrUpdateAppConfigResponse struct {
	ID                    uint    `json:"id"`
	ProductionCostPerItem float32 `json:"productitemcost"`
	ProductionCostPerUnit float32 `json:"productunitcost"`
	PercentageProfit      float32 `json:"profit"`
}

type AppConfigResponse struct {
	ID                    uint    `json:"id"`
	ProductionCostPerItem float32 `json:"productitemcost"`
	ProductionCostPerUnit float32 `json:"productunitcost"`
	PercentageProfit      float32 `json:"profit"`
}
