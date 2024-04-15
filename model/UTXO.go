package model

import "bytes"

type UTXO struct {
	Script []byte
	Value  uint64
}

func NewUTXO(script []byte, value uint64) *UTXO {
	return &UTXO{
		Script: script,
		Value:  value,
	}
}

func (u *UTXO) Equal(other *UTXO) bool {
	return u.Value == other.Value && bytes.Equal(u.Script, other.Script)
}
