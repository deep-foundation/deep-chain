package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
)

//func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
//	cdc.RegisterConcrete(&MsgCyberlink{}, "cyber/MsgCyberlink", nil)
//}

//func RegisterInterfaces(registry types.InterfaceRegistry) {
//	registry.RegisterImplementations((*sdk.Msg)(nil),
//		&MsgCyberlink{},
//	)
//
//	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
//}

var (
	amino = codec.NewLegacyAmino()

	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func init() {
	//RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
