package dto

type GetSubscriptionResp struct {
	ID       uint   `json:"id"`
	BlogID   uint   `json:"blog_id"`
	BlogName string `json:"blog_name"`
}

type CreateSubscriptionReq struct {
	BlogID uint `json:"blog_id" binding:"required"`
}

type CreateSubscriptionsReq struct {
	BlogIDs []uint `json:"blog_ids" binding:"required,dive,required"`
}
