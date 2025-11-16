package dto

type GetSubscriptionResp struct {
	ID       uint   `json:"id"`
	BlogID   uint   `json:"blog_id"`
	BlogName string `json:"blog_name"`
}
