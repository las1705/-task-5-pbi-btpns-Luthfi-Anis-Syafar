package request_model

type UserRequest struct {
	Username string `json:"username" form:"username" binding:"required" `
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
