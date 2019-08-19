package bitset

type BitSet struct {
	bit uint64
}

func NewBitSet() *BitSet {
	return &BitSet{}
}

func (bs *BitSet) Set(mask uint64) {
	bs.bit |= mask
}

func (bs *BitSet) Reset(mask uint64) {
	bs.bit &= ^(mask)
}

func (bs BitSet) IsSet(mask uint64) bool {
	return bs.bit&(mask) == mask
}
