package model

import "time"

type Item struct {
	ID               int64     `json:"id"`
	ItemName         string    `json:"item_name"`
	AbstractMessage  string    `json:"abstract_message"`
	ItemType1ID      int64     `json:"item_type1_id"`
	ItemType1Name    string    `json:"item_type1_name"`
	ItemType2ID      int64     `json:"item_type2_id"`
	ItemType2Name    string    `json:"item_type2_name"`
	State            int32     `json:"state"`
	BasicDescription string    `json:"basic_description"`
	ProjectDescription string  `json:"project_description"`
	ReminderDescription string `json:"reminder_description"`
	Longitude        string    `json:"longitude"`
	Latitude         string    `json:"latitude"`
	AvgScore         float64   `json:"avg_score"`
	CommentCount     int32     `json:"comment_count"`
	PlaceID          int64     `json:"place_id"`
	StartTime        time.Time `json:"start_time"`
	EndTime          time.Time `json:"end_time"`
	CreatedTime      time.Time `json:"created_time"`
	UpdatedTime      time.Time `json:"updated_time"`
	Status           bool      `json:"status"`
}