package ticket_item

type Place struct {
	ID         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address    string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	AreaID     int64  `protobuf:"varint,4,opt,name=area_id,proto3" json:"area_id,omitempty"`
	AreaName   string `protobuf:"bytes,5,opt,name=area_name,proto3" json:"area_name,omitempty"`
	XLength    int32  `protobuf:"varint,6,opt,name=x_length,proto3" json:"x_length,omitempty"`
	YLength    int32  `protobuf:"varint,7,opt,name=y_length,proto3" json:"y_length,omitempty"`
	Latitude   string `protobuf:"bytes,8,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude  string `protobuf:"bytes,9,opt,name=longitude,proto3" json:"longitude,omitempty"`
	CreatedTime string `protobuf:"bytes,10,opt,name=created_time,proto3" json:"created_time,omitempty"`
	UpdatedTime string `protobuf:"bytes,11,opt,name=updated_time,proto3" json:"updated_time,omitempty"`
	Status     bool   `protobuf:"varint,12,opt,name=status,proto3" json:"status,omitempty"`
}