package responses_model

import "time"

type UserResponse struct {
	ID       *int       `json:"id"`
	Username *string    `json:"username"`
	Email    *string    `json:"email"`
	Password *string    `json:"password"`
	CreateAt *string    `json:"create_at"`
	UpdateAt *time.Time `json:"update_at"`
}
