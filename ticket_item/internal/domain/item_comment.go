package domain

type ItemComment struct {
	Id      int64  `json:"id,omitempty"`
	ItemId  int64  `json:"itemId,omitempty"`
	UserId  int64  `json:"userId,omitempty"`
	Score   int64  `json:"score,omitempty"`
	Content string `json:"content,omitempty"`
}
