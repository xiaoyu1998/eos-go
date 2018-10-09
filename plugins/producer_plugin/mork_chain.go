package producer_plugin

import (
	"fmt"
	Chain "github.com/eosspark/eos-go/chain"
	"github.com/eosspark/eos-go/chain/types"
	"github.com/eosspark/eos-go/common"
	"github.com/eosspark/eos-go/crypto"
	"github.com/eosspark/eos-go/crypto/ecc"
)

var chain *mockChain

var initPriKey, _ = ecc.NewPrivateKey("5KYZdUEo39z3FPrtuX2QbbwGnNP5zTd7yyr2SC1j299sBCnWjss")
var initPubKey = initPriKey.PublicKey()
var initPriKey2, _ = ecc.NewPrivateKey("5Ja3h2wJNUnNcoj39jDMHGigsazvbGHAeLYEHM5uTwtfUoRDoYP")
var initPubKey2 = initPriKey2.PublicKey()
var eosio = common.AccountName(common.N("eosio"))
var yuanc = common.AccountName(common.N("yuanc"))

type mockChain struct {
	head    *types.BlockState
	pending *types.BlockState
	forkDb  forkDatabase
}

type forkDatabase struct {
	head  *types.BlockState
	index []*types.BlockState
}

func (db *forkDatabase) find(id common.BlockIdType) *types.BlockState {
	for _, n := range db.index {
		if n.ID == id {
			return n
		}
	}
	return nil
}

func (db *forkDatabase) findByNum(num uint32) *types.BlockState {
	for _, n := range db.index {
		if n.BlockNum == num {
			return n
		}
	}
	return nil
}

func (db *forkDatabase) add(n *types.BlockState) *types.BlockState {
	db.index = append(db.index, n)
	db.head = n
	return n
}

func (db *forkDatabase) add2(b *types.SignedBlock, trust bool) *types.BlockState {
	prior := db.find(b.Previous)
	result := types.NewBlockState3(prior.BlockHeaderState, b, trust)
	return db.add(result)
}

func init() {
	fmt.Println(initPubKey)
	fmt.Println(initPubKey2)
	chain = new(mockChain)

	initSchedule := types.ProducerScheduleType{0, []types.ProducerKey{
		{eosio, initPubKey},
		{yuanc, initPubKey2},
	}}

	genHeader := types.BlockHeaderState{}
	genHeader.ActiveSchedule = initSchedule
	genHeader.PendingSchedule = initSchedule
	genHeader.Header.Timestamp = common.NewBlockTimeStamp(common.Now())
	genHeader.ID = genHeader.Header.BlockID()
	genHeader.BlockNum = genHeader.Header.BlockNumber()

	genHeader.ProducerToLastProduced = make(map[common.AccountName]uint32)
	genHeader.ProducerToLastImpliedIrb = make(map[common.AccountName]uint32)

	chain.head = types.NewBlockState(genHeader)
	chain.head.SignedBlock = new(types.SignedBlock)
	chain.head.SignedBlock.SignedBlockHeader = genHeader.Header
	chain.forkDb.add(chain.head)

	fmt.Println("now", common.Now())
	fmt.Println("init", genHeader.Header.Timestamp.ToTimePoint())
}

func (c mockChain) LastIrreversibleBlockNum() uint32 {
	return c.head.DposIrreversibleBlocknum
}

func (c mockChain) HeadBlockState() *types.BlockState {
	return c.head

}
func (c mockChain) HeadBlockTime() common.TimePoint {
	return c.head.Header.Timestamp.ToTimePoint()
}

func (c mockChain) PendingBlockTime() common.TimePoint {
	return c.pending.Header.Timestamp.ToTimePoint()
}

func (c mockChain) HeadBlockNum() uint32 {
	return c.head.BlockNum
}

func (c mockChain) PendingBlockState() *types.BlockState {
	return c.pending
}

func (c mockChain) GetUnappliedTransactions() []*types.TransactionMetadata {
	return make([]*types.TransactionMetadata, 0)
}

func (c mockChain) GetScheduledTransactions() []common.TransactionIdType {
	return make([]common.TransactionIdType, 0)
}

func (c *mockChain) AbortBlock() {
	fmt.Println("abort block...")
	if c.pending != nil {
		c.pending = nil
	}
}

func (c *mockChain) StartBlock(when common.BlockTimeStamp, confirmBlockCount uint16) {
	fmt.Println("start block...")
	chain.pending = types.NewBlockState2(c.head.BlockHeaderState, when)
	chain.pending.SetConfirmed(confirmBlockCount)

}
func (c *mockChain) FinalizeBlock() {
	fmt.Println("finalize block...")
	c.pending.ID = c.pending.Header.BlockID()
}

func (c *mockChain) SignBlock(callback func(sha256 crypto.Sha256) ecc.Signature) {
	fmt.Println("sign block...")
	p := c.pending
	p.Sign(callback)
	p.SignedBlock.SignedBlockHeader = p.Header
}

func (c *mockChain) CommitBlock(addToForkDb bool) {
	fmt.Println("commit block...")

	if addToForkDb {
		c.pending.Validated = true
		c.forkDb.add(c.pending)
		c.head = c.forkDb.head
	}

	//c.pending = nil
}

func (c *mockChain) PushTransaction(trx *types.TransactionMetadata, deadline common.TimePoint) *types.TransactionTrace {
	//c.pending.SignedBlock.Transactions = append(c.pending.SignedBlock.Transactions, )
	c.pending.Trxs = append(c.pending.Trxs, trx)
	return nil
}
func (c *mockChain) PushScheduledTransaction(trx common.TransactionIdType, deadline common.TimePoint) *types.TransactionTrace {
	return nil
}

func (c *mockChain) PushReceipt(trx interface{}) types.TransactionReceipt {
	//c.pending.SignedBlock.Transactions = append(c.pending.SignedBlock.Transactions, )
	return types.TransactionReceipt{}
}

func (c *mockChain) PushBlock(b *types.SignedBlock) {
	c.forkDb.add2(b, false)
	if c.GetReadMode() != Chain.IRREVERSIBLE {
		c.MaybeSwitchForks()
	}

}

func (c *mockChain) MaybeSwitchForks() {
	newHead := c.forkDb.head

	if newHead.Header.Previous == c.head.ID {
		c.ApplyBlock(newHead.SignedBlock)

	} else if newHead.ID != c.head.ID {
		fmt.Println(" newHead.ID != c.head.ID ")
	}
}

func (c *mockChain) ApplyBlock(b *types.SignedBlock) {
	c.StartBlock(b.Timestamp, b.Confirmed)
	c.FinalizeBlock()

	c.pending.Header.ProducerSignature = b.ProducerSignature
	c.pending.SignedBlock.SignedBlockHeader = c.pending.Header

	c.CommitBlock(false)
}

func (c *mockChain) FetchBlockById(id common.BlockIdType) *types.SignedBlock {
	state := c.forkDb.find(id)
	if state != nil {
		return state.SignedBlock
	}
	bptr := c.FetchBlockByNumber(types.NumFromID(id))
	if bptr != nil && bptr.BlockID() == id {
		return bptr
	}
	return nil
}

func (c *mockChain) FetchBlockByNumber(num uint32) *types.SignedBlock {
	state := c.forkDb.findByNum(num)
	if state != nil {
		return state.SignedBlock
	}
	return nil
}

func (c *mockChain) IsKnownUnexpiredTransaction(id common.TransactionIdType) bool {
	return false
}

func (c *mockChain) DropUnappliedTransaction(trx *types.TransactionMetadata) {}

func (c *mockChain) GetReadMode() Chain.DBReadMode {
	return Chain.DBReadMode(Chain.SPECULATIVE)
}