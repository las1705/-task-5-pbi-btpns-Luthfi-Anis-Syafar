package request_model

type PhotoRequest struct {
	Title   string `json:"title" form:"title" binding:"required" `
	Caption string `json:"caption" form:"caption" binding:"required"`
}
