package ethkeystore

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/okex/exchain/app/crypto/ethsecp256k1"
	tmcrypto "github.com/okex/exchain/libs/tendermint/crypto"
)

// CreateKeystoreByTmKey  create a eth keystore by accountname from keybase
func CreateKeystoreByTmKey(privKey tmcrypto.PrivKey, dir, encryptPassword string) (string, error) {
	// dir must be absolute
	if !filepath.IsAbs(dir) {
		return "", fmt.Errorf("invalid directory")
	}
	// converts tendermint  key to ethereum key
	ethKey, err := EncodeTmKeyToEthKey(privKey)
	if err != nil {
		return "", fmt.Errorf("invalid private key type, must be Ethereum key: %T", privKey)
	}

	// export Key to keystore file
	// if filename isn't set ,use default ethereum name
	addr := common.BytesToAddress(privKey.PubKey().Address())
	fileName := filepath.Join(dir, keyFileName(addr))
	err = ExportKeyStoreFile(ethKey, encryptPassword, fileName)
	return fileName, err
}

// EncodeTmKeyToEthKey  transfer tendermint key  to a ethereum key
func EncodeTmKeyToEthKey(privKey tmcrypto.PrivKey) (*ecdsa.PrivateKey, error) {
	// Converts key to Ethermint secp256 implementation
	emintKey, ok := privKey.(ethsecp256k1.PrivKey)
	if !ok {
		return nil, fmt.Errorf("invalid private key type, must be Ethereum key: %T", privKey)
	}

	return emintKey.ToECDSA(), nil
}

// ExportKeyStoreFile Export Key to  keystore file
func ExportKeyStoreFile(privateKeyECDSA *ecdsa.PrivateKey, encryptPassword, fileName string) error {
	//new keystore key
	ethKey, err := newEthKeyFromECDSA(privateKeyECDSA)
	if err != nil {
		return err
	}
	// encrypt Key to get keystore file
	content, err := keystore.EncryptKey(ethKey, encryptPassword, keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		return fmt.Errorf("failed to encrypt key: %s", err.Error())
	}

	// write to keystore file
	err = ioutil.WriteFile(fileName, content, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write keystore: %s", err.Error())
	}
	return nil
}

// newEthKeyFromECDSA new eth.keystore Key
func newEthKeyFromECDSA(privateKeyECDSA *ecdsa.PrivateKey) (*keystore.Key, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("Could not create random uuid: %v", err)
	}
	key := &keystore.Key{
		Id:         id,
		Address:    ethcrypto.PubkeyToAddress(privateKeyECDSA.PublicKey),
		PrivateKey: privateKeyECDSA,
	}
	return key, nil
}

//keyFileName return the default keystore file name in the ethereum
func keyFileName(keyAddr common.Address) string {
	ts := time.Now().UTC()
	return fmt.Sprintf("UTC--%s--%s", toISO8601(ts), hex.EncodeToString(keyAddr[:]))
}

func toISO8601(t time.Time) string {
	var tz string
	name, offset := t.Zone()
	if name == "UTC" {
		tz = "Z"
	} else {
		tz = fmt.Sprintf("%03d00", offset/3600)
	}
	return fmt.Sprintf("%04d-%02d-%02dT%02d-%02d-%02d.%09d%s",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), tz)
}
