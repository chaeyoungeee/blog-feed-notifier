package dto

type GetBlogResp struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	MainURL string `json:"main_url"`
}
