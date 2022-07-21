package character

import (
	context "context"
	"errors"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const NAME_DATABASE = "user"
const NAME_TABLE = "character"

type characterServer struct {
	GetDataHandler      func(ctx context.Context, req *GetDataRequest) (*GetDataResponse, error)
	SetNameHandler      func(ctx context.Context, req *SetNameRequest) error
	SetMMRHandler       func(ctx context.Context, req *SetMMRRequest) error
	IncrementEXPHandler func(ctx context.Context, req *IncrementEXPRequest) error
}

func NewCharacterServer() *characterServer {
	return &characterServer{}
}

func (server *characterServer) GetData(ctx context.Context, req *GetDataRequest) (*GetDataResponse, error) {
	if server.GetDataHandler == nil {
		return nil, errors.New("GetData handler not implement")
	}
	return server.GetDataHandler(ctx, req)
}
func (server *characterServer) SetName(ctx context.Context, req *SetNameRequest) (*emptypb.Empty, error) {
	if server.SetNameHandler == nil {
		return nil, errors.New("SetName handler not implement")
	}
	return &emptypb.Empty{}, server.SetNameHandler(ctx, req)
}
func (server *characterServer) SetMMR(ctx context.Context, req *SetMMRRequest) (*emptypb.Empty, error) {
	if server.SetMMRHandler == nil {
		return nil, errors.New("SetMMR handler not implement")
	}
	return &emptypb.Empty{}, server.SetMMRHandler(ctx, req)
}
func (server *characterServer) IncrementEXP(ctx context.Context, req *IncrementEXPRequest) (*emptypb.Empty, error) {
	if server.IncrementEXPHandler == nil {
		return nil, errors.New("IncrementEXP handler not implement")
	}
	return &emptypb.Empty{}, server.IncrementEXPHandler(ctx, req)
}
func (server *characterServer) mustEmbedUnimplementedCharacterServer() {}
