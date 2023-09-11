package models

type Product struct {
	UserID         string         `json:"user_id"`
	UserName       string         `json:"user_name"`
	UserProfile    string         `json:"user_profile"`
	ProductID      string         `json:"product_id"`
	ProductProfile string         `json:"product_profile"`
	ProductName    string         `json:"product_name"`
	ProductDetail  string         `json:"product_detail"`
	OldPrice       string         `json:"old_price"`
	NewPrice       string         `json:"new_price"`
	LikesCount     int            `json:"likes_count"`
	CommentsCount  int            `json:"comments_count"`
	Comments       []WatchComment `json:"comments"`
}

type WatchComment struct {
	UserName string   `json:"User-Name"`
	Comment  string   `json:"Comment"`
	Replies  []string `json:"Replies"`
}
