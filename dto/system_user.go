package dto

//========== System Users ==========
type SystemUsersLoginRequest struct {
	Email    string `form:"email"  binding:"required"`
	Password string `form:"password" binding:"required"`
}

type SystemUsersRequest struct {
	Email     string `form:"email"  binding:"required"`
	Password  string `form:"password" binding:"required"`
	FirstName string `form:"firstname"  binding:"required"`
	LastName  string `form:"lastname" binding:"required"`
	FullName  string `form:"fullname"  binding:"required"`
	NickName  string `form:"nickname" binding:"required"`
	Mobile    string `form:"mobile"  binding:"required"`
	Status    uint8  `form:"status" binding:"required"`
	Verify    bool   `form:"verify" binding:"required"`

	TitleID uint `form:"titleId" binding:"required"`
	SalesID uint `form:"salesId" binding:"required"`
}

type CreateOrUpdateSystemUsersResponse struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	FullName  string `json:"fullname"`
	NickName  string `json:"nickname"`
	Mobile    string `json:"mobile"`
	Status    uint8  `json:"status"`
	Verify    bool   `json:"verify"`

	TitleID uint `json:"titleId"`
	SalesID uint `json:"salesId"`
}

type SystemUsersResponse struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	FullName  string `json:"fullname"`
	NickName  string `json:"nickname"`
	Mobile    string `json:"mobile"`
	Status    uint8  `json:"status"`
	Verify    bool   `json:"verify"`

	Title TitleResponse     `json:"title"`
	Sales SalesOnlyResponse `json:"sales"`
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

// //========== Users Group ==========
// type UsersGroupRequest struct {
// 	Name        string `form:"name"  binding:"required"`
// 	Description string `form:"description"  binding:"required"`
// }

// type CreateOrUpdateUsersGroupResponse struct {
// 	ID          uint   `json:"id"`
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// }

// type UsersGroupResponse struct {
// 	ID          uint   `json:"id"`
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
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
