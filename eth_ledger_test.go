// +build eth_ledger_app

/*
 * Ledger devices follow the ISO/IEC 7816-4 APDU protocol.
 * Link to full details here: https://en.wikipedia.org/wiki/Smart_card_application_protocol_data_unit .
 * The specs for ledger-app-eth can be found here: https://github.com/LedgerHQ/ledger-app-eth/blob/master/doc/ethapp.asc .
 */

package ledger_go

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/assert"
)

const (
	ETH_CLA        byte = 0xE0
	ETH_INS_CONFIG byte = 0x06
	ETH_INS_DERIVE byte = 0x02
	ETH_INS_SIGN   byte = 0x04
)

func Test_AppConfiguration(t *testing.T) {
	ledger, err := FindLedger()
	if err != nil {
		fmt.Println("\n*********************************")
		fmt.Println("Did you enter the password??")
		fmt.Println("*********************************")
		t.Fatalf("Error: %s", err.Error())
	}
	assert.NotNil(t, ledger)

	message := []byte{ETH_CLA, ETH_INS_CONFIG, 0, 0, 0}

	for i := 0; i < 10; i++ {
		response, err := ledger.Exchange(message)
		assert.Nil(t, err, "App config instruction failed at iteration %d\n", i)

		assert.Equal(t, 4, len(response))
	}
	ledger.Close()
}

func Test_Derive(t *testing.T) {
	ledger, err := FindLedger()
	assert.Nil(t, err)

	derivationPath := accounts.DefaultRootDerivationPath
	path := make([]byte, 1+4*len(derivationPath))
	path[0] = byte(len(derivationPath))
	for i, component := range derivationPath {
		binary.BigEndian.PutUint32(path[1+4*i:], component)
	}

	header := []byte{ETH_CLA, ETH_INS_DERIVE, 0, 0, byte(len(path))}
	message := append(header, path...)

	response, err := ledger.Exchange(message)
	assert.Nil(t, err)
	pkLength := int(response[0])
	publicKey := hex.EncodeToString(response[1 : pkLength+1])
	addrLength := int(response[pkLength+1])
	address := hex.EncodeToString(response[pkLength+2 : pkLength+2+addrLength])
	fmt.Printf("Public Key Length: %d\nPublic Key: %s\nAddress Length: %d\nAddress: %s\n", pkLength, publicKey, addrLength, address)
	assert.Equal(t, len(response), pkLength+addrLength+2)

	ledger.Close()
}

// The sign transaction instruction can only be used on transactions with empty data fields.
// Transaction must be confirmed on device to sign, and for the test to proceed.
func Test_EthereumSignTransaction(t *testing.T) {
	ledger, err := FindLedger()
	assert.Nil(t, err)

	derivationPath := accounts.DefaultRootDerivationPath
	path := make([]byte, 1+4*len(derivationPath))
	path[0] = byte(len(derivationPath))
	for i, component := range derivationPath {
		binary.BigEndian.PutUint32(path[1+4*i:], component)
	}

	tx := types.NewTransaction(
		3,
		common.HexToAddress("b94f5374fce5edbc8e2a8697c15331677e6ebf0b"),
		big.NewInt(10),
		2000,
		big.NewInt(1),
		[]byte{},
	)
	tx_serialized, err := rlp.EncodeToBytes([]interface{}{tx.Nonce(), tx.GasPrice(), tx.Gas(), tx.To(), tx.Value(), tx.Data()})
	assert.Nil(t, err)

	payload := append(path, tx_serialized...)

	header := []byte{ETH_CLA, ETH_INS_SIGN, 0, 0, byte(len(payload))}

	message := append(header, payload...)
	response, err := ledger.Exchange(message)
	assert.Nil(t, err)

	fmt.Printf("Signed transaction: %s\n", hex.EncodeToString(response))
	ledger.Close()
}
