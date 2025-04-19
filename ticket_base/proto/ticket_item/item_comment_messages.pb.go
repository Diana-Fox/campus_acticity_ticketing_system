package ticket_item

type ItemComment struct {
	ID         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ItemID     int64  `protobuf:"varint,2,opt,name=item_id,proto3" json:"item_id,omitempty"`
	UserID     int64  `protobuf:"varint,3,opt,name=user_id,proto3" json:"user_id,omitempty"`
	Score      int32  `protobuf:"varint,4,opt,name=score,proto3" json:"score,omitempty"`
	Content    string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	CreatedTime string `protobuf:"bytes,6,opt,name=created_time,proto3" json:"created_time,omitempty"`
	UpdatedTime string `protobuf:"bytes,7,opt,name=updated_time,proto3" json:"updated_time,omitempty"`
}