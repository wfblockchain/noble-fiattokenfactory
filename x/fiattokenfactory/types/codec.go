package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func init() {
	RegisterLegacyAminoCodec(amino)
	amino.Seal()
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgUpdateMasterMinter{}, "fiattokenfactory-nobel/UpdateMasterMinter", nil)
	cdc.RegisterConcrete(&MsgUpdatePauser{}, "fiattokenfactory-nobel/UpdatePauser", nil)
	cdc.RegisterConcrete(&MsgUpdateBlacklister{}, "fiattokenfactory-nobel/UpdateBlacklister", nil)
	cdc.RegisterConcrete(&MsgUpdateOwner{}, "fiattokenfactory-nobel/UpdateOwner", nil)
	cdc.RegisterConcrete(&MsgConfigureMinter{}, "fiattokenfactory-nobel/ConfigureMinter", nil)
	cdc.RegisterConcrete(&MsgRemoveMinter{}, "fiattokenfactory-nobel/RemoveMinter", nil)
	cdc.RegisterConcrete(&MsgMint{}, "fiattokenfactory-nobel/Mint", nil)
	cdc.RegisterConcrete(&MsgBurn{}, "fiattokenfactory-nobel/Burn", nil)
	cdc.RegisterConcrete(&MsgBlacklist{}, "fiattokenfactory-nobel/Blacklist", nil)
	cdc.RegisterConcrete(&MsgUnblacklist{}, "fiattokenfactory-nobel/Unblacklist", nil)
	cdc.RegisterConcrete(&MsgPause{}, "fiattokenfactory-nobel/Pause", nil)
	cdc.RegisterConcrete(&MsgUnpause{}, "fiattokenfactory-nobel/Unpause", nil)
	cdc.RegisterConcrete(&MsgConfigureMinterController{}, "fiattokenfactory-nobel/ConfigureMinterController", nil)
	cdc.RegisterConcrete(&MsgRemoveMinterController{}, "fiattokenfactory-nobel/RemoveMinterController", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateMasterMinter{},
		&MsgUpdatePauser{},
		&MsgUpdateBlacklister{},
		&MsgUpdateOwner{},
		&MsgConfigureMinter{},
		&MsgRemoveMinter{},
		&MsgMint{},
		&MsgBurn{},
		&MsgBlacklist{},
		&MsgUnblacklist{},
		&MsgPause{},
		&MsgUnpause{},
		&MsgConfigureMinterController{},
		&MsgRemoveMinterController{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
