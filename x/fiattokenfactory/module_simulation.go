package fiattokenfactory

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/wfblockchain/noble-fiattokenfactory/testutil/sample"
	tokenfactorysimulation "github.com/wfblockchain/noble-fiattokenfactory/x/fiattokenfactory/simulation"
	"github.com/wfblockchain/noble-fiattokenfactory/x/fiattokenfactory/types"

	// simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	// paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	// simcli "github.com/cosmos/cosmos-sdk/x/simulation/client/cli"
	"github.com/cosmos/cosmos-sdk/types/module"
	sdk "github.com/cosmos/cosmos-sdk/types/simulation"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = tokenfactorysimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgUpdateMasterMinter = "op_weight_msg_update_master_minter"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateMasterMinter int = 100

	opWeightMsgUpdatePauser = "op_weight_msg_update_pauser"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatePauser int = 100

	opWeightMsgUpdateBlacklister = "op_weight_msg_update_blacklister"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateBlacklister int = 100

	opWeightMsgUpdateOwner = "op_weight_msg_update_owner"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateOwner int = 100

	opWeightMsgConfigureMinter = "op_weight_msg_configure_minter"
	// TODO: Determine the simulation weight value
	defaultWeightMsgConfigureMinter int = 100

	opWeightMsgRemoveMinter = "op_weight_msg_remove_minter"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRemoveMinter int = 100

	opWeightMsgMint = "op_weight_msg_mint"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMint int = 100

	opWeightMsgBurn = "op_weight_msg_burn"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBurn int = 100

	opWeightMsgBlacklist = "op_weight_msg_blacklist"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBlacklist int = 100

	opWeightMsgUnblacklist = "op_weight_msg_unblacklist"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUnblacklist int = 100

	opWeightMsgPause = "op_weight_msg_pause"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPause int = 100

	opWeightMsgUnpause = "op_weight_msg_unpause"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUnpause int = 100

	opWeightMsgConfigureMinterController = "op_weight_msg_configure_minter_controller"
	// TODO: Determine the simulation weight value
	defaultWeightMsgConfigureMinterController int = 100

	opWeightMsgRemoveMinterController = "op_weight_msg_remove_minter_controller"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRemoveMinterController int = 100
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	// x/fiattokenfactory

	genesis := types.GenesisState{
		MintersList: []types.Minters{
			{
				Address: authtypes.NewModuleAddress("cctp").String(),
			},
		},
		MinterControllerList: []types.MinterController{
			{
				Minter: authtypes.NewModuleAddress("cctp").String(),
			},
		},
		MintingDenom: &types.MintingDenom{Denom: "uusdc"},
	}

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&genesis)

	// x/bank

	bankGenesisBz := simState.GenState[banktypes.ModuleName]
	var bankGenesis banktypes.GenesisState
	simState.Cdc.MustUnmarshalJSON(bankGenesisBz, &bankGenesis)

	bankGenesis.DenomMetadata = append(bankGenesis.DenomMetadata, banktypes.Metadata{
		Description: "USD Coin",
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    "uusdc",
				Exponent: 0,
				Aliases:  []string{"microusdc"},
			},
			{
				Denom:    "usdc",
				Exponent: 6,
				Aliases:  []string{},
			},
		},
		Base:    "uusdc",
		Display: "usdc",
		Name:    "usdc",
		Symbol:  "USDC",
	})

	simState.GenState[banktypes.ModuleName] = simState.Cdc.MustMarshalJSON(&bankGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
// func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
// 	return []simtypes.ParamChange{}
// }

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgUpdateMasterMinter int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateMasterMinter, &weightMsgUpdateMasterMinter, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMasterMinter = defaultWeightMsgUpdateMasterMinter
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateMasterMinter,
		tokenfactorysimulation.SimulateMsgUpdateMasterMinter(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdatePauser int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdatePauser, &weightMsgUpdatePauser, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePauser = defaultWeightMsgUpdatePauser
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePauser,
		tokenfactorysimulation.SimulateMsgUpdatePauser(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateBlacklister int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateBlacklister, &weightMsgUpdateBlacklister, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateBlacklister = defaultWeightMsgUpdateBlacklister
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateBlacklister,
		tokenfactorysimulation.SimulateMsgUpdateBlacklister(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateOwner int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateOwner, &weightMsgUpdateOwner, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateOwner = defaultWeightMsgUpdateOwner
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateOwner,
		tokenfactorysimulation.SimulateMsgUpdateOwner(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgConfigureMinter int
	simState.AppParams.GetOrGenerate(opWeightMsgConfigureMinter, &weightMsgConfigureMinter, nil,
		func(_ *rand.Rand) {
			weightMsgConfigureMinter = defaultWeightMsgConfigureMinter
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgConfigureMinter,
		tokenfactorysimulation.SimulateMsgConfigureMinter(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRemoveMinter int
	simState.AppParams.GetOrGenerate(opWeightMsgRemoveMinter, &weightMsgRemoveMinter, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveMinter = defaultWeightMsgRemoveMinter
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveMinter,
		tokenfactorysimulation.SimulateMsgRemoveMinter(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMint int
	simState.AppParams.GetOrGenerate(opWeightMsgMint, &weightMsgMint, nil,
		func(_ *rand.Rand) {
			weightMsgMint = defaultWeightMsgMint
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMint,
		tokenfactorysimulation.SimulateMsgMint(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBurn int
	simState.AppParams.GetOrGenerate(opWeightMsgBurn, &weightMsgBurn, nil,
		func(_ *rand.Rand) {
			weightMsgBurn = defaultWeightMsgBurn
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBurn,
		tokenfactorysimulation.SimulateMsgBurn(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBlacklist int
	simState.AppParams.GetOrGenerate(opWeightMsgBlacklist, &weightMsgBlacklist, nil,
		func(_ *rand.Rand) {
			weightMsgBlacklist = defaultWeightMsgBlacklist
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBlacklist,
		tokenfactorysimulation.SimulateMsgBlacklist(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUnblacklist int
	simState.AppParams.GetOrGenerate(opWeightMsgUnblacklist, &weightMsgUnblacklist, nil,
		func(_ *rand.Rand) {
			weightMsgUnblacklist = defaultWeightMsgUnblacklist
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUnblacklist,
		tokenfactorysimulation.SimulateMsgUnblacklist(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgPause int
	simState.AppParams.GetOrGenerate(opWeightMsgPause, &weightMsgPause, nil,
		func(_ *rand.Rand) {
			weightMsgPause = defaultWeightMsgPause
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPause,
		tokenfactorysimulation.SimulateMsgPause(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUnpause int
	simState.AppParams.GetOrGenerate(opWeightMsgUnpause, &weightMsgUnpause, nil,
		func(_ *rand.Rand) {
			weightMsgUnpause = defaultWeightMsgUnpause
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUnpause,
		tokenfactorysimulation.SimulateMsgUnpause(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgConfigureMinterController int
	simState.AppParams.GetOrGenerate(opWeightMsgConfigureMinterController, &weightMsgConfigureMinterController, nil,
		func(_ *rand.Rand) {
			weightMsgConfigureMinterController = defaultWeightMsgConfigureMinterController
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgConfigureMinterController,
		tokenfactorysimulation.SimulateMsgConfigureMinterController(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRemoveMinterController int
	simState.AppParams.GetOrGenerate(opWeightMsgRemoveMinterController, &weightMsgRemoveMinterController, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveMinterController = defaultWeightMsgRemoveMinterController
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveMinterController,
		tokenfactorysimulation.SimulateMsgRemoveMinterController(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	return operations
}
