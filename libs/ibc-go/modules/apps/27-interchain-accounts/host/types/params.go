package types

import (
	"fmt"
	"strings"

	paramtypes "github.com/okex/exchain/x/params"
)

const (
	// DefaultHostEnabled is the default value for the host param (set to true)
	DefaultHostEnabled = true
)

var (
	// KeyHostEnabled is the store key for HostEnabled Params
	KeyHostEnabled = []byte("HostEnabled")
	// KeyAllowMessages is the store key for the AllowMessages Params
	KeyAllowMessages = []byte("AllowMessages")
)

const (
	bankMsgSend = "/cosmos.bank.v1beta1.MsgSend"
)

var (
	DefaultAllowMessages []string = []string{
		bankMsgSend,
	}
)

// ParamKeyTable type declaration for parameters
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new parameter configuration for the host submodule
func NewParams(enableHost bool, allowMsgs []string) Params {
	return Params{
		HostEnabled:   enableHost,
		AllowMessages: allowMsgs,
	}
}

// DefaultParams is the default parameter configuration for the host submodule
func DefaultParams() Params {
	return NewParams(DefaultHostEnabled, DefaultAllowMessages)
}

// Validate validates all host submodule parameters
func (p Params) Validate() error {
	if err := validateEnabled(p.HostEnabled); err != nil {
		return err
	}

	if err := validateAllowlist(p.AllowMessages); err != nil {
		return err
	}

	return nil
}

// ParamSetPairs implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyHostEnabled, p.HostEnabled, validateEnabled),
		paramtypes.NewParamSetPair(KeyAllowMessages, p.AllowMessages, validateAllowlist),
	}
}

func validateEnabled(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateAllowlist(i interface{}) error {
	allowMsgs, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	for _, typeURL := range allowMsgs {
		if strings.TrimSpace(typeURL) == "" {
			return fmt.Errorf("parameter must not contain empty strings: %s", allowMsgs)
		}
	}

	return nil
}
