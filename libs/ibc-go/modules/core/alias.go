package ibc

import (
	"github.com/okex/exchain/libs/ibc-go/modules/core/keeper"
	"github.com/okex/exchain/libs/ibc-go/modules/core/types"
)

type (
	Keeper = keeper.FacadedKeeper
)

const ()

var (
	NewKeeper           = keeper.NewKeeper
	NewV4Keeper         = keeper.NewV4Keeper
	NewFacadedKeeper    = keeper.NewFacadedKeeper
	ModuleCdc           = types.ModuleCdc
	DefaultGenesisState = types.DefaultGenesisState
)

var (
	DefaultSelectorFactory = keeper.DefaultFactory
)

const (
	IBCV4 = 4
	IBCV2 = 2
)
