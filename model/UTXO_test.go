package model_test

import (
	"blockchain_seeder/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUTXO(t *testing.T) {
	script := []byte{0x00, 0x01, 0x02, 0x03, 0x04}

	utxo := model.NewUTXO(script, 1000)

	assert.Equal(t, script, utxo.Script)
	assert.Equal(t, 5, len(utxo.Script))
	assert.Equal(t, uint64(1000), utxo.Value)
}

func TestEqual(t *testing.T) {
	script1 := []byte{0x00, 0x01, 0x02, 0x03, 0x04}
	script2 := []byte{0x00, 0x01, 0x02, 0x03, 0x04}

	utxo1 := model.NewUTXO(script1, 1000)
	utxo2 := model.NewUTXO(script2, 1000)

	assert.Equal(t, utxo1, utxo2)
	assert.True(t, utxo1.Equal(utxo2))
}
