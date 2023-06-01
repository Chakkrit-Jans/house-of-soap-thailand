package dto

// ========== Product Formula ==========
type ProductFormulaRequest struct {
	Code            string  `form:"code" binding:"required"`
	Name            string  `form:"name" binding:"required"`
	Properties      string  `form:"properties"`
	ActiveIngedient string  `form:"activeIngedient"`
	UnitCost        float32 `form:"unitcost"`
	SalePrice       float32 `form:"saleprice"`

	ProductID        uint `form:"productId"`
	ProductTypeID    uint `form:"producttypeId"`
	ProductTextureID uint `form:"producttextureId"`
	SmellID          uint `form:"smellId"`
	Color1ID         uint `form:"color1Id"`
	Color2ID         uint `form:"color2Id"`
	Color3ID         uint `form:"clolor3Id"`
	TestSizeID       uint `form:"testsizeId"`
	ClaimID          uint `form:"claimId"`
	HowtoID          uint `form:"howtoId"`
}

type CreateOrUpdateProductFormulaResponse struct {
	ID              uint    `json:"id"`
	Code            string  `json:"code"`
	Name            string  `json:"name"`
	Properties      string  `json:"properties"`
	ActiveIngedient string  `json:"activeIngedient"`
	UnitCost        float32 `json:"unitcost"`
	SalePrice       float32 `json:"saleprice"`
	Image           string  `json:"image"`

	ProductID        uint `json:"productId"`
	ProductTypeID    uint `json:"producttypeId"`
	ProductTextureID uint `json:"producttextureId"`
	SmellID          uint `json:"smellId"`
	Color1ID         uint `json:"color1Id"`
	Color2ID         uint `json:"color2Id"`
	Color3ID         uint `json:"clolor3Id"`
	TestSizeID       uint `json:"testsizeId"`
	ClaimID          uint `json:"claimId"`
	HowtoID          uint `json:"howtoId"`
}

type ProductFormulaResponse struct {
	ID              uint    `json:"id"`
	Code            string  `json:"code"`
	Name            string  `json:"name"`
	Properties      string  `json:"properties"`
	ActiveIngedient string  `json:"activeIngedient"`
	UnitCost        float32 `json:"unitcost"`
	SalePrice       float32 `json:"saleprice"`
	Image           string  `json:"image"`

	Product        ProductResponse         `json:"product"`
	ProductType    ProductTypeOnlyResponse `json:"producttype"`
	ProductTexture ProductTextureResponse  `json:"producttexture"`
	Smell          SmellOnlyResponse       `json:"smell"`
	Color1         Color1Response          `json:"color1"`
	Color2         Color2Response          `json:"color2"`
	Color3         Color3Response          `json:"clolor3"`
	TestSize       TestSizeResponse        `json:"testsize"`
	Claim          ClaimResponse           `json:"claim"`
	Howto          HowtoResponse           `json:"howto"`
}
