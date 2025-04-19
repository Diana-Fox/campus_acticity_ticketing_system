package ticket_item

import (
	context "context"
)

type TicketItemServiceClient interface {
	GetItemById(ctx context.Context, in *GetItemByIdRequest, opts ...grpc.CallOption) (*GetItemByIdResponse, error)
	GetItemCommentsByItemId(ctx context.Context, in *GetItemCommentsByItemIdRequest, opts ...grpc.CallOption) (*GetItemCommentsByItemIdResponse, error)
	GetPlaceById(ctx context.Context, in *GetPlaceByIdRequest, opts ...grpc.CallOption) (*GetPlaceByIdResponse, error)
	GetPlaceSeatsByActivityId(ctx context.Context, in *GetPlaceSeatsByActivityIdRequest, opts ...grpc.CallOption) (*GetPlaceSeatsByActivityIdResponse, error)
}

type TicketItemServiceServer interface {
	GetItemById(context.Context, *GetItemByIdRequest) (*GetItemByIdResponse, error)
	GetItemCommentsByItemId(context.Context, *GetItemCommentsByItemIdRequest) (*GetItemCommentsByItemIdResponse, error)
	GetPlaceById(context.Context, *GetPlaceByIdRequest) (*GetPlaceByIdResponse, error)
	GetPlaceSeatsByActivityId(context.Context, *GetPlaceSeatsByActivityIdRequest) (*GetPlaceSeatsByActivityIdResponse, error)
}

func RegisterTicketItemServiceServer(s *grpc.Server, srv TicketItemServiceServer) {
	s.RegisterService(&_TicketItemService_serviceDesc, srv)
}

type GetItemByIdRequest struct {
	ID int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

type GetItemByIdResponse struct {
	Item *Item `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

type GetItemCommentsByItemIdRequest struct {
	ItemID int64 `protobuf:"varint,1,opt,name=item_id,proto3" json:"item_id,omitempty"`
}

type GetItemCommentsByItemIdResponse struct {
	ItemComments []*ItemComment `protobuf:"bytes,1,rep,name=item_comments,proto3" json:"item_comments,omitempty"`
}

type GetPlaceByIdRequest struct {
	ID int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

type GetPlaceByIdResponse struct {
	Place *Place `protobuf:"bytes,1,opt,name=place,proto3" json:"place,omitempty"`
}

type GetPlaceSeatsByActivityIdRequest struct {
	ActivityID int64 `protobuf:"varint,1,opt,name=activity_id,proto3" json:"activity_id,omitempty"`
}

type GetPlaceSeatsByActivityIdResponse struct {
	PlaceSeats []*PlaceSeat `protobuf:"bytes,1,rep,name=place_seats,proto3" json:"place_seats,omitempty"`
}