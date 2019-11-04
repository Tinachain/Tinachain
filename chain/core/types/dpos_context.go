package types

import (
	_ "bytes"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/Tinachain/Tina/chain/boker/protocol"
	"github.com/Tinachain/Tina/chain/common"
	"github.com/Tinachain/Tina/chain/crypto/sha3"
	"github.com/Tinachain/Tina/chain/ethdb"
	"github.com/Tinachain/Tina/chain/log"
	"github.com/Tinachain/Tina/chain/rlp"
	"github.com/Tinachain/Tina/chain/trie"
)

type DposContext struct {
	epochTrie     *trie.Trie
	validatorTrie *trie.Trie
	voteTrie      *trie.Trie
	db            ethdb.Database
}

func NewEpochTrie(root common.Hash, db ethdb.Database) (*trie.Trie, error) {

	return trie.NewTrieWithPrefix(root, protocol.EpochPrefix, db)
}

func NewValidatorTrie(root common.Hash, db ethdb.Database) (*trie.Trie, error) {

	return trie.NewTrieWithPrefix(root, protocol.ValidatorPrefix, db)
}

func NewVoteTrie(root common.Hash, db ethdb.Database) (*trie.Trie, error) {

	return trie.NewTrieWithPrefix(root, protocol.VotePrefix, db)
}

func NewDposContext(db ethdb.Database) (*DposContext, error) {

	log.Info("dpos_context.go NewDposContext")

	epochTrie, err := NewEpochTrie(common.Hash{}, db)
	if err != nil {
		return nil, err
	}
	log.Info("NewEpochTrie")

	validatorTrie, err := NewValidatorTrie(common.Hash{}, db)
	if err != nil {
		return nil, err
	}
	log.Info("NewValidatorTrie")

	voteTrie, err := NewVoteTrie(common.Hash{}, db)
	if err != nil {
		return nil, err
	}
	log.Info("NewVoteTrie")

	return &DposContext{
		epochTrie:     epochTrie,
		validatorTrie: validatorTrie,
		voteTrie:      voteTrie,
		db:            db,
	}, nil
}

func NewDposContextFromProto(db ethdb.Database, ctxProto *DposContextProto) (*DposContext, error) {

	epochTrie, err := NewEpochTrie(ctxProto.EpochHash, db)
	if err != nil {
		log.Error("NewEpochTrie", "Error", err)
		return nil, err
	}

	validatorTrie, err := NewValidatorTrie(ctxProto.ValidatorHash, db)
	if err != nil {
		log.Error("NewValidatorTrie", "Error", err)
		return nil, err
	}

	voteTrie, err := NewVoteTrie(ctxProto.VoteHash, db)
	if err != nil {
		log.Error("NewBlockCntTrie", "Error", err)
		return nil, err
	}

	dposContext := new(DposContext)
	dposContext.epochTrie = epochTrie
	dposContext.validatorTrie = validatorTrie
	dposContext.voteTrie = voteTrie
	dposContext.db = db

	return dposContext, nil
}

func (d *DposContext) Copy() *DposContext {

	epochTrie := *d.epochTrie
	validatorTrie := *d.validatorTrie
	voteTrie := *d.voteTrie

	return &DposContext{
		epochTrie:     &epochTrie,
		validatorTrie: &validatorTrie,
		voteTrie:      &voteTrie,
	}
}

func (d *DposContext) Root() (h common.Hash) {

	hw := sha3.NewKeccak256()
	rlp.Encode(hw, d.epochTrie.Hash())
	rlp.Encode(hw, d.validatorTrie.Hash())
	rlp.Encode(hw, d.voteTrie.Hash())
	hw.Sum(h[:0])
	return h
}

func (d *DposContext) Snapshot() *DposContext {
	return d.Copy()
}

func (d *DposContext) RevertToSnapShot(snapshot *DposContext) {

	d.epochTrie = snapshot.epochTrie
	d.validatorTrie = snapshot.validatorTrie
	d.voteTrie = snapshot.voteTrie
}

func (d *DposContext) FromProto(dcp *DposContextProto) error {

	var err error
	d.epochTrie, err = NewEpochTrie(dcp.EpochHash, d.db)
	if err != nil {
		return err
	}

	d.validatorTrie, err = NewValidatorTrie(dcp.ValidatorHash, d.db)
	if err != nil {
		return err
	}

	d.voteTrie, err = NewVoteTrie(dcp.VoteHash, d.db)
	return err
}

type DposContextProto struct {
	EpochHash     common.Hash `json:"epochRoot"        gencodec:"required"`
	ValidatorHash common.Hash `json:"validatorRoot"     gencodec:"required"`
	VoteHash      common.Hash `json:"voteRoot"      gencodec:"required"`
}

func (d *DposContext) ToProto() *DposContextProto {
	return &DposContextProto{
		EpochHash:     d.epochTrie.Hash(),
		ValidatorHash: d.validatorTrie.Hash(),
		VoteHash:      d.voteTrie.Hash(),
	}
}

func (p *DposContextProto) Root() (h common.Hash) {

	hw := sha3.NewKeccak256()
	rlp.Encode(hw, p.EpochHash)
	rlp.Encode(hw, p.ValidatorHash)
	rlp.Encode(hw, p.VoteHash)
	hw.Sum(h[:0])
	return h
}

func (d *DposContext) CommitTo(dbw trie.DatabaseWriter) (*DposContextProto, error) {

	epochRoot, err := d.epochTrie.CommitTo(dbw)
	if err != nil {
		return nil, err
	}

	validatorRoot, err := d.validatorTrie.CommitTo(dbw)
	if err != nil {
		return nil, err
	}

	voteRoot, err := d.voteTrie.CommitTo(dbw)
	if err != nil {
		return nil, err
	}

	return &DposContextProto{
		EpochHash:     epochRoot,
		ValidatorHash: validatorRoot,
		VoteHash:      voteRoot,
	}, nil
}

func (d *DposContext) EpochTrie() *trie.Trie             { return d.epochTrie }
func (d *DposContext) ValidatorTrie() *trie.Trie         { return d.validatorTrie }
func (d *DposContext) VoteTrie() *trie.Trie              { return d.voteTrie }
func (d *DposContext) SetEpoch(epoch *trie.Trie)         { d.epochTrie = epoch }
func (d *DposContext) SetValidator(validator *trie.Trie) { d.validatorTrie = validator }
func (d *DposContext) SetVote(voteTrie *trie.Trie)       { d.voteTrie = voteTrie }
func (d *DposContext) DB() ethdb.Database                { return d.db }

func (dc *DposContext) GetEpochTrie() ([]common.Address, error) {

	//log.Info("(dc *DposContext) GetEpochTrie", "epochTrie", dc.epochTrie.Hash().String())

	var validators []common.Address
	validatorsRLP := dc.epochTrie.Get(protocol.ValidatorsPrefix)

	if err := rlp.DecodeBytes(validatorsRLP, &validators); err != nil {

		return nil, protocol.ErrDecodeValidators
	}
	return validators, nil
}

func (dc *DposContext) SetEpochTrie(validators []common.Address) error {

	log.Info("(dc *DposContext) SetEpochTrie", "size", len(validators))

	validatorsRLP, err := rlp.EncodeToBytes(validators)
	if err != nil {
		return protocol.ErrEncodeValidators
	}
	dc.epochTrie.Update(protocol.ValidatorsPrefix, validatorsRLP)
	return nil
}

func (dc *DposContext) Clean() error {

	log.Info("(dc *DposContext) Clean")

	//得到当前的验证者
	var validators []common.Address
	validatorsRLP := dc.epochTrie.Get(protocol.ValidatorsPrefix)
	if err := rlp.DecodeBytes(validatorsRLP, &validators); err != nil {
		return protocol.ErrDecodeValidators
	}

	//删除当前验证人
	for _, v := range validators {

		err := dc.validatorTrie.TryDelete(v.Bytes())
		if err != nil {
			if _, ok := err.(*trie.MissingNodeError); !ok {
				return err
			}
		}
	}

	//清空周期树
	validators = make([]common.Address, 0)
	validatorsRLP, err := rlp.EncodeToBytes(validators)
	if err != nil {
		return protocol.ErrEncodeValidators
	}
	dc.epochTrie.Update(protocol.ValidatorsPrefix, validatorsRLP)
	return nil
}

func (dc *DposContext) InsertValidator(validator common.Address, votes *big.Int) error {

	log.Info("(dc *DposContext) InsertValidator", "validator", validator, "votes", votes.String())

	//得到当前所有的验证者
	var validators []common.Address
	validatorsRLP := dc.epochTrie.Get(protocol.ValidatorsPrefix)
	if err := rlp.DecodeBytes(validatorsRLP, &validators); err != nil {
		return protocol.ErrDecodeValidators
	}
	if len(validators) >= protocol.MaxValidatorSize {
		return protocol.ErrValidatorsIsFull
	}

	key := validator.Bytes()
	validatorRLP, err := rlp.EncodeToBytes(votes.String())
	if err != nil {
		return fmt.Errorf("failed to encode validators to rlp bytes: %s", err)
	}
	if err := dc.validatorTrie.TryUpdate(key, validatorRLP); err != nil {
		return err
	}

	//加入
	validators = append(validators, validator)
	validatorsRLP, encodeErr := rlp.EncodeToBytes(validators)
	if err != nil {
		return encodeErr
	}
	dc.epochTrie.Update(protocol.ValidatorsPrefix, validatorsRLP)

	return nil
}

func (dc *DposContext) GetValidatorCnt(validator common.Address) (*big.Int, error) {

	//根据地址获取数据
	key := validator.Bytes()
	v, err := dc.epochTrie.TryGet(key)
	if err != nil {
		return big.NewInt(0), err
	}

	//转换成交易类型
	cnt, err := strconv.Atoi(string(v[:]))
	return big.NewInt(int64(cnt)), nil
}

func (dc *DposContext) SetValidatorVotes(validators []common.Address, votes []*big.Int) error {

	//清空验证人
	dc.Clean()

	//重建验证人
	for index, validator := range validators {
		cnt := votes[index].Int64()
		if err := dc.validatorTrie.TryUpdate(validator.Bytes(), []byte(strconv.Itoa(int(cnt)))); err != nil {
			return fmt.Errorf("failed to TryUpdate validator: %s", err)
		}
	}
	return dc.SetEpochTrie(validators)
}

func (dc *DposContext) IsValidator(address common.Address) bool {

	validators, err := dc.GetEpochTrie()
	if err != nil {
		return false
	}
	for _, v := range validators {
		if address == v {
			return true
		}
	}
	return false
}

func (dc *DposContext) IsValidatorFull() bool {

	validators, err := dc.GetEpochTrie()
	if err != nil {
		return true
	}
	if len(validators) >= protocol.MaxValidatorSize {
		return true
	}
	return false
}

func (dc *DposContext) GetCurrentProducer(firstTimer int64) (common.Address, error) {

	producers, err := dc.GetEpochTrie()
	if err != nil {
		return common.Address{}, errors.New("failed to GetValidators")
	}
	producerSize := len(producers)
	if producerSize == 0 {
		return common.Address{}, protocol.ErrEpochTrieNil
	}

	offset := (time.Now().Unix() - firstTimer) % protocol.EpochInterval
	offset /= protocol.BlockInterval

	offset %= int64(producerSize)
	return producers[offset], nil
}

func (dc *DposContext) GetCurrentNowProducer(firstTimer int64, now int64) (common.Address, error) {

	producers, err := dc.GetEpochTrie()
	if err != nil {
		return common.Address{}, errors.New("failed to GetValidators")
	}
	producerSize := len(producers)
	if producerSize == 0 {
		return common.Address{}, protocol.ErrEpochTrieNil
	}

	offset := (now - firstTimer) % protocol.EpochInterval
	offset /= protocol.BlockInterval

	offset %= int64(producerSize)
	return producers[offset], nil
}

func (dc *DposContext) GetLastProducer(indexOffset int, firstTimer int64) (common.Address, error) {

	producers, err := dc.GetEpochTrie()
	if err != nil {
		return common.Address{}, errors.New("failed to GetValidators")
	}
	producerSize := len(producers)
	if producerSize == 0 {
		return common.Address{}, protocol.ErrEpochTrieNil
	}

	offset := (time.Now().Unix() - firstTimer) % protocol.EpochInterval
	offset /= protocol.BlockInterval

	offset %= int64(producerSize)

	if offset+int64(indexOffset) < 0 {
		return producers[len(producers)-1], nil
	} else {
		return producers[int(offset)+indexOffset], nil
	}
}

//根据时间得到当时的打包节点
func (dc *DposContext) GetProducer(now int64, firstTimer int64) (producer common.Address, err error) {

	producer = common.Address{}
	offset := (now - firstTimer) % protocol.EpochInterval
	if offset%protocol.BlockInterval != 0 {

		log.Info("GetProducer", "offset", offset%protocol.BlockInterval)
		return common.Address{}, protocol.ErrInvalidProducerTime
	}
	offset /= protocol.BlockInterval

	//得到验证者数组
	producers, err := dc.GetEpochTrie()
	if err != nil {
		return common.Address{}, err
	}

	//得到验证者数量并判断验证者数量是否为0
	producerSize := len(producers)
	if producerSize == 0 {
		return common.Address{}, protocol.ErrInvalidProducer
	}

	//根据当前的移位偏移量确定当前应该出块的验证者
	offset %= int64(producerSize)
	return producers[offset], nil
}

func (dc *DposContext) CreateEpoch() error {

	return nil
}

type sortableAddress struct {
	address common.Address
	weight  *big.Int
}
type sortableAddresses []*sortableAddress

func (p sortableAddresses) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p sortableAddresses) Len() int      { return len(p) }
func (p sortableAddresses) Less(i, j int) bool {
	if p[i].weight.Cmp(p[j].weight) < 0 {
		return false
	} else if p[i].weight.Cmp(p[j].weight) > 0 {
		return true
	} else {
		return p[i].address.String() < p[j].address.String()
	}
}
