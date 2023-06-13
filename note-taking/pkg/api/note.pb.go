package api

import (
	context "context"
	fmt "fmt"

	grpc "google.golang.org/grpc"
)

type NoteClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type noteClient struct {
	cc *grpc.ClientConn
}

func NewNoteClient(cc *grpc.ClientConn) NoteClient {
	return &noteClient{cc}
}

func (c *noteClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/api.Note/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/api.Note/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/api.Note/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/api.Note/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type NoteServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterNoteServer(s *grpc.Server, srv NoteServer) {
	s.RegisterService(&_Note_serviceDesc, srv)
}

var _Note_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Note",
	HandlerType: (*NoteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Note_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Note_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Note_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Note_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/note.proto",
}

func _Note_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Note/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

type AddRequest struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return fmt.Sprintf("AddRequest{%v}", *m) }
func (*AddRequest) ProtoMessage()    {}

func (m *AddRequest) GetTitle() string {
	return m.Title
}

func (m *AddRequest) GetDescription() string {
	return m.Description
}

type AddResponse struct {
	NoteId int32 `protobuf:"varint,1,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return fmt.Sprintf("AddResponse{%v}", *m) }
func (*AddResponse) ProtoMessage()    {}

func (m *AddResponse) GetNoteId() int32 {
	return m.NoteId
}

func _Note_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Note/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

type GetRequest struct {
	NoteID string `protobuf:"bytes,1,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return fmt.Sprintf("GetRequest{%v}", *m) }
func (*GetRequest) ProtoMessage()    {}

type GetResponse struct {
	Note *Note `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return fmt.Sprintf("GetResponse{%v}", *m) }
func (*GetResponse) ProtoMessage()    {}

func (m *GetResponse) GetNote() *Note {
	return m.Note
}

func _Note_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Note/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

type UpdateRequest struct {
	NoteID   string `protobuf:"bytes,1,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
	NewTitle string `protobuf:"bytes,2,opt,name=new_title,json=newTitle,proto3" json:"new_title,omitempty"`
	NewBody  string `protobuf:"bytes,3,opt,name=new_body,json=newBody,proto3" json:"new_body,omitempty"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return fmt.Sprintf("UpdateRequest{%v}", *m) }
func (*UpdateRequest) ProtoMessage()    {}

type UpdateResponse struct {
	Note *Note `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
}

func (m *UpdateResponse) GetNote() *Note {
	return m.Note
}

func (m *UpdateResponse) Reset() {
	*m = UpdateResponse{}
}

func (m *UpdateResponse) String() string {
	return fmt.Sprintf("UpdateResponse{%v}", *m)
}

func (*UpdateResponse) ProtoMessage() {}

func _Note_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Note/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

type DeleteRequest struct {
	NoteID string `protobuf:"bytes,1,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return fmt.Sprintf("DeleteRequest{%v}", *m) }
func (*DeleteRequest) ProtoMessage()    {}

type DeleteResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (m *DeleteResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *DeleteResponse) Reset() {
	*m = DeleteResponse{}
}

func (m *DeleteResponse) String() string {
	return fmt.Sprintf("DeleteResponse{%v}", *m)
}

func (*DeleteResponse) ProtoMessage() {}

type Note struct {
	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *Note) Reset()         { *m = Note{} }
func (m *Note) String() string { return fmt.Sprintf("Note{%v}", *m) }
func (*Note) ProtoMessage()    {}

func (m *Note) GetId() int32 {
	return m.Id
}

func (m *Note) GetTitle() string {
	return m.Title
}

func (m *Note) GetDescription() string {
	return m.Description
}
