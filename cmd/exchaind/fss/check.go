package fss

import (
	"bytes"
	"fmt"
	"log"

	"github.com/okex/exchain/cmd/exchaind/base"
	"github.com/okex/exchain/libs/iavl"
	dbm "github.com/okex/exchain/libs/tm-db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// checkCmd represents the create command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check with the fast index with IAVL original nodes",
	Long: `Check fast index with IAVL original nodes:
This command is a tool to check the IAVL fast index.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		iavl.SetEnableFastStorage(true)
		storeKeys := getStoreKeys()
		outputModules(storeKeys)

		return check(storeKeys)
	},
}

func init() {
	fssCmd.AddCommand(checkCmd)
}

func check(storeKeys []string) error {
	dataDir := viper.GetString(flagDataDir)
	dbBackend := viper.GetString(flagDBBackend)
	db, err := base.OpenDB(dataDir+base.AppDBName, dbm.BackendType(dbBackend))
	if err != nil {
		return fmt.Errorf("error opening dir %v backend %v DB: %w", dataDir, dbBackend, err)
	}
	defer db.Close()

	for _, key := range storeKeys {
		log.Printf("Checking.... %v\n", key)
		prefix := []byte(fmt.Sprintf("s/k:%s/", key))

		prefixDB := dbm.NewPrefixDB(db, prefix)

		mutableTree, err := iavl.NewMutableTree(prefixDB, 0)
		if err != nil {
			return err
		}
		if err := checkIndex(mutableTree); err != nil {
			return fmt.Errorf("%v iavl fast index not match %v", key, err.Error())
		}
	}

	return nil
}

func checkIndex(mutableTree *iavl.MutableTree) error {
	fastIterator := mutableTree.Iterator(nil, nil, true)
	defer fastIterator.Close()
	iterator := iavl.NewIterator(nil, nil, true, mutableTree.ImmutableTree)
	defer iterator.Close()

	for fastIterator.Valid() && iterator.Valid() {
		if bytes.Compare(fastIterator.Key(), iterator.Key()) != 0 ||
			bytes.Compare(fastIterator.Value(), iterator.Value()) != 0 {
			return fmt.Errorf("fast index key:%v value:%v, iavl node key:%v iavl node value:%v",
				fastIterator.Key(), fastIterator.Value(), iterator.Key(), iterator.Value())
		}
		fastIterator.Next()
		iterator.Next()
	}

	if fastIterator.Valid() {
		return fmt.Errorf("fast index key:%v value:%v", fastIterator.Key(), fastIterator.Value())
	}

	if iterator.Valid() {
		return fmt.Errorf("iavl node key:%v value:%v", iterator.Key(), iterator.Value())
	}

	return nil
}
