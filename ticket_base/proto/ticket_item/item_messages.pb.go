package ticket_item

type Item struct {
	ID               int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ItemName         string  `protobuf:"bytes,2,opt,name=item_name,proto3" json:"item_name,omitempty"`
	AbstractMessage  string  `protobuf:"bytes,3,opt,name=abstract_message,proto3" json:"abstract_message,omitempty"`
	ItemType1ID      int64   `protobuf:"varint,4,opt,name=item_type1_id,proto3" json:"item_type1_id,omitempty"`
	ItemType1Name    string  `protobuf:"bytes,5,opt,name=item_type1_name,proto3" json:"item_type1_name,omitempty"`
	ItemType2ID      int64   `protobuf:"varint,6,opt,name=item_type2_id,proto3" json:"item_type2_id,omitempty"`
	ItemType2Name    string  `protobuf:"bytes,7,opt,name=item_type2_name,proto3" json:"item_type2_name,omitempty"`
	State            int32   `protobuf:"varint,8,opt,name=state,proto3" json:"state,omitempty"`
	BasicDescription string  `protobuf:"bytes,9,opt,name=basic_description,proto3" json:"basic_description,omitempty"`
	ProjectDescription string `protobuf:"bytes,10,opt,name=project_description,proto3" json:"project_description,omitempty"`
	ReminderDescription string `protobuf:"bytes,11,opt,name=reminder_description,proto3" json:"reminder_description,omitempty"`
	Longitude        string  `protobuf:"bytes,12,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude         string  `protobuf:"bytes,13,opt,name=latitude,proto3" json:"latitude,omitempty"`
	AvgScore         float64 `protobuf:"fixed64,14,opt,name=avg_score,proto3" json:"avg_score,omitempty"`
	CommentCount     int32   `protobuf:"varint,15,opt,name=comment_count,proto3" json:"comment_count,omitempty"`
	PlaceID          int64   `protobuf:"varint,16,opt,name=place_id,proto3" json:"place_id,omitempty"`
	StartTime        string  `protobuf:"bytes,17,opt,name=start_time,proto3" json:"start_time,omitempty"`
	EndTime          string  `protobuf:"bytes,18,opt,name=end_time,proto3" json:"end_time,omitempty"`
	CreatedTime      string  `protobuf:"bytes,19,opt,name=created_time,proto3" json:"created_time,omitempty"`
	UpdatedTime      string  `protobuf:"bytes,20,opt,name=updated_time,proto3" json:"updated_time,omitempty"`
	Status           bool    `protobuf:"varint,21,opt,name=status,proto3" json:"status,omitempty"`
}