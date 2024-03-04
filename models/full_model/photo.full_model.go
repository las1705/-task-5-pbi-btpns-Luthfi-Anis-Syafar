package full_model

type Photo struct {
	ID       *int    `json:"id"`
	Title    *string `json:"title"`
	Caption  *string `json:"caption"`
	PhotoUrl *string `json:"photo_url"`
	UserId   *int    `json:"user_id"`
}
