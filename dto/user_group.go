package dto

//========== Users Group ==========
type UsersGroupRequest struct {
	Name        string `form:"name"  binding:"required"`
	Description string `form:"description"  binding:"required"`
}

type CreateOrUpdateUsersGroupResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UsersGroupResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// //========== System Authorized ==========
// type SystemAuthorizedRequest struct {
// 	Name        string `form:"name"  binding:"required"`
// 	Description string `form:"description" binding:"required"`
// }

// type CreateOrUpdateSystemAuthorizedResponse struct {
// 	ID          uint   `json:"id"`
// 	Name        string `json:"name"`
// 	Description string `form:"description" binding:"required"`
// }

// type SystemAuthorizedResponse struct {
// 	ID          uint   `json:"id"`
// 	Name        string `json:"name"`
// 	Description string `form:"description" binding:"required"`
// }

// //========== System Users ==========
// type SystemUsersRequest struct {
// 	Name        string `form:"name"  binding:"required"`
// 	Description string `form:"description" binding:"required"`
// }

// type CreateOrUpdateSystemUsersResponse struct {
// 	ID          uint   `json:"id"`
// 	Name        string `json:"name"`
// 	Description string `form:"description" binding:"required"`
// }

// type SystemUsersResponse struct {
// 	ID          uint   `json:"id"`
// 	Name        string `json:"name"`
// 	Description string `form:"description" binding:"required"`
// }

// //========== UserApprovedDiscount ==========
// type UserApprovedDiscountRequest struct {
// 	Name string `form:"name"  binding:"required"`
// }

// type CreateOrUpdateUserApprovedDiscountResponse struct {
// 	ID   uint   `json:"id"`
// 	Name string `json:"name"`

// 	SystemUsers []SystemUsersRequest `json:"user"`
// 	Discount    []DiscountRequest    `json:"discount"`
// }

// type SystemUserApprovedDiscountResponse struct {
// 	ID   uint   `json:"id"`
// 	Name string `json:"name"`

// 	SystemUsers []SystemUsersRequest `json:"user"`
// 	Discount    []DiscountRequest    `json:"discount"`
// }

// //========== User Permissions ==========
// type UserPermissionsRequest struct {
// 	Name        string `form:"name"  binding:"required"`
// 	Description string `form:"description"  binding:"required"`
// }

// type CreateOrUpdateUserPermissionsResponse struct {
// 	ID          uint   `json:"id"`
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// }

// type UserPermissionsResponse struct {
// 	ID          uint   `json:"id"`
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// }

// //========== Users Group Authorized ==========
// type UsersGroupAuthorizedRequest struct {
// 	Name string `form:"name"  binding:"required"`
// }

// type CreateOrUpdateUsersGroupAuthorizedResponse struct {
// 	ID   uint   `json:"id"`
// 	Name string `json:"name"`

// 	UsersGroup       []UsersGroupRequest       `json:"usergroup"`
// 	SystemAuthorized []SystemAuthorizedRequest `json:"systemauthorized"`
// 	UserPermissions  []UserPermissionsRequest  `json:"userpermissions"`
// }

// type UsersGroupAuthorizedResponse struct {
// 	ID   uint   `json:"id"`
// 	Name string `json:"name"`

// 	UsersGroup       []UsersGroupRequest       `json:"usergroup"`
// 	SystemAuthorized []SystemAuthorizedRequest `json:"systemauthorized"`
// 	UserPermissions  []UserPermissionsRequest  `json:"userpermissions"`
// }

// //========== Users Group Detail ==========
// type UsersGroupDetailRequest struct {
// 	Name string `form:"name"  binding:"required"`
// }

// type CreateOrUpdateUsersGroupDetailResponse struct {
// 	ID   uint   `json:"id"`
// 	Name string `json:"name"`

// 	SystemUsers []SystemUsersRequest `json:"user"`
// 	UsersGroup  []UsersGroupRequest  `json:"usergroup"`
// }

// type UsersGroupDetailResponse struct {
// 	ID   uint   `json:"id"`
// 	Name string `json:"name"`

// 	SystemUsers []SystemUsersRequest `json:"user"`
// 	UsersGroup  []UsersGroupRequest  `json:"usergroup"`
// }
