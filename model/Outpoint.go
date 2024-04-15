package model

import (
	"encoding/binary"
	"fmt"

	"github.com/libsv/go-bt/v2/chainhash"
)

// Outpoint represents a bitcoin transaction output.
type Outpoint struct {
	TxID  chainhash.Hash
	Index uint32
}

// NewOutpoint creates a new Outpoint.
func NewOutpoint(txID chainhash.Hash, index uint32) *Outpoint {
	return &Outpoint{
		TxID:  txID,
		Index: index,
	}
}

// NewOutpointFromBytes creates a new Outpoint from a byte slice. It expects a byte slice of exactly 36 bytes,
// where the first 32 bytes are the transaction ID (little endian) and the last 4 bytes are the index (little endian).
func NewOutpointFromBytes(b []byte) (*Outpoint, error) {
	if len(b) != 36 {
		return nil, fmt.Errorf("invalid outpoint length: expected 36 bytes, got %d", len(b))
	}

	txID, err := chainhash.NewHash(b[:32])
	if err != nil {
		return nil, fmt.Errorf("failed to create hash from bytes: %v", err)
	}

	index := binary.LittleEndian.Uint32(b[32:])

	return &Outpoint{
		TxID:  *txID,
		Index: index,
	}, nil
}

// Bytes returns a byte slice representation of the Outpoint. The first 32 bytes are
// the transaction ID (little endian) and the last 4 bytes are the index (little endian).
func (o *Outpoint) Bytes() []byte {
	// Write the txid and a varint of the index to a byte slice
	serialized := make([]byte, 36)
	copy(serialized, o.TxID[:])
	binary.LittleEndian.PutUint32(serialized[32:], o.Index)

	return serialized
}

// String returns a string representation of the Outpoint, formatted as "txid:index". In this case,
// the txid is the big-endian representation of the transaction ID in hex format (64 characters).
func (o *Outpoint) String() string {
	return fmt.Sprintf("%v:%d", o.TxID, o.Index)
}

func (o *Outpoint) Equal(other *Outpoint) bool {
	return o.TxID.IsEqual(&other.TxID) && o.Index == other.Index
}
