package main


type Watch struct {
	UserID				string	`json:"User_Id"`
	UserName 			string	`json:"User-Name"`
	UserProfile			string	`json:"User-Profile"`
	ProductID			string	`json:"Product-ID"`
	ProductProfile		string	`json:"Product_Link"`
	WatchName 			string	`json:"Product-Name"`
	ProductDetail 		string	`json:"Product-Detail"`
	OldPrice 			string	`json:"Old-Price"`
	NewPrice 			string	`json:"New-Price"`
	LikesCount			int 	`json:"Like-Count"`
	CommentsCount		int 	`json:"Comments-Count"`
	Comments			[]WatchComment `json:"Comments"`

}
type WatchComment struct {
	UserName 			string	`json:"User-Name"`
	Comment 			  string
	Replies 			[]SubComment `json:"Replies"`
}
type SubComment struct {
	SubComment 			  	[]string
}

