package ticket_item

type PlaceSeat struct {
	ID         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	X          int32  `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
	Y          int32  `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`
	ActivityID int64  `protobuf:"varint,4,opt,name=activity_id,proto3" json:"activity_id,omitempty"`
	Sort       int32  `protobuf:"varint,5,opt,name=sort,proto3" json:"sort,omitempty"`
	State      int32  `protobuf:"varint,6,opt,name=state,proto3" json:"state,omitempty"`
	CreatedTime string `protobuf:"bytes,7,opt,name=created_time,proto3" json:"created_time,omitempty"`
	UpdatedTime string `protobuf:"bytes,8,opt,name=updated_time,proto3" json:"updated_time,omitempty"`
	Status     bool   `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`
}