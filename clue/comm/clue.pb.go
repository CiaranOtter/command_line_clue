// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: proto/clue.proto

package comm

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CommandType int32

const (
	CommandType_PICK_CHAR CommandType = 0
)

// Enum value maps for CommandType.
var (
	CommandType_name = map[int32]string{
		0: "PICK_CHAR",
	}
	CommandType_value = map[string]int32{
		"PICK_CHAR": 0,
	}
)

func (x CommandType) Enum() *CommandType {
	p := new(CommandType)
	*p = x
	return p
}

func (x CommandType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommandType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_clue_proto_enumTypes[0].Descriptor()
}

func (CommandType) Type() protoreflect.EnumType {
	return &file_proto_clue_proto_enumTypes[0]
}

func (x CommandType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommandType.Descriptor instead.
func (CommandType) EnumDescriptor() ([]byte, []int) {
	return file_proto_clue_proto_rawDescGZIP(), []int{0}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*Message_Con
	//	*Message_Opp
	//	*Message_Com
	//	*Message_End
	//	*Message_SetChar
	//	*Message_Start
	//	*Message_RAns
	//	*Message_Cards
	Data isMessage_Data `protobuf_oneof:"data"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto_clue_proto_rawDescGZIP(), []int{0}
}

func (m *Message) GetData() isMessage_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *Message) GetCon() *Connect {
	if x, ok := x.GetData().(*Message_Con); ok {
		return x.Con
	}
	return nil
}

func (x *Message) GetOpp() *Opponent {
	if x, ok := x.GetData().(*Message_Opp); ok {
		return x.Opp
	}
	return nil
}

func (x *Message) GetCom() *Command {
	if x, ok := x.GetData().(*Message_Com); ok {
		return x.Com
	}
	return nil
}

func (x *Message) GetEnd() *GameEnd {
	if x, ok := x.GetData().(*Message_End); ok {
		return x.End
	}
	return nil
}

func (x *Message) GetSetChar() *SetPlayer {
	if x, ok := x.GetData().(*Message_SetChar); ok {
		return x.SetChar
	}
	return nil
}

func (x *Message) GetStart() *StartGame {
	if x, ok := x.GetData().(*Message_Start); ok {
		return x.Start
	}
	return nil
}

func (x *Message) GetRAns() *RemoveAnswer {
	if x, ok := x.GetData().(*Message_RAns); ok {
		return x.RAns
	}
	return nil
}

func (x *Message) GetCards() *Cards {
	if x, ok := x.GetData().(*Message_Cards); ok {
		return x.Cards
	}
	return nil
}

type isMessage_Data interface {
	isMessage_Data()
}

type Message_Con struct {
	Con *Connect `protobuf:"bytes,1,opt,name=con,proto3,oneof"`
}

type Message_Opp struct {
	Opp *Opponent `protobuf:"bytes,2,opt,name=opp,proto3,oneof"`
}

type Message_Com struct {
	Com *Command `protobuf:"bytes,3,opt,name=com,proto3,oneof"`
}

type Message_End struct {
	End *GameEnd `protobuf:"bytes,4,opt,name=end,proto3,oneof"`
}

type Message_SetChar struct {
	SetChar *SetPlayer `protobuf:"bytes,5,opt,name=setChar,proto3,oneof"`
}

type Message_Start struct {
	Start *StartGame `protobuf:"bytes,6,opt,name=start,proto3,oneof"`
}

type Message_RAns struct {
	RAns *RemoveAnswer `protobuf:"bytes,7,opt,name=rAns,proto3,oneof"`
}

type Message_Cards struct {
	Cards *Cards `protobuf:"bytes,8,opt,name=cards,proto3,oneof"`
}

func (*Message_Con) isMessage_Data() {}

func (*Message_Opp) isMessage_Data() {}

func (*Message_Com) isMessage_Data() {}

func (*Message_End) isMessage_Data() {}

func (*Message_SetChar) isMessage_Data() {}

func (*Message_Start) isMessage_Data() {}

func (*Message_RAns) isMessage_Data() {}

func (*Message_Cards) isMessage_Data() {}

type RemoveAnswer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room int32 `protobuf:"varint,1,opt,name=Room,proto3" json:"Room,omitempty"`
	Char int32 `protobuf:"varint,2,opt,name=Char,proto3" json:"Char,omitempty"`
	Weap int32 `protobuf:"varint,3,opt,name=Weap,proto3" json:"Weap,omitempty"`
}

func (x *RemoveAnswer) Reset() {
	*x = RemoveAnswer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveAnswer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveAnswer) ProtoMessage() {}

func (x *RemoveAnswer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveAnswer.ProtoReflect.Descriptor instead.
func (*RemoveAnswer) Descriptor() ([]byte, []int) {
	return file_proto_clue_proto_rawDescGZIP(), []int{1}
}

func (x *RemoveAnswer) GetRoom() int32 {
	if x != nil {
		return x.Room
	}
	return 0
}

func (x *RemoveAnswer) GetChar() int32 {
	if x != nil {
		return x.Char
	}
	return 0
}

func (x *RemoveAnswer) GetWeap() int32 {
	if x != nil {
		return x.Weap
	}
	return 0
}

type Cards struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index []int32 `protobuf:"varint,1,rep,packed,name=index,proto3" json:"index,omitempty"`
}

func (x *Cards) Reset() {
	*x = Cards{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clue_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cards) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cards) ProtoMessage() {}

func (x *Cards) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clue_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cards.ProtoReflect.Descriptor instead.
func (*Cards) Descriptor() ([]byte, []int) {
	return file_proto_clue_proto_rawDescGZIP(), []int{2}
}

func (x *Cards) GetIndex() []int32 {
	if x != nil {
		return x.Index
	}
	return nil
}

type Opponent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerName string `protobuf:"bytes,1,opt,name=PlayerName,proto3" json:"PlayerName,omitempty"`
	// Types that are assignable to Data:
	//
	//	*Opponent_Con
	//	*Opponent_SetChar
	Data isOpponent_Data `protobuf_oneof:"data"`
}

func (x *Opponent) Reset() {
	*x = Opponent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clue_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Opponent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Opponent) ProtoMessage() {}

func (x *Opponent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clue_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Opponent.ProtoReflect.Descriptor instead.
func (*Opponent) Descriptor() ([]byte, []int) {
	return file_proto_clue_proto_rawDescGZIP(), []int{3}
}

func (x *Opponent) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (m *Opponent) GetData() isOpponent_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *Opponent) GetCon() *Connect {
	if x, ok := x.GetData().(*Opponent_Con); ok {
		return x.Con
	}
	return nil
}

func (x *Opponent) GetSetChar() *SetPlayer {
	if x, ok := x.GetData().(*Opponent_SetChar); ok {
		return x.SetChar
	}
	return nil
}

type isOpponent_Data interface {
	isOpponent_Data()
}

type Opponent_Con struct {
	Con *Connect `protobuf:"bytes,2,opt,name=con,proto3,oneof"`
}

type Opponent_SetChar struct {
	SetChar *SetPlayer `protobuf:"bytes,3,opt,name=setChar,proto3,oneof"`
}

func (*Opponent_Con) isOpponent_Data() {}

func (*Opponent_SetChar) isOpponent_Data() {}

type GameEnd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GameEnd) Reset() {
	*x = GameEnd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clue_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameEnd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameEnd) ProtoMessage() {}

func (x *GameEnd) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clue_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameEnd.ProtoReflect.Descriptor instead.
func (*GameEnd) Descriptor() ([]byte, []int) {
	return file_proto_clue_proto_rawDescGZIP(), []int{4}
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type CommandType `protobuf:"varint,1,opt,name=type,proto3,enum=CommandType" json:"type,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clue_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clue_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_proto_clue_proto_rawDescGZIP(), []int{5}
}

func (x *Command) GetType() CommandType {
	if x != nil {
		return x.Type
	}
	return CommandType_PICK_CHAR
}

type Connect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerName string `protobuf:"bytes,1,opt,name=playerName,proto3" json:"playerName,omitempty"`
}

func (x *Connect) Reset() {
	*x = Connect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clue_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connect) ProtoMessage() {}

func (x *Connect) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clue_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connect.ProtoReflect.Descriptor instead.
func (*Connect) Descriptor() ([]byte, []int) {
	return file_proto_clue_proto_rawDescGZIP(), []int{6}
}

func (x *Connect) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

type SetPlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CharacterName string `protobuf:"bytes,1,opt,name=CharacterName,proto3" json:"CharacterName,omitempty"`
}

func (x *SetPlayer) Reset() {
	*x = SetPlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clue_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPlayer) ProtoMessage() {}

func (x *SetPlayer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clue_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPlayer.ProtoReflect.Descriptor instead.
func (*SetPlayer) Descriptor() ([]byte, []int) {
	return file_proto_clue_proto_rawDescGZIP(), []int{7}
}

func (x *SetPlayer) GetCharacterName() string {
	if x != nil {
		return x.CharacterName
	}
	return ""
}

type StartGame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartGame) Reset() {
	*x = StartGame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clue_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartGame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartGame) ProtoMessage() {}

func (x *StartGame) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clue_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartGame.ProtoReflect.Descriptor instead.
func (*StartGame) Descriptor() ([]byte, []int) {
	return file_proto_clue_proto_rawDescGZIP(), []int{8}
}

var File_proto_clue_proto protoreflect.FileDescriptor

var file_proto_clue_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9b, 0x02, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c,
	0x0a, 0x03, 0x63, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x48, 0x00, 0x52, 0x03, 0x63, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x03,
	0x6f, 0x70, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4f, 0x70, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x03, 0x6f, 0x70, 0x70, 0x12, 0x1c, 0x0a, 0x03, 0x63,
	0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x48, 0x00, 0x52, 0x03, 0x63, 0x6f, 0x6d, 0x12, 0x1c, 0x0a, 0x03, 0x65, 0x6e, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x64,
	0x48, 0x00, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x26, 0x0a, 0x07, 0x73, 0x65, 0x74, 0x43, 0x68,
	0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x48, 0x00, 0x52, 0x07, 0x73, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x12,
	0x22, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x48, 0x00, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x72, 0x41, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x48, 0x00, 0x52, 0x04, 0x72, 0x41, 0x6e, 0x73, 0x12, 0x1e, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x64,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x73, 0x48,
	0x00, 0x52, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x4a, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x43, 0x68, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x57, 0x65, 0x61, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x57, 0x65, 0x61, 0x70, 0x22, 0x1d, 0x0a, 0x05,
	0x43, 0x61, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x78, 0x0a, 0x08, 0x4f,
	0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x03, 0x63, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x48, 0x00,
	0x52, 0x03, 0x63, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x07, 0x73, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x48, 0x00, 0x52, 0x07, 0x73, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x42, 0x06, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x09, 0x0a, 0x07, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x64,
	0x22, 0x2b, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x20, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x29, 0x0a,
	0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x31, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x43, 0x68,
	0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x0b, 0x0a, 0x09, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x2a, 0x1c, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x49, 0x43, 0x4b, 0x5f,
	0x43, 0x48, 0x41, 0x52, 0x10, 0x00, 0x32, 0x33, 0x0a, 0x0b, 0x43, 0x6c, 0x75, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0a, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x08, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x08, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x63,
	0x6c, 0x75, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_clue_proto_rawDescOnce sync.Once
	file_proto_clue_proto_rawDescData = file_proto_clue_proto_rawDesc
)

func file_proto_clue_proto_rawDescGZIP() []byte {
	file_proto_clue_proto_rawDescOnce.Do(func() {
		file_proto_clue_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_clue_proto_rawDescData)
	})
	return file_proto_clue_proto_rawDescData
}

var file_proto_clue_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_clue_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_clue_proto_goTypes = []interface{}{
	(CommandType)(0),     // 0: CommandType
	(*Message)(nil),      // 1: Message
	(*RemoveAnswer)(nil), // 2: RemoveAnswer
	(*Cards)(nil),        // 3: Cards
	(*Opponent)(nil),     // 4: Opponent
	(*GameEnd)(nil),      // 5: GameEnd
	(*Command)(nil),      // 6: Command
	(*Connect)(nil),      // 7: Connect
	(*SetPlayer)(nil),    // 8: SetPlayer
	(*StartGame)(nil),    // 9: StartGame
}
var file_proto_clue_proto_depIdxs = []int32{
	7,  // 0: Message.con:type_name -> Connect
	4,  // 1: Message.opp:type_name -> Opponent
	6,  // 2: Message.com:type_name -> Command
	5,  // 3: Message.end:type_name -> GameEnd
	8,  // 4: Message.setChar:type_name -> SetPlayer
	9,  // 5: Message.start:type_name -> StartGame
	2,  // 6: Message.rAns:type_name -> RemoveAnswer
	3,  // 7: Message.cards:type_name -> Cards
	7,  // 8: Opponent.con:type_name -> Connect
	8,  // 9: Opponent.setChar:type_name -> SetPlayer
	0,  // 10: Command.type:type_name -> CommandType
	1,  // 11: ClueService.GameStream:input_type -> Message
	1,  // 12: ClueService.GameStream:output_type -> Message
	12, // [12:13] is the sub-list for method output_type
	11, // [11:12] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_proto_clue_proto_init() }
func file_proto_clue_proto_init() {
	if File_proto_clue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_clue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_clue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveAnswer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_clue_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cards); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_clue_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Opponent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_clue_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameEnd); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_clue_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_clue_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connect); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_clue_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPlayer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_clue_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartGame); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_proto_clue_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Message_Con)(nil),
		(*Message_Opp)(nil),
		(*Message_Com)(nil),
		(*Message_End)(nil),
		(*Message_SetChar)(nil),
		(*Message_Start)(nil),
		(*Message_RAns)(nil),
		(*Message_Cards)(nil),
	}
	file_proto_clue_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Opponent_Con)(nil),
		(*Opponent_SetChar)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_clue_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_clue_proto_goTypes,
		DependencyIndexes: file_proto_clue_proto_depIdxs,
		EnumInfos:         file_proto_clue_proto_enumTypes,
		MessageInfos:      file_proto_clue_proto_msgTypes,
	}.Build()
	File_proto_clue_proto = out.File
	file_proto_clue_proto_rawDesc = nil
	file_proto_clue_proto_goTypes = nil
	file_proto_clue_proto_depIdxs = nil
}
