// Code generated by protoc-gen-go.
// source: base.data.proto
// DO NOT EDIT!

package protodata

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

// *****************************玩家数据信息******************************************
type RoleData struct {
	RoleId            *int64  `protobuf:"varint,1,opt,name=roleId" json:"roleId,omitempty"`
	RoleName          *string `protobuf:"bytes,2,opt,name=roleName" json:"roleName,omitempty"`
	KillNum           *int32  `protobuf:"varint,3,opt,name=killNum" json:"killNum,omitempty"`
	Stamina           *int32  `protobuf:"varint,4,opt,name=stamina" json:"stamina,omitempty"`
	MaxStamina        *int32  `protobuf:"varint,5,opt,name=maxStamina" json:"maxStamina,omitempty"`
	Coin              *int32  `protobuf:"varint,6,opt,name=coin" json:"coin,omitempty"`
	Diamond           *int32  `protobuf:"varint,7,opt,name=diamond" json:"diamond,omitempty"`
	SuppleStaminaTime *int32  `protobuf:"varint,9,opt,name=suppleStaminaTime" json:"suppleStaminaTime,omitempty"`
	SuppleStaDiamond  *int32  `protobuf:"varint,10,opt,name=suppleStaDiamond" json:"suppleStaDiamond,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (m *RoleData) Reset()         { *m = RoleData{} }
func (m *RoleData) String() string { return proto.CompactTextString(m) }
func (*RoleData) ProtoMessage()    {}

func (m *RoleData) GetRoleId() int64 {
	if m != nil && m.RoleId != nil {
		return *m.RoleId
	}
	return 0
}

func (m *RoleData) GetRoleName() string {
	if m != nil && m.RoleName != nil {
		return *m.RoleName
	}
	return ""
}

func (m *RoleData) GetKillNum() int32 {
	if m != nil && m.KillNum != nil {
		return *m.KillNum
	}
	return 0
}

func (m *RoleData) GetStamina() int32 {
	if m != nil && m.Stamina != nil {
		return *m.Stamina
	}
	return 0
}

func (m *RoleData) GetMaxStamina() int32 {
	if m != nil && m.MaxStamina != nil {
		return *m.MaxStamina
	}
	return 0
}

func (m *RoleData) GetCoin() int32 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *RoleData) GetDiamond() int32 {
	if m != nil && m.Diamond != nil {
		return *m.Diamond
	}
	return 0
}

func (m *RoleData) GetSuppleStaminaTime() int32 {
	if m != nil && m.SuppleStaminaTime != nil {
		return *m.SuppleStaminaTime
	}
	return 0
}

func (m *RoleData) GetSuppleStaDiamond() int32 {
	if m != nil && m.SuppleStaDiamond != nil {
		return *m.SuppleStaDiamond
	}
	return 0
}

// *****************************英雄数据******************************************
type GeneralData struct {
	GeneralId   *int32  `protobuf:"varint,1,opt,name=generalId" json:"generalId,omitempty"`
	GeneralName *string `protobuf:"bytes,2,opt,name=generalName" json:"generalName,omitempty"`
	GeneralDesc *string `protobuf:"bytes,3,opt,name=generalDesc" json:"generalDesc,omitempty"`
	Level       *int32  `protobuf:"varint,4,opt,name=level" json:"level,omitempty"`
	Hp          *int32  `protobuf:"varint,5,opt,name=hp" json:"hp,omitempty"`
	Atk         *int32  `protobuf:"varint,6,opt,name=atk" json:"atk,omitempty"`
	Def         *int32  `protobuf:"varint,7,opt,name=def" json:"def,omitempty"`
	Speed       *int32  `protobuf:"varint,8,opt,name=speed" json:"speed,omitempty"`
	Dex         *int32  `protobuf:"varint,9,opt,name=dex" json:"dex,omitempty"`
	TriggerR    *int32  `protobuf:"varint,10,opt,name=triggerR" json:"triggerR,omitempty"`
	AtkR        *int32  `protobuf:"varint,11,opt,name=atkR" json:"atkR,omitempty"`
	// 	optional int32 atkType = 12;			//攻击类型1近程 2远程 3炸弹
	GeneralType      *int32 `protobuf:"varint,13,opt,name=generalType" json:"generalType,omitempty"`
	BuyDiamond       *int32 `protobuf:"varint,14,opt,name=buyDiamond" json:"buyDiamond,omitempty"`
	LevelUpCoin      *int32 `protobuf:"varint,15,opt,name=levelUpCoin" json:"levelUpCoin,omitempty"`
	IsUnlock         *bool  `protobuf:"varint,16,opt,name=isUnlock" json:"isUnlock,omitempty"`
	KillNum          *int32 `protobuf:"varint,17,opt,name=killNum" json:"killNum,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GeneralData) Reset()         { *m = GeneralData{} }
func (m *GeneralData) String() string { return proto.CompactTextString(m) }
func (*GeneralData) ProtoMessage()    {}

func (m *GeneralData) GetGeneralId() int32 {
	if m != nil && m.GeneralId != nil {
		return *m.GeneralId
	}
	return 0
}

func (m *GeneralData) GetGeneralName() string {
	if m != nil && m.GeneralName != nil {
		return *m.GeneralName
	}
	return ""
}

func (m *GeneralData) GetGeneralDesc() string {
	if m != nil && m.GeneralDesc != nil {
		return *m.GeneralDesc
	}
	return ""
}

func (m *GeneralData) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *GeneralData) GetHp() int32 {
	if m != nil && m.Hp != nil {
		return *m.Hp
	}
	return 0
}

func (m *GeneralData) GetAtk() int32 {
	if m != nil && m.Atk != nil {
		return *m.Atk
	}
	return 0
}

func (m *GeneralData) GetDef() int32 {
	if m != nil && m.Def != nil {
		return *m.Def
	}
	return 0
}

func (m *GeneralData) GetSpeed() int32 {
	if m != nil && m.Speed != nil {
		return *m.Speed
	}
	return 0
}

func (m *GeneralData) GetDex() int32 {
	if m != nil && m.Dex != nil {
		return *m.Dex
	}
	return 0
}

func (m *GeneralData) GetTriggerR() int32 {
	if m != nil && m.TriggerR != nil {
		return *m.TriggerR
	}
	return 0
}

func (m *GeneralData) GetAtkR() int32 {
	if m != nil && m.AtkR != nil {
		return *m.AtkR
	}
	return 0
}

func (m *GeneralData) GetGeneralType() int32 {
	if m != nil && m.GeneralType != nil {
		return *m.GeneralType
	}
	return 0
}

func (m *GeneralData) GetBuyDiamond() int32 {
	if m != nil && m.BuyDiamond != nil {
		return *m.BuyDiamond
	}
	return 0
}

func (m *GeneralData) GetLevelUpCoin() int32 {
	if m != nil && m.LevelUpCoin != nil {
		return *m.LevelUpCoin
	}
	return 0
}

func (m *GeneralData) GetIsUnlock() bool {
	if m != nil && m.IsUnlock != nil {
		return *m.IsUnlock
	}
	return false
}

func (m *GeneralData) GetKillNum() int32 {
	if m != nil && m.KillNum != nil {
		return *m.KillNum
	}
	return 0
}

// *****************************道具信息******************************************
type ItemData struct {
	ItemId           *int32  `protobuf:"varint,1,opt,name=itemId" json:"itemId,omitempty"`
	ItemName         *string `protobuf:"bytes,2,opt,name=itemName" json:"itemName,omitempty"`
	ItemDesc         *string `protobuf:"bytes,3,opt,name=itemDesc" json:"itemDesc,omitempty"`
	Level            *int32  `protobuf:"varint,4,opt,name=level" json:"level,omitempty"`
	LevelUpCoin      *int32  `protobuf:"varint,5,opt,name=levelUpCoin" json:"levelUpCoin,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ItemData) Reset()         { *m = ItemData{} }
func (m *ItemData) String() string { return proto.CompactTextString(m) }
func (*ItemData) ProtoMessage()    {}

func (m *ItemData) GetItemId() int32 {
	if m != nil && m.ItemId != nil {
		return *m.ItemId
	}
	return 0
}

func (m *ItemData) GetItemName() string {
	if m != nil && m.ItemName != nil {
		return *m.ItemName
	}
	return ""
}

func (m *ItemData) GetItemDesc() string {
	if m != nil && m.ItemDesc != nil {
		return *m.ItemDesc
	}
	return ""
}

func (m *ItemData) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *ItemData) GetLevelUpCoin() int32 {
	if m != nil && m.LevelUpCoin != nil {
		return *m.LevelUpCoin
	}
	return 0
}

// *****************************副本信息******************************************
type SectionData struct {
	SectionId        *int32  `protobuf:"varint,1,opt,name=sectionId" json:"sectionId,omitempty"`
	SectionName      *string `protobuf:"bytes,2,opt,name=sectionName" json:"sectionName,omitempty"`
	SectionDesc      *string `protobuf:"bytes,3,opt,name=sectionDesc" json:"sectionDesc,omitempty"`
	IsUnlock         *bool   `protobuf:"varint,4,opt,name=isUnlock" json:"isUnlock,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SectionData) Reset()         { *m = SectionData{} }
func (m *SectionData) String() string { return proto.CompactTextString(m) }
func (*SectionData) ProtoMessage()    {}

func (m *SectionData) GetSectionId() int32 {
	if m != nil && m.SectionId != nil {
		return *m.SectionId
	}
	return 0
}

func (m *SectionData) GetSectionName() string {
	if m != nil && m.SectionName != nil {
		return *m.SectionName
	}
	return ""
}

func (m *SectionData) GetSectionDesc() string {
	if m != nil && m.SectionDesc != nil {
		return *m.SectionDesc
	}
	return ""
}

func (m *SectionData) GetIsUnlock() bool {
	if m != nil && m.IsUnlock != nil {
		return *m.IsUnlock
	}
	return false
}

// *****************************章节信息******************************************
type ChapterData struct {
	ChapterId        *int32         `protobuf:"varint,1,opt,name=chapterId" json:"chapterId,omitempty"`
	ChapterName      *string        `protobuf:"bytes,2,opt,name=chapterName" json:"chapterName,omitempty"`
	ChapterDesc      *string        `protobuf:"bytes,3,opt,name=chapterDesc" json:"chapterDesc,omitempty"`
	Sections         []*SectionData `protobuf:"bytes,4,rep" json:"Sections,omitempty"`
	IsUnlock         *bool          `protobuf:"varint,5,opt,name=isUnlock" json:"isUnlock,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *ChapterData) Reset()         { *m = ChapterData{} }
func (m *ChapterData) String() string { return proto.CompactTextString(m) }
func (*ChapterData) ProtoMessage()    {}

func (m *ChapterData) GetChapterId() int32 {
	if m != nil && m.ChapterId != nil {
		return *m.ChapterId
	}
	return 0
}

func (m *ChapterData) GetChapterName() string {
	if m != nil && m.ChapterName != nil {
		return *m.ChapterName
	}
	return ""
}

func (m *ChapterData) GetChapterDesc() string {
	if m != nil && m.ChapterDesc != nil {
		return *m.ChapterDesc
	}
	return ""
}

func (m *ChapterData) GetSections() []*SectionData {
	if m != nil {
		return m.Sections
	}
	return nil
}

func (m *ChapterData) GetIsUnlock() bool {
	if m != nil && m.IsUnlock != nil {
		return *m.IsUnlock
	}
	return false
}

// *****************************金币充值商品信息**********************************
type CoinProductData struct {
	ProductIndex     *int32  `protobuf:"varint,1,opt,name=productIndex" json:"productIndex,omitempty"`
	ProductName      *string `protobuf:"bytes,2,opt,name=productName" json:"productName,omitempty"`
	ProductDesc      *string `protobuf:"bytes,3,opt,name=productDesc" json:"productDesc,omitempty"`
	ProductCoin      *int32  `protobuf:"varint,4,opt,name=productCoin" json:"productCoin,omitempty"`
	PriceDiamond     *int32  `protobuf:"varint,5,opt,name=priceDiamond" json:"priceDiamond,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CoinProductData) Reset()         { *m = CoinProductData{} }
func (m *CoinProductData) String() string { return proto.CompactTextString(m) }
func (*CoinProductData) ProtoMessage()    {}

func (m *CoinProductData) GetProductIndex() int32 {
	if m != nil && m.ProductIndex != nil {
		return *m.ProductIndex
	}
	return 0
}

func (m *CoinProductData) GetProductName() string {
	if m != nil && m.ProductName != nil {
		return *m.ProductName
	}
	return ""
}

func (m *CoinProductData) GetProductDesc() string {
	if m != nil && m.ProductDesc != nil {
		return *m.ProductDesc
	}
	return ""
}

func (m *CoinProductData) GetProductCoin() int32 {
	if m != nil && m.ProductCoin != nil {
		return *m.ProductCoin
	}
	return 0
}

func (m *CoinProductData) GetPriceDiamond() int32 {
	if m != nil && m.PriceDiamond != nil {
		return *m.PriceDiamond
	}
	return 0
}

// *****************************钻石充值商品信息**********************************
type DiamondProductData struct {
	ProductIndex     *int32  `protobuf:"varint,1,opt,name=productIndex" json:"productIndex,omitempty"`
	ProductName      *string `protobuf:"bytes,2,opt,name=productName" json:"productName,omitempty"`
	ProductDesc      *string `protobuf:"bytes,3,opt,name=productDesc" json:"productDesc,omitempty"`
	ProductDiamond   *int32  `protobuf:"varint,4,opt,name=productDiamond" json:"productDiamond,omitempty"`
	Price            *int32  `protobuf:"varint,5,opt,name=price" json:"price,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DiamondProductData) Reset()         { *m = DiamondProductData{} }
func (m *DiamondProductData) String() string { return proto.CompactTextString(m) }
func (*DiamondProductData) ProtoMessage()    {}

func (m *DiamondProductData) GetProductIndex() int32 {
	if m != nil && m.ProductIndex != nil {
		return *m.ProductIndex
	}
	return 0
}

func (m *DiamondProductData) GetProductName() string {
	if m != nil && m.ProductName != nil {
		return *m.ProductName
	}
	return ""
}

func (m *DiamondProductData) GetProductDesc() string {
	if m != nil && m.ProductDesc != nil {
		return *m.ProductDesc
	}
	return ""
}

func (m *DiamondProductData) GetProductDiamond() int32 {
	if m != nil && m.ProductDiamond != nil {
		return *m.ProductDiamond
	}
	return 0
}

func (m *DiamondProductData) GetPrice() int32 {
	if m != nil && m.Price != nil {
		return *m.Price
	}
	return 0
}

// *****************************奖励信息**********************************
type RewardData struct {
	RewardCoin       *int32       `protobuf:"varint,1,opt,name=rewardCoin" json:"rewardCoin,omitempty"`
	RewardDiamond    *int32       `protobuf:"varint,2,opt,name=rewardDiamond" json:"rewardDiamond,omitempty"`
	General          *GeneralData `protobuf:"bytes,3,opt,name=general" json:"general,omitempty"`
	Stamina          *int32       `protobuf:"varint,4,opt,name=stamina" json:"stamina,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *RewardData) Reset()         { *m = RewardData{} }
func (m *RewardData) String() string { return proto.CompactTextString(m) }
func (*RewardData) ProtoMessage()    {}

func (m *RewardData) GetRewardCoin() int32 {
	if m != nil && m.RewardCoin != nil {
		return *m.RewardCoin
	}
	return 0
}

func (m *RewardData) GetRewardDiamond() int32 {
	if m != nil && m.RewardDiamond != nil {
		return *m.RewardDiamond
	}
	return 0
}

func (m *RewardData) GetGeneral() *GeneralData {
	if m != nil {
		return m.General
	}
	return nil
}

func (m *RewardData) GetStamina() int32 {
	if m != nil && m.Stamina != nil {
		return *m.Stamina
	}
	return 0
}

// *****************************签到信息**********************************
type SignRewardData struct {
	SignIndex        *int32      `protobuf:"varint,1,opt,name=signIndex" json:"signIndex,omitempty"`
	Reward           *RewardData `protobuf:"bytes,2,opt,name=reward" json:"reward,omitempty"`
	IsReceive        *bool       `protobuf:"varint,3,opt,name=isReceive" json:"isReceive,omitempty"`
	SignDay          *int32      `protobuf:"varint,4,opt,name=signDay" json:"signDay,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *SignRewardData) Reset()         { *m = SignRewardData{} }
func (m *SignRewardData) String() string { return proto.CompactTextString(m) }
func (*SignRewardData) ProtoMessage()    {}

func (m *SignRewardData) GetSignIndex() int32 {
	if m != nil && m.SignIndex != nil {
		return *m.SignIndex
	}
	return 0
}

func (m *SignRewardData) GetReward() *RewardData {
	if m != nil {
		return m.Reward
	}
	return nil
}

func (m *SignRewardData) GetIsReceive() bool {
	if m != nil && m.IsReceive != nil {
		return *m.IsReceive
	}
	return false
}

func (m *SignRewardData) GetSignDay() int32 {
	if m != nil && m.SignDay != nil {
		return *m.SignDay
	}
	return 0
}

// *****************************邮件信息**********************************
type MailData struct {
	MailId           *int32      `protobuf:"varint,1,opt,name=mailId" json:"mailId,omitempty"`
	MailTitle        *string     `protobuf:"bytes,2,opt,name=mailTitle" json:"mailTitle,omitempty"`
	MailContent      *string     `protobuf:"bytes,3,opt,name=mailContent" json:"mailContent,omitempty"`
	Reward           *RewardData `protobuf:"bytes,4,opt,name=reward" json:"reward,omitempty"`
	IsReceive        *bool       `protobuf:"varint,6,opt,name=isReceive" json:"isReceive,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *MailData) Reset()         { *m = MailData{} }
func (m *MailData) String() string { return proto.CompactTextString(m) }
func (*MailData) ProtoMessage()    {}

func (m *MailData) GetMailId() int32 {
	if m != nil && m.MailId != nil {
		return *m.MailId
	}
	return 0
}

func (m *MailData) GetMailTitle() string {
	if m != nil && m.MailTitle != nil {
		return *m.MailTitle
	}
	return ""
}

func (m *MailData) GetMailContent() string {
	if m != nil && m.MailContent != nil {
		return *m.MailContent
	}
	return ""
}

func (m *MailData) GetReward() *RewardData {
	if m != nil {
		return m.Reward
	}
	return nil
}

func (m *MailData) GetIsReceive() bool {
	if m != nil && m.IsReceive != nil {
		return *m.IsReceive
	}
	return false
}

// *****************************通用规则信息******************************************
type FuncData struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *FuncData) Reset()         { *m = FuncData{} }
func (m *FuncData) String() string { return proto.CompactTextString(m) }
func (*FuncData) ProtoMessage()    {}

func init() {
}
