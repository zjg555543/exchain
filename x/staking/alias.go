// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/okex/exchain/x/staking/keeper
// ALIASGEN: github.com/okex/exchain/x/staking/types
// ALIASGEN: github.com/okex/exchain/x/staking/exported
package staking

import (
	"github.com/okex/exchain/x/staking/exported"
	"github.com/okex/exchain/x/staking/keeper"
	"github.com/okex/exchain/x/staking/types"
)

const (
	DefaultParamspace = keeper.DefaultParamspace
	ModuleName        = types.ModuleName
	StoreKey          = types.StoreKey
	TStoreKey         = types.TStoreKey
	QuerierRoute      = types.QuerierRoute
	RouterKey         = types.RouterKey
	NotBondedPoolName = types.NotBondedPoolName
	BondedPoolName    = types.BondedPoolName
	QueryParameters   = types.QueryParameters
)

var (
	// functions aliases
	NewKeeper                          = keeper.NewKeeper
	NewQuerier                         = keeper.NewQuerier
	RegisterCodec                      = types.RegisterCodec
	NewCommission                      = types.NewCommission
	NewCommissionRates                 = types.NewCommissionRates
	ErrNoValidatorFound                = types.ErrNoValidatorFound
	ErrValidatorOwnerExists            = types.ErrValidatorOwnerExists
	ErrValidatorPubKeyExists           = types.ErrValidatorPubKeyExists
	ErrValidatorPubKeyTypeNotSupported = types.ErrValidatorPubKeyTypeNotSupported
	ErrBadDenom                        = types.ErrBadDenom
	DefaultGenesisState                = types.DefaultGenesisState
	NewMultiStakingHooks               = types.NewMultiStakingHooks
	GetValidatorsByPowerIndexKey       = types.GetValidatorsByPowerIndexKey
	NewMsgCreateValidator              = types.NewMsgCreateValidator
	NewMsgEditValidator                = types.NewMsgEditValidator
	NewMsgDestroyValidator             = types.NewMsgDestroyValidator
	NewMsgDeposit                      = types.NewMsgDeposit
	NewMsgWithdraw                     = types.NewMsgWithdraw
	DefaultParams                      = types.DefaultParams
	NewValidator                       = types.NewValidator
	NewDescription                     = types.NewDescription
	NewMsgAddShares                    = types.NewMsgAddShares
	NewGenesisState                    = types.NewGenesisState
	DelegatorAddSharesInvariant        = keeper.DelegatorAddSharesInvariant
	NewMsgRegProxy                     = types.NewMsgRegProxy
	NewMsgBindProxy                    = types.NewMsgBindProxy
	NewMsgUnbindProxy                  = types.NewMsgUnbindProxy

	// variable aliases
	ModuleCdc     = types.ModuleCdc
	ValidatorsKey = types.ValidatorsKey
)

type (
	Keeper                    = keeper.Keeper
	GenesisState              = types.GenesisState
	Validator                 = types.Validator
	Validators                = types.Validators
	ValidatorExport           = types.ValidatorExported
	Description               = types.Description
	ValidatorI                = exported.ValidatorI
	Delegator                 = types.Delegator
	UndelegationInfo          = types.UndelegationInfo
	ProxyDelegatorKeyExported = types.ProxyDelegatorKeyExported
	SharesResponses           = types.SharesResponses
)
