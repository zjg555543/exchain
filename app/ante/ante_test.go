package ante_test

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	ethcmn "github.com/ethereum/go-ethereum/common"

	abci "github.com/okex/exchain/libs/tendermint/abci/types"
	tmcrypto "github.com/okex/exchain/libs/tendermint/crypto"

	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"

	"github.com/okex/exchain/app"
	"github.com/okex/exchain/app/ante"
	"github.com/okex/exchain/app/types"
	evmtypes "github.com/okex/exchain/x/evm/types"
)

func requireValidTx(
	t *testing.T, anteHandler sdk.AnteHandler, ctx sdk.Context, tx sdk.Tx, sim bool,
) {
	_, err := anteHandler(ctx, tx, sim)
	require.NoError(t, err)
}

func requireInvalidTx(
	t *testing.T, anteHandler sdk.AnteHandler, ctx sdk.Context,
	tx sdk.Tx, sim bool,
) {
	_, err := anteHandler(ctx, tx, sim)
	require.Error(t, err)
}

func (suite *AnteTestSuite) TestValidEthTx() {
	suite.ctx.SetBlockHeight(1)

	addr1, priv1 := newTestAddrKey()
	addr2, _ := newTestAddrKey()

	acc1 := suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, addr1)
	_ = acc1.SetCoins(newTestCoins())
	suite.app.AccountKeeper.SetAccount(suite.ctx, acc1)

	acc2 := suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, addr2)
	_ = acc2.SetCoins(newTestCoins())
	suite.app.AccountKeeper.SetAccount(suite.ctx, acc2)

	// require a valid Ethereum tx to pass
	to := ethcmn.BytesToAddress(addr2.Bytes())
	amt := big.NewInt(32)
	gas := big.NewInt(20)
	ethMsg := evmtypes.NewMsgEthereumTx(0, &to, amt, 22000, gas, []byte("test"))

	tx, err := newTestEthTx(suite.ctx, ethMsg, priv1)
	suite.Require().NoError(err)
	requireValidTx(suite.T(), suite.anteHandler, suite.ctx, tx, false)
}

func (suite *AnteTestSuite) TestValidTx() {
	suite.ctx.SetBlockHeight(1)

	addr1, priv1 := newTestAddrKey()
	addr2, priv2 := newTestAddrKey()

	acc1 := suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, addr1)
	_ = acc1.SetCoins(newTestCoins())
	suite.app.AccountKeeper.SetAccount(suite.ctx, acc1)

	acc2 := suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, addr2)
	_ = acc2.SetCoins(newTestCoins())
	suite.app.AccountKeeper.SetAccount(suite.ctx, acc2)

	// require a valid SDK tx to pass
	fee := newTestStdFee()
	msg1 := newTestMsg(addr1, addr2)
	msgs := []sdk.Msg{msg1}

	privKeys := []tmcrypto.PrivKey{priv1, priv2}
	accNums := []uint64{acc1.GetAccountNumber(), acc2.GetAccountNumber()}
	accSeqs := []uint64{acc1.GetSequence(), acc2.GetSequence()}

	tx := newTestSDKTx(suite.ctx, msgs, privKeys, accNums, accSeqs, fee)

	requireValidTx(suite.T(), suite.anteHandler, suite.ctx, tx, false)
}

func (suite *AnteTestSuite) TestSDKInvalidSigs() {
	suite.ctx.SetBlockHeight(1)

	addr1, priv1 := newTestAddrKey()
	addr2, priv2 := newTestAddrKey()
	addr3, priv3 := newTestAddrKey()

	acc1 := suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, addr1)
	_ = acc1.SetCoins(newTestCoins())
	suite.app.AccountKeeper.SetAccount(suite.ctx, acc1)

	acc2 := suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, addr2)
	_ = acc2.SetCoins(newTestCoins())
	suite.app.AccountKeeper.SetAccount(suite.ctx, acc2)

	fee := newTestStdFee()
	msg1 := newTestMsg(addr1, addr2)

	// require validation failure with no signers
	msgs := []sdk.Msg{msg1}

	privKeys := []tmcrypto.PrivKey{}
	accNums := []uint64{acc1.GetAccountNumber(), acc2.GetAccountNumber()}
	accSeqs := []uint64{acc1.GetSequence(), acc2.GetSequence()}

	tx := newTestSDKTx(suite.ctx, msgs, privKeys, accNums, accSeqs, fee)
	requireInvalidTx(suite.T(), suite.anteHandler, suite.ctx, tx, false)

	// require validation failure with invalid number of signers
	msgs = []sdk.Msg{msg1}

	privKeys = []tmcrypto.PrivKey{priv1}
	accNums = []uint64{acc1.GetAccountNumber(), acc2.GetAccountNumber()}
	accSeqs = []uint64{acc1.GetSequence(), acc2.GetSequence()}

	tx = newTestSDKTx(suite.ctx, msgs, privKeys, accNums, accSeqs, fee)
	requireInvalidTx(suite.T(), suite.anteHandler, suite.ctx, tx, false)

	// require validation failure with an invalid signer
	msg2 := newTestMsg(addr1, addr3)
	msgs = []sdk.Msg{msg1, msg2}

	privKeys = []tmcrypto.PrivKey{priv1, priv2, priv3}
	accNums = []uint64{acc1.GetAccountNumber(), acc2.GetAccountNumber(), 0}
	accSeqs = []uint64{acc1.GetSequence(), acc2.GetSequence(), 0}

	tx = newTestSDKTx(suite.ctx, msgs, privKeys, accNums, accSeqs, fee)
	requireInvalidTx(suite.T(), suite.anteHandler, suite.ctx, tx, false)
}

func (suite *AnteTestSuite) TestSDKInvalidAcc() {
	suite.ctx.SetBlockHeight(1)

	addr1, priv1 := newTestAddrKey()

	acc1 := suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, addr1)
	_ = acc1.SetCoins(newTestCoins())
	suite.app.AccountKeeper.SetAccount(suite.ctx, acc1)

	fee := newTestStdFee()
	msg1 := newTestMsg(addr1)
	msgs := []sdk.Msg{msg1}
	privKeys := []tmcrypto.PrivKey{priv1}

	// require validation failure with invalid account number
	accNums := []uint64{1}
	accSeqs := []uint64{acc1.GetSequence()}

	tx := newTestSDKTx(suite.ctx, msgs, privKeys, accNums, accSeqs, fee)
	requireInvalidTx(suite.T(), suite.anteHandler, suite.ctx, tx, false)

	// require validation failure with invalid sequence (nonce)
	accNums = []uint64{acc1.GetAccountNumber()}
	accSeqs = []uint64{1}

	tx = newTestSDKTx(suite.ctx, msgs, privKeys, accNums, accSeqs, fee)
	requireInvalidTx(suite.T(), suite.anteHandler, suite.ctx, tx, false)
}

func (suite *AnteTestSuite) TestEthInvalidSig() {
	suite.ctx.SetBlockHeight(1)

	_, priv1 := newTestAddrKey()
	addr2, _ := newTestAddrKey()
	to := ethcmn.BytesToAddress(addr2.Bytes())
	amt := big.NewInt(32)
	gas := big.NewInt(20)
	ethMsg := evmtypes.NewMsgEthereumTx(0, &to, amt, 22000, gas, []byte("test"))

	tx, err := newTestEthTx(suite.ctx, ethMsg, priv1)
	suite.Require().NoError(err)

	ctx := suite.ctx
	ctx.SetChainID("ethermint-4")
	requireInvalidTx(suite.T(), suite.anteHandler, ctx, tx, false)
}

func (suite *AnteTestSuite) TestEthInvalidNonce() {

	suite.ctx.SetBlockHeight(1)

	addr1, priv1 := newTestAddrKey()
	addr2, _ := newTestAddrKey()

	acc := suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, addr1)
	err := acc.SetSequence(10)
	suite.Require().NoError(err)
	_ = acc.SetCoins(newTestCoins())
	suite.app.AccountKeeper.SetAccount(suite.ctx, acc)

	// require a valid Ethereum tx to pass
	to := ethcmn.BytesToAddress(addr2.Bytes())
	amt := big.NewInt(32)
	gas := big.NewInt(20)
	ethMsg := evmtypes.NewMsgEthereumTx(0, &to, amt, 22000, gas, []byte("test"))

	tx, err := newTestEthTx(suite.ctx, ethMsg, priv1)
	suite.Require().NoError(err)
	requireInvalidTx(suite.T(), suite.anteHandler, suite.ctx, tx, false)
}

func (suite *AnteTestSuite) TestEthInsufficientBalance() {
	suite.ctx.SetBlockHeight(1)

	addr1, priv1 := newTestAddrKey()
	addr2, _ := newTestAddrKey()

	acc := suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, addr1)
	suite.app.AccountKeeper.SetAccount(suite.ctx, acc)

	// require a valid Ethereum tx to pass
	to := ethcmn.BytesToAddress(addr2.Bytes())
	amt := big.NewInt(32)
	gas := big.NewInt(20)
	ethMsg := evmtypes.NewMsgEthereumTx(0, &to, amt, 22000, gas, []byte("test"))

	tx, err := newTestEthTx(suite.ctx, ethMsg, priv1)
	suite.Require().NoError(err)
	requireInvalidTx(suite.T(), suite.anteHandler, suite.ctx, tx, false)
}

func (suite *AnteTestSuite) TestEthInvalidIntrinsicGas() {
	suite.ctx.SetBlockHeight(1)

	addr1, priv1 := newTestAddrKey()
	addr2, _ := newTestAddrKey()

	acc := suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, addr1)
	_ = acc.SetCoins(newTestCoins())
	suite.app.AccountKeeper.SetAccount(suite.ctx, acc)

	// require a valid Ethereum tx to pass
	to := ethcmn.BytesToAddress(addr2.Bytes())
	amt := big.NewInt(32)
	gas := big.NewInt(20)
	gasLimit := uint64(1000)
	ethMsg := evmtypes.NewMsgEthereumTx(0, &to, amt, gasLimit, gas, []byte("test"))

	tx, err := newTestEthTx(suite.ctx, ethMsg, priv1)
	suite.Require().NoError(err)
	requireInvalidTx(suite.T(), suite.anteHandler, suite.ctx.WithIsCheckTx(true), tx, false)
}

func (suite *AnteTestSuite) TestEthInvalidMempoolFees() {
	// setup app with checkTx = true
	suite.app = app.Setup(true)
	suite.ctx = suite.app.BaseApp.NewContext(true, abci.Header{Height: 1, ChainID: "ethermint-3", Time: time.Now().UTC()})
	suite.app.EvmKeeper.SetParams(suite.ctx, evmtypes.DefaultParams())

	suite.anteHandler = ante.NewAnteHandler(suite.app.AccountKeeper, suite.app.EvmKeeper, suite.app.SupplyKeeper, nil, suite.app.WasmHandler, suite.app.IBCKeeper)
	suite.ctx.SetMinGasPrices(sdk.NewDecCoins(sdk.NewDecCoinFromDec(types.NativeToken, sdk.NewDecFromBigIntWithPrec(big.NewInt(500000), sdk.Precision))))
	addr1, priv1 := newTestAddrKey()
	addr2, _ := newTestAddrKey()

	acc := suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, addr1)
	_ = acc.SetCoins(newTestCoins())
	suite.app.AccountKeeper.SetAccount(suite.ctx, acc)

	// require a valid Ethereum tx to pass
	to := ethcmn.BytesToAddress(addr2.Bytes())
	amt := big.NewInt(32)
	gas := big.NewInt(20)
	ethMsg := evmtypes.NewMsgEthereumTx(0, &to, amt, 22000, gas, []byte("payload"))

	tx, err := newTestEthTx(suite.ctx, ethMsg, priv1)
	suite.Require().NoError(err)
	requireInvalidTx(suite.T(), suite.anteHandler, suite.ctx, tx, false)
}

// TestCase represents a test case used in test tables.
type TestCase struct {
	desc     string
	simulate bool
	expPass  bool
}

func (suite *AnteTestSuite) TestAnteHandlerSequences() {
	addr, priv := newTestAddrKey()
	acc := suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, addr)
	_ = acc.SetCoins(newTestCoins())
	suite.app.AccountKeeper.SetAccount(suite.ctx, acc)

	testCases := []TestCase{
		{
			"good ibctx with right sequence",
			false,
			true,
		},
		{
			"bad ibctx with wrong sequence (replay protected)",
			false,
			false,
		},
	}
	ibcTx := mockIbcTx([]uint64{acc.GetAccountNumber()}, []uint64{acc.GetSequence()}, priv, suite.ctx.ChainID(), addr)
	suite.Require().NotNil(ibcTx)
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.desc), func() {
			if tc.expPass {
				requireValidTx(suite.T(), suite.anteHandler, suite.ctx, *ibcTx, false)
			} else {
				requireInvalidTx(suite.T(), suite.anteHandler, suite.ctx, *ibcTx, false)
			}
		})
	}
}
