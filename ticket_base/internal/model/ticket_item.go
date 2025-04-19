package model

import (
	"time"
)

type Item struct {
	ID                  int64     `json:"id"`
	ItemName            string    `json:"itemName"`
	AbstractMessage     string    `json:"abstractMessage"`
	ItemType1Id         int64     `json:"itemType1Id"`
	ItemType1Name       string    `json:"itemType1Name"`
	ItemType2Id         int64     `json:"itemType2Id"`
	ItemType2Name       string    `json:"itemType2Name"`
	State               int       `json:"state"`
	BasicDescription    string    `json:"basicDescription"`
	ProjectDescription  string    `json:"projectDescription"`
	ReminderDescription string    `json:"reminderDescription"`
	Longitude           string    `json:"longitude"`
	Latitude            string    `json:"latitude"`
	AvgScore            float64   `json:"avgScore"`
	CommentCount        int       `json:"commentCount"`
	PlaceId             int64     `json:"placeId"`
	StartTime           time.Time `json:"startTime"`
	EndTime             time.Time `json:"endTime"`
	CreatedTime         time.Time `json:"createdTime"`
	UpdatedTime         time.Time `json:"updatedTime"`
	Status              bool      `json:"status"`
}

type ItemComment struct {
	ID          int64     `json:"id"`
	ItemId      int64     `json:"itemId"`
	UserId      int64     `json:"userId"`
	Score       int       `json:"score"`
	Content     string    `json:"content"`
	CreatedTime time.Time `json:"createdTime"`
	UpdatedTime time.Time `json:"updatedTime"`
}

type Place struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	AreaId      int64     `json:"areaId"`
	AreaName    string    `json:"areaName"`
	XLength     int       `json:"xLength"`
	YLength     int       `json:"yLength"`
	Latitude    string    `json:"latitude"`
	Longitude   string    `json:"longitude"`
	CreatedTime time.Time `json:"createdTime"`
	UpdatedTime time.Time `json:"updatedTime"`
	Status      bool      `json:"status"`
}

type PlaceSeat struct {
	ID          int64     `json:"id"`
	X           int       `json:"x"`
	Y           int       `json:"y"`
	ActivityId  int64     `json:"activityId"`
	Sort        int       `json:"sort"`
	State       int       `json:"state"`
	CreatedTime time.Time `json:"createdTime"`
	UpdatedTime time.Time `json:"updatedTime"`
	Status      bool      `json:"status"`
}
