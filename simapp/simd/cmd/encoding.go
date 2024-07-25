package cmd

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/testutil"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
)

// EncodingConfig specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Marshaler         codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

// makeEncodingConfig creates an EncodingConfig for an amino based test configuration.
func makeEncodingConfig() EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := testutil.CodecOptions{}.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(marshaler, tx.DefaultSignModes)

	encCfg := EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             amino,
	}

	mb := module.NewBasicManager()

	std.RegisterLegacyAminoCodec(encCfg.Amino)
	std.RegisterInterfaces(encCfg.InterfaceRegistry)
	mb.RegisterLegacyAminoCodec(encCfg.Amino)
	mb.RegisterInterfaces(encCfg.InterfaceRegistry)

	return encCfg
	// return EncodingConfig{
	// 	InterfaceRegistry: interfaceRegistry,
	// 	Marshaler:         marshaler,
	// 	TxConfig:          txCfg,
	// 	Amino:             amino,
	// }
}

func MakeEncodingConfig(moduleBasics module.BasicManager) EncodingConfig {
	aminoCodec := codec.NewLegacyAmino()
	interfaceRegistry := testutil.CodecOptions{}.NewInterfaceRegistry()
	codec := codec.NewProtoCodec(interfaceRegistry)

	encCfg := EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         codec,
		TxConfig:          tx.NewTxConfig(codec, tx.DefaultSignModes),
		Amino:             aminoCodec,
	}

	//mb := module.NewBasicManager(modules...)

	std.RegisterLegacyAminoCodec(encCfg.Amino)
	std.RegisterInterfaces(encCfg.InterfaceRegistry)
	moduleBasics.RegisterLegacyAminoCodec(encCfg.Amino)
	moduleBasics.RegisterInterfaces(encCfg.InterfaceRegistry)

	return encCfg
}

// MakeEncodingConfig creates an EncodingConfig for testing
// func MakeEncodingConfig(moduleBasics module.BasicManager) EncodingConfig {
// 	encodingConfig := makeEncodingConfig()
// 	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
// 	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
// 	moduleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
// 	moduleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
// 	return encodingConfig
// }
