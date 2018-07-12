// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/zombie_battleground/types/zb/zb.proto

/*
Package zb is a generated protocol buffer package.

It is generated from these files:
	github.com/loomnetwork/zombie_battleground/types/zb/zb.proto

It has these top-level messages:
	ZBAccount
	UpsertAccountRequest
	GetAccountRequest
	GetDecksRequest
	DecksResponse
	ZBDeck
	UserDecks
	CardInDeck
	GetDeckRequest
	GetDecksResponse
	InitRequest
*/
package zb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/loomnetwork/go-loom/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ZBAccount struct {
	UserId              string         `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PhoneNumberVerified bool           `protobuf:"varint,2,opt,name=phone_number_verified,json=phoneNumberVerified,proto3" json:"phone_number_verified,omitempty"`
	RewardRedeemed      bool           `protobuf:"varint,3,opt,name=reward_redeemed,json=rewardRedeemed,proto3" json:"reward_redeemed,omitempty"`
	IsKickstarter       bool           `protobuf:"varint,4,opt,name=is_kickstarter,json=isKickstarter,proto3" json:"is_kickstarter,omitempty"`
	Image               string         `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	EmailNotification   bool           `protobuf:"varint,6,opt,name=email_notification,json=emailNotification,proto3" json:"email_notification,omitempty"`
	EloScore            int64          `protobuf:"varint,7,opt,name=elo_score,json=eloScore,proto3" json:"elo_score,omitempty"`
	CurrentTier         int32          `protobuf:"varint,8,opt,name=current_tier,json=currentTier,proto3" json:"current_tier,omitempty"`
	GameMembershipTier  int32          `protobuf:"varint,9,opt,name=game_membership_tier,json=gameMembershipTier,proto3" json:"game_membership_tier,omitempty"`
	Owner               *types.Address `protobuf:"bytes,10,opt,name=owner" json:"owner,omitempty"`
}

func (m *ZBAccount) Reset()                    { *m = ZBAccount{} }
func (m *ZBAccount) String() string            { return proto.CompactTextString(m) }
func (*ZBAccount) ProtoMessage()               {}
func (*ZBAccount) Descriptor() ([]byte, []int) { return fileDescriptorZb, []int{0} }

func (m *ZBAccount) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ZBAccount) GetPhoneNumberVerified() bool {
	if m != nil {
		return m.PhoneNumberVerified
	}
	return false
}

func (m *ZBAccount) GetRewardRedeemed() bool {
	if m != nil {
		return m.RewardRedeemed
	}
	return false
}

func (m *ZBAccount) GetIsKickstarter() bool {
	if m != nil {
		return m.IsKickstarter
	}
	return false
}

func (m *ZBAccount) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *ZBAccount) GetEmailNotification() bool {
	if m != nil {
		return m.EmailNotification
	}
	return false
}

func (m *ZBAccount) GetEloScore() int64 {
	if m != nil {
		return m.EloScore
	}
	return 0
}

func (m *ZBAccount) GetCurrentTier() int32 {
	if m != nil {
		return m.CurrentTier
	}
	return 0
}

func (m *ZBAccount) GetGameMembershipTier() int32 {
	if m != nil {
		return m.GameMembershipTier
	}
	return 0
}

func (m *ZBAccount) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

type UpsertAccountRequest struct {
	UserId              string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PhoneNumberVerified bool   `protobuf:"varint,2,opt,name=phone_number_verified,json=phoneNumberVerified,proto3" json:"phone_number_verified,omitempty"`
	RewardRedeemed      bool   `protobuf:"varint,3,opt,name=reward_redeemed,json=rewardRedeemed,proto3" json:"reward_redeemed,omitempty"`
	IsKickstarter       bool   `protobuf:"varint,4,opt,name=is_kickstarter,json=isKickstarter,proto3" json:"is_kickstarter,omitempty"`
	Image               string `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	EmailNotification   bool   `protobuf:"varint,6,opt,name=email_notification,json=emailNotification,proto3" json:"email_notification,omitempty"`
	EloScore            int64  `protobuf:"varint,7,opt,name=elo_score,json=eloScore,proto3" json:"elo_score,omitempty"`
	CurrentTier         int32  `protobuf:"varint,8,opt,name=current_tier,json=currentTier,proto3" json:"current_tier,omitempty"`
	GameMembershipTier  int32  `protobuf:"varint,9,opt,name=game_membership_tier,json=gameMembershipTier,proto3" json:"game_membership_tier,omitempty"`
}

func (m *UpsertAccountRequest) Reset()                    { *m = UpsertAccountRequest{} }
func (m *UpsertAccountRequest) String() string            { return proto.CompactTextString(m) }
func (*UpsertAccountRequest) ProtoMessage()               {}
func (*UpsertAccountRequest) Descriptor() ([]byte, []int) { return fileDescriptorZb, []int{1} }

func (m *UpsertAccountRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UpsertAccountRequest) GetPhoneNumberVerified() bool {
	if m != nil {
		return m.PhoneNumberVerified
	}
	return false
}

func (m *UpsertAccountRequest) GetRewardRedeemed() bool {
	if m != nil {
		return m.RewardRedeemed
	}
	return false
}

func (m *UpsertAccountRequest) GetIsKickstarter() bool {
	if m != nil {
		return m.IsKickstarter
	}
	return false
}

func (m *UpsertAccountRequest) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *UpsertAccountRequest) GetEmailNotification() bool {
	if m != nil {
		return m.EmailNotification
	}
	return false
}

func (m *UpsertAccountRequest) GetEloScore() int64 {
	if m != nil {
		return m.EloScore
	}
	return 0
}

func (m *UpsertAccountRequest) GetCurrentTier() int32 {
	if m != nil {
		return m.CurrentTier
	}
	return 0
}

func (m *UpsertAccountRequest) GetGameMembershipTier() int32 {
	if m != nil {
		return m.GameMembershipTier
	}
	return 0
}

type GetAccountRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (m *GetAccountRequest) Reset()                    { *m = GetAccountRequest{} }
func (m *GetAccountRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAccountRequest) ProtoMessage()               {}
func (*GetAccountRequest) Descriptor() ([]byte, []int) { return fileDescriptorZb, []int{2} }

func (m *GetAccountRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetDecksRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (m *GetDecksRequest) Reset()                    { *m = GetDecksRequest{} }
func (m *GetDecksRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDecksRequest) ProtoMessage()               {}
func (*GetDecksRequest) Descriptor() ([]byte, []int) { return fileDescriptorZb, []int{3} }

func (m *GetDecksRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type DecksResponse struct {
	UserId string    `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Decks  []*ZBDeck `protobuf:"bytes,2,rep,name=decks" json:"decks,omitempty"`
}

func (m *DecksResponse) Reset()                    { *m = DecksResponse{} }
func (m *DecksResponse) String() string            { return proto.CompactTextString(m) }
func (*DecksResponse) ProtoMessage()               {}
func (*DecksResponse) Descriptor() ([]byte, []int) { return fileDescriptorZb, []int{4} }

func (m *DecksResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *DecksResponse) GetDecks() []*ZBDeck {
	if m != nil {
		return m.Decks
	}
	return nil
}

type ZBDeck struct {
	Name   string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	HeroId int64         `protobuf:"varint,2,opt,name=hero_id,json=heroId,proto3" json:"hero_id,omitempty"`
	Cards  []*CardInDeck `protobuf:"bytes,3,rep,name=cards" json:"cards,omitempty"`
}

func (m *ZBDeck) Reset()                    { *m = ZBDeck{} }
func (m *ZBDeck) String() string            { return proto.CompactTextString(m) }
func (*ZBDeck) ProtoMessage()               {}
func (*ZBDeck) Descriptor() ([]byte, []int) { return fileDescriptorZb, []int{5} }

func (m *ZBDeck) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ZBDeck) GetHeroId() int64 {
	if m != nil {
		return m.HeroId
	}
	return 0
}

func (m *ZBDeck) GetCards() []*CardInDeck {
	if m != nil {
		return m.Cards
	}
	return nil
}

type UserDecks struct {
	Decks []*ZBDeck `protobuf:"bytes,1,rep,name=decks" json:"decks,omitempty"`
}

func (m *UserDecks) Reset()                    { *m = UserDecks{} }
func (m *UserDecks) String() string            { return proto.CompactTextString(m) }
func (*UserDecks) ProtoMessage()               {}
func (*UserDecks) Descriptor() ([]byte, []int) { return fileDescriptorZb, []int{6} }

func (m *UserDecks) GetDecks() []*ZBDeck {
	if m != nil {
		return m.Decks
	}
	return nil
}

type CardInDeck struct {
	CardId int64 `protobuf:"varint,1,opt,name=card_id,json=cardId,proto3" json:"card_id,omitempty"`
	Amount int64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *CardInDeck) Reset()                    { *m = CardInDeck{} }
func (m *CardInDeck) String() string            { return proto.CompactTextString(m) }
func (*CardInDeck) ProtoMessage()               {}
func (*CardInDeck) Descriptor() ([]byte, []int) { return fileDescriptorZb, []int{7} }

func (m *CardInDeck) GetCardId() int64 {
	if m != nil {
		return m.CardId
	}
	return 0
}

func (m *CardInDeck) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type GetDeckRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DeckId string `protobuf:"bytes,2,opt,name=deck_id,json=deckId,proto3" json:"deck_id,omitempty"`
}

func (m *GetDeckRequest) Reset()                    { *m = GetDeckRequest{} }
func (m *GetDeckRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDeckRequest) ProtoMessage()               {}
func (*GetDeckRequest) Descriptor() ([]byte, []int) { return fileDescriptorZb, []int{8} }

func (m *GetDeckRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *GetDeckRequest) GetDeckId() string {
	if m != nil {
		return m.DeckId
	}
	return ""
}

type GetDecksResponse struct {
	Decks []*ZBDeck `protobuf:"bytes,1,rep,name=decks" json:"decks,omitempty"`
}

func (m *GetDecksResponse) Reset()                    { *m = GetDecksResponse{} }
func (m *GetDecksResponse) String() string            { return proto.CompactTextString(m) }
func (*GetDecksResponse) ProtoMessage()               {}
func (*GetDecksResponse) Descriptor() ([]byte, []int) { return fileDescriptorZb, []int{9} }

func (m *GetDecksResponse) GetDecks() []*ZBDeck {
	if m != nil {
		return m.Decks
	}
	return nil
}

type InitRequest struct {
	DefaultDecks []*ZBDeck `protobuf:"bytes,1,rep,name=default_decks,json=defaultDecks" json:"default_decks,omitempty"`
}

func (m *InitRequest) Reset()                    { *m = InitRequest{} }
func (m *InitRequest) String() string            { return proto.CompactTextString(m) }
func (*InitRequest) ProtoMessage()               {}
func (*InitRequest) Descriptor() ([]byte, []int) { return fileDescriptorZb, []int{10} }

func (m *InitRequest) GetDefaultDecks() []*ZBDeck {
	if m != nil {
		return m.DefaultDecks
	}
	return nil
}

func init() {
	proto.RegisterType((*ZBAccount)(nil), "ZBAccount")
	proto.RegisterType((*UpsertAccountRequest)(nil), "UpsertAccountRequest")
	proto.RegisterType((*GetAccountRequest)(nil), "GetAccountRequest")
	proto.RegisterType((*GetDecksRequest)(nil), "GetDecksRequest")
	proto.RegisterType((*DecksResponse)(nil), "DecksResponse")
	proto.RegisterType((*ZBDeck)(nil), "ZBDeck")
	proto.RegisterType((*UserDecks)(nil), "UserDecks")
	proto.RegisterType((*CardInDeck)(nil), "CardInDeck")
	proto.RegisterType((*GetDeckRequest)(nil), "GetDeckRequest")
	proto.RegisterType((*GetDecksResponse)(nil), "GetDecksResponse")
	proto.RegisterType((*InitRequest)(nil), "InitRequest")
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/zombie_battleground/types/zb/zb.proto", fileDescriptorZb)
}

var fileDescriptorZb = []byte{
	// 579 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x54, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0x96, 0x93, 0xc6, 0x49, 0x26, 0xfd, 0x78, 0xbb, 0x6f, 0xa1, 0x16, 0x08, 0x94, 0x5a, 0x42,
	0x58, 0x55, 0x9b, 0x94, 0x72, 0x04, 0x0e, 0x2d, 0x48, 0x95, 0x85, 0xe8, 0xc1, 0x50, 0x84, 0x7a,
	0xb1, 0xd6, 0xde, 0x69, 0xb2, 0x8a, 0xbd, 0x6b, 0x76, 0xd7, 0x54, 0xf4, 0x27, 0xf3, 0x27, 0x40,
	0xbb, 0x76, 0x1b, 0x0e, 0xad, 0xca, 0x85, 0x1b, 0x17, 0xcb, 0xcf, 0x3c, 0x1f, 0x3b, 0x3b, 0x63,
	0x19, 0x5e, 0xcf, 0xb8, 0x99, 0xd7, 0xd9, 0x24, 0x97, 0xe5, 0xb4, 0x90, 0xb2, 0x14, 0x68, 0x2e,
	0xa5, 0x5a, 0x4c, 0xaf, 0x64, 0x99, 0x71, 0x4c, 0x33, 0x6a, 0x4c, 0x81, 0x33, 0x25, 0x6b, 0xc1,
	0xa6, 0xe6, 0x7b, 0x85, 0x7a, 0x7a, 0x95, 0x4d, 0xaf, 0xb2, 0x49, 0xa5, 0xa4, 0x91, 0x8f, 0x0e,
	0xee, 0x70, 0xcf, 0xe4, 0xbe, 0x85, 0xad, 0xc3, 0x3d, 0x1b, 0x47, 0xf8, 0xb3, 0x03, 0xc3, 0xf3,
	0xe3, 0xa3, 0x3c, 0x97, 0xb5, 0x30, 0x64, 0x1b, 0xfa, 0xb5, 0x46, 0x95, 0x72, 0x16, 0x78, 0x63,
	0x2f, 0x1a, 0x26, 0xbe, 0x85, 0x31, 0x23, 0x87, 0xf0, 0xa0, 0x9a, 0x4b, 0x81, 0xa9, 0xa8, 0xcb,
	0x0c, 0x55, 0xfa, 0x0d, 0x15, 0xbf, 0xe0, 0xc8, 0x82, 0xce, 0xd8, 0x8b, 0x06, 0xc9, 0xff, 0x8e,
	0x3c, 0x75, 0xdc, 0xe7, 0x96, 0x22, 0xcf, 0x61, 0x43, 0xe1, 0x25, 0x55, 0x2c, 0x55, 0xc8, 0x10,
	0x4b, 0x64, 0x41, 0xd7, 0xa9, 0xd7, 0x9b, 0x72, 0xd2, 0x56, 0xc9, 0x33, 0x58, 0xe7, 0x3a, 0x5d,
	0xf0, 0x7c, 0xa1, 0x0d, 0x55, 0x06, 0x55, 0xb0, 0xe2, 0x74, 0x6b, 0x5c, 0xbf, 0x5f, 0x16, 0xc9,
	0x16, 0xf4, 0x78, 0x49, 0x67, 0x18, 0xf4, 0x5c, 0x6b, 0x0d, 0x20, 0xfb, 0x40, 0xb0, 0xa4, 0xbc,
	0x48, 0x85, 0x34, 0xfc, 0x82, 0xe7, 0xd4, 0x70, 0x29, 0x02, 0xdf, 0x05, 0x6c, 0x3a, 0xe6, 0xf4,
	0x37, 0x82, 0x3c, 0x86, 0x21, 0x16, 0x32, 0xd5, 0xb9, 0x54, 0x18, 0xf4, 0xc7, 0x5e, 0xd4, 0x4d,
	0x06, 0x58, 0xc8, 0x8f, 0x16, 0x93, 0x1d, 0x58, 0xcd, 0x6b, 0xa5, 0x50, 0x98, 0xd4, 0x70, 0x54,
	0xc1, 0x60, 0xec, 0x45, 0xbd, 0x64, 0xd4, 0xd6, 0x3e, 0x71, 0x54, 0xe4, 0x00, 0xb6, 0x66, 0xb4,
	0xc4, 0xb4, 0x44, 0x7b, 0x57, 0x3d, 0xe7, 0x55, 0x23, 0x1d, 0x3a, 0x29, 0xb1, 0xdc, 0x87, 0x1b,
	0xca, 0x39, 0x9e, 0x42, 0x4f, 0x5e, 0x0a, 0x54, 0x01, 0x8c, 0xbd, 0x68, 0x74, 0x38, 0x98, 0x1c,
	0x31, 0xa6, 0x50, 0xeb, 0xa4, 0x29, 0x87, 0x3f, 0x3a, 0xb0, 0x75, 0x56, 0x69, 0x54, 0xa6, 0xdd,
	0x42, 0x82, 0x5f, 0x6b, 0xd4, 0xff, 0x96, 0xf1, 0x17, 0x96, 0x11, 0xee, 0xc1, 0xe6, 0x09, 0xfe,
	0xe9, 0xa0, 0xc3, 0x5d, 0xd8, 0x38, 0x41, 0xf3, 0x0e, 0xf3, 0x85, 0xbe, 0x57, 0x7b, 0x02, 0x6b,
	0xad, 0x50, 0x57, 0x52, 0x68, 0xbc, 0x7b, 0x7d, 0x4f, 0xa0, 0xc7, 0xac, 0x32, 0xe8, 0x8c, 0xbb,
	0xd1, 0xe8, 0xb0, 0x3f, 0x39, 0x3f, 0xb6, 0xce, 0xa4, 0xa9, 0x86, 0x5f, 0xc0, 0x6f, 0x0a, 0x84,
	0xc0, 0x8a, 0xa0, 0x25, 0xb6, 0x76, 0xf7, 0x6e, 0x53, 0xe7, 0xa8, 0xa4, 0x4d, 0xed, 0xb8, 0x81,
	0xf9, 0x16, 0xc6, 0x8c, 0xec, 0x40, 0x2f, 0xa7, 0x8a, 0xe9, 0xa0, 0xeb, 0x52, 0x47, 0x93, 0xb7,
	0x54, 0xb1, 0x58, 0x34, 0xc9, 0x8e, 0x09, 0x77, 0x61, 0x78, 0xa6, 0x51, 0xb9, 0x36, 0x97, 0x5d,
	0x78, 0xb7, 0x76, 0xf1, 0x06, 0x60, 0x19, 0x60, 0x4f, 0xb5, 0x11, 0xd7, 0x77, 0xe9, 0x26, 0xbe,
	0x85, 0x31, 0x23, 0x0f, 0xc1, 0xa7, 0xa5, 0x9d, 0xe5, 0x75, 0x37, 0x0d, 0x0a, 0x8f, 0x61, 0xbd,
	0x9d, 0xdc, 0xbd, 0x5f, 0xf3, 0x36, 0xf4, 0xed, 0x91, 0xd7, 0x37, 0x1a, 0x26, 0xbe, 0x85, 0x31,
	0x0b, 0x5f, 0xc0, 0x7f, 0xcb, 0xe9, 0xb7, 0x43, 0xbd, 0xa7, 0xeb, 0x57, 0x30, 0x8a, 0x05, 0xbf,
	0x59, 0xec, 0x1e, 0xac, 0x31, 0xbc, 0xa0, 0x75, 0x61, 0xd2, 0x5b, 0x5d, 0xab, 0x2d, 0xeb, 0xce,
	0xc8, 0x7c, 0xf7, 0x47, 0x7c, 0xf9, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x09, 0x32, 0xa4, 0x05, 0x83,
	0x05, 0x00, 0x00,
}
