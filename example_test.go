package iota_test

import (
	"fmt"
	"testing"

	"github.com/iotaledger/iota.go"
	"github.com/stretchr/testify/require"
)

func TestSignedTransactionPayloadSize(t *testing.T) {
	sigTxPayload := oneInputOutputSignedTransactionPayload()
	m := &iota.Message{
		Version: 1,
		Parent1: randTxHash(),
		Parent2: randTxHash(),
		Payload: sigTxPayload,
		Nonce:   0,
	}

	data, err := m.Serialize(iota.DeSeriModeNoValidation)
	require.NoError(t, err)
	fmt.Printf("length of message cotaining a signed transaction payload: %d\n", len(data))
}