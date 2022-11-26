//go:build rocksdb
// +build rocksdb

package db

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/cosmos/gorocksdb"
	"github.com/spf13/viper"
)

func init() {
	dbCreator := func(name string, dir string) (DB, error) {
		return NewRocksDB(name, dir)
	}
	registerDBCreator(RocksDBBackend, dbCreator, false)
}

// RocksDB is a RocksDB backend.
type RocksDB struct {
	db     *gorocksdb.DB
	ro     *gorocksdb.ReadOptions
	wo     *gorocksdb.WriteOptions
	woSync *gorocksdb.WriteOptions
}

var _ DB = (*RocksDB)(nil)

const (
	blockSize      = "block_size"
	blockCache     = "block_cache"
	statistics     = "statistics"
	maxOpenFiles   = "max_open_files"
	mmapRead       = "allow_mmap_reads"
	mmapWrite      = "allow_mmap_writes"
	bloomFilter    = "bloom_filter"
	newFormatBloom = "bloom_filter_new"
	disableComp    = "disable_comp"
	writeBuff      = "write_buff_size"
	level0Trigger  = "level0_trigger"
	parallelism    = "parallelism"
	backComp       = "back_comp"
)

func NewRocksDB(name string, dir string) (*RocksDB, error) {
	// default rocksdb option, good enough for most cases, including heavy workloads.
	// 1GB table cache, 512MB write buffer(may use 50% more on heavy workloads).
	// compression: snappy as default, need to -lsnappy to enable.
	params := parseOptParams(viper.GetString(FlagRocksdbOpts))

	bbto := gorocksdb.NewDefaultBlockBasedTableOptions()
	if v, ok := params[blockSize]; ok {
		size, err := toBytes(v)
		if err != nil {
			panic(fmt.Sprintf("Invalid options parameter %s: %s", blockSize, err))
		}
		bbto.SetBlockSize(int(size))
	}
	bbto.SetBlockCache(gorocksdb.NewLRUCache(1 << 30))
	if v, ok := params[blockCache]; ok {
		cache, err := toBytes(v)
		if err != nil {
			panic(fmt.Sprintf("Invalid options parameter %s: %s", blockCache, err))
		}
		bbto.SetBlockCache(gorocksdb.NewLRUCache(cache))
	}

	bitsPerKey := 10
	if v, ok := params[bloomFilter]; ok {
		bit, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			bitsPerKey = int(bit)
		}
	}

	if _, ok := params[newFormatBloom]; ok {
		fmt.Println("*****lyh***** NewBloomFilterFull", bitsPerKey)
		bbto.SetFilterPolicy(gorocksdb.NewBloomFilterFull(bitsPerKey))
	} else {
		fmt.Println("*****lyh***** NewBloomFilter", bitsPerKey)
		bbto.SetFilterPolicy(gorocksdb.NewBloomFilter(bitsPerKey))
	}

	//bbto.SetCacheIndexAndFilterBlocks(true)
	//bbto.SetCacheIndexAndFilterBlocksWithHighPriority(true)

	opts := gorocksdb.NewDefaultOptions()
	opts.SetBlockBasedTableFactory(bbto)
	opts.SetCreateIfMissing(true)

	if v, ok := params[parallelism]; ok {
		size, err := strconv.Atoi(v)
		if err == nil {
			fmt.Println("*****lyh***** parallelism", size, runtime.NumCPU())
			opts.IncreaseParallelism(size)
		}
	} else {
		opts.IncreaseParallelism(runtime.NumCPU())
	}

	if v, ok := params[backComp]; ok {
		size, err := strconv.Atoi(v)
		if err == nil {
			fmt.Println("*****lyh***** backComp", size)
			opts.SetMaxBackgroundCompactions(size)
		}
	}

	if v, ok := params[statistics]; ok {
		enable, err := strconv.ParseBool(v)
		if err != nil {
			panic(fmt.Sprintf("Invalid options parameter %s: %s", statistics, err))
		}
		if enable {
			opts.EnableStatistics()
		}
	}

	opts.SetMaxOpenFiles(-1)
	if v, ok := params[maxOpenFiles]; ok {
		maxOpenFiles, err := strconv.Atoi(v)
		if err != nil {
			panic(fmt.Sprintf("Invalid options parameter %s: %s", maxOpenFiles, err))
		}
		opts.SetMaxOpenFiles(maxOpenFiles)
	}

	opts.SetAllowMmapReads(false)
	if v, ok := params[mmapRead]; ok {
		enable, err := strconv.ParseBool(v)
		if err != nil {
			panic(fmt.Sprintf("Invalid options parameter %s: %s", mmapRead, err))
		}
		opts.SetAllowMmapReads(enable)
	}

	if v, ok := params[mmapWrite]; ok {
		enable, err := strconv.ParseBool(v)
		if err != nil {
			panic(fmt.Sprintf("Invalid options parameter %s: %s", mmapWrite, err))
		}
		if enable {
			opts.SetAllowMmapWrites(enable)
		}
	}

	if _, ok := params[disableComp]; ok {
		fmt.Println("*****lyh***** disableComp")
		opts.SetDisableAutoCompactions(true)
	}

	if v, ok := params[writeBuff]; ok {
		size, err := strconv.Atoi(v)
		if err == nil {
			size = size * 1024 * 1024 * 128
			fmt.Println("*****lyh***** writeBuffsize", size)
			opts.SetWriteBufferSize(size)
			opts.OptimizeLevelStyleCompaction(uint64(size * 4))
		}
	} else {
		opts.OptimizeLevelStyleCompaction(512 * 1024 * 1024)
	}

	if v, ok := params[level0Trigger]; ok {
		size, err := strconv.Atoi(v)
		if err == nil {
			fmt.Println("*****lyh***** SetLevel0FileNumCompactionTrigger", size)
			opts.SetLevel0FileNumCompactionTrigger(size)
		}
		opts.SetMaxBytesForLevelBase(1024 * 10 * 1024 * 1024)
	}

	// 1.5GB maximum memory use for writebuffer.

	return NewRocksDBWithOptions(name, dir, opts)
}

func NewRocksDBWithOptions(name string, dir string, opts *gorocksdb.Options) (*RocksDB, error) {
	dbPath := filepath.Join(dir, name+".db")
	db, err := gorocksdb.OpenDb(opts, dbPath)
	if err != nil {
		return nil, err
	}
	ro := gorocksdb.NewDefaultReadOptions()
	wo := gorocksdb.NewDefaultWriteOptions()
	woSync := gorocksdb.NewDefaultWriteOptions()
	woSync.SetSync(true)
	database := &RocksDB{
		db:     db,
		ro:     ro,
		wo:     wo,
		woSync: woSync,
	}
	return database, nil
}

// Get implements DB.
func (db *RocksDB) Get(key []byte) ([]byte, error) {
	key = nonNilBytes(key)
	res, err := db.db.Get(db.ro, key)
	if err != nil {
		return nil, err
	}
	return moveSliceToBytes(res), nil
}

func (db *RocksDB) GetUnsafeValue(key []byte, processor UnsafeValueProcessor) (interface{}, error) {
	key = nonNilBytes(key)
	res, err := db.db.Get(db.ro, key)
	if err != nil {
		return nil, err
	}
	defer res.Free()
	if !res.Exists() {
		return processor(nil)
	}
	return processor(res.Data())
}

// Has implements DB.
func (db *RocksDB) Has(key []byte) (bool, error) {
	bytes, err := db.Get(key)
	if err != nil {
		return false, err
	}
	return bytes != nil, nil
}

// Set implements DB.
func (db *RocksDB) Set(key []byte, value []byte) error {
	key = nonNilBytes(key)
	value = nonNilBytes(value)
	err := db.db.Put(db.wo, key, value)
	if err != nil {
		return err
	}
	return nil
}

// SetSync implements DB.
func (db *RocksDB) SetSync(key []byte, value []byte) error {
	key = nonNilBytes(key)
	value = nonNilBytes(value)
	err := db.db.Put(db.woSync, key, value)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements DB.
func (db *RocksDB) Delete(key []byte) error {
	key = nonNilBytes(key)
	err := db.db.Delete(db.wo, key)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSync implements DB.
func (db *RocksDB) DeleteSync(key []byte) error {
	key = nonNilBytes(key)
	err := db.db.Delete(db.woSync, key)
	if err != nil {
		return nil
	}
	return nil
}

func (db *RocksDB) DB() *gorocksdb.DB {
	return db.db
}

// Close implements DB.
func (db *RocksDB) Close() error {
	db.ro.Destroy()
	db.wo.Destroy()
	db.woSync.Destroy()
	db.db.Close()
	return nil
}

// Print implements DB.
func (db *RocksDB) Print() error {
	itr, err := db.Iterator(nil, nil)
	if err != nil {
		return err
	}
	defer itr.Close()
	for ; itr.Valid(); itr.Next() {
		key := itr.Key()
		value := itr.Value()
		fmt.Printf("[%X]:\t[%X]\n", key, value)
	}
	return nil
}

// Stats implements DB.
func (db *RocksDB) Stats() map[string]string {
	keys := []string{"rocksdb.stats"}
	stats := make(map[string]string, len(keys))
	for _, key := range keys {
		stats[key] = db.db.GetProperty(key)
	}
	return stats
}

// NewBatch implements DB.
func (db *RocksDB) NewBatch() Batch {
	return NewRocksDBBatch(db)
}

// Iterator implements DB.
func (db *RocksDB) Iterator(start, end []byte) (Iterator, error) {
	itr := db.db.NewIterator(db.ro)
	return NewRocksDBIterator(itr, start, end, false), nil
}

// ReverseIterator implements DB.
func (db *RocksDB) ReverseIterator(start, end []byte) (Iterator, error) {
	itr := db.db.NewIterator(db.ro)
	return NewRocksDBIterator(itr, start, end, true), nil
}

func (db *RocksDB) Compact() error {
	db.DB().CompactRange(gorocksdb.Range{})
	return nil
}
