package core

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tetrex/custom-blockchain/types"
)

func TestHeader_Encode_Decode(t *testing.T) {

	h := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		TimeStamp: time.Now().UnixNano(),
		Height:    10,
		Nounce:    9568410,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(buf))

	h_decode := &Header{}
	assert.Nil(t, h_decode.DecodeBinary(buf))

	assert.Equal(t, h, h_decode)
}

func TestBlock_Encode_Decode(t *testing.T) {

	h := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			TimeStamp: time.Now().UnixNano(),
			Height:    10,
			Nounce:    9568410,
		},
		Transactions: nil,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(buf))

	b_decode := &Block{}
	assert.Nil(t, b_decode.DecodeBinary(buf))

	assert.Equal(t, h, b_decode)
}

func TestBlockHash(t *testing.T) {

	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			TimeStamp: time.Now().UnixNano(),
			Height:    10,
			Nounce:    9568410,
		},
		Transactions: nil,
	}
	h := b.Hash()
	assert.False(t, h.IsZero())
}
