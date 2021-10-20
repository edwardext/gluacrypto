package gluacrypto_crypto_test

import (
	"testing"

	"github.com/edwardext/gluacrypto"
	"github.com/stretchr/testify/assert"
	lua "github.com/yuin/gopher-lua"
)

// consts
var (
	Data   []byte
	Key    []byte
	Key128 []byte
	IV     []byte
	IV128  []byte
)

func init() {
	Data = []byte("abcd")
	Key = []byte("key12345")
	Key128 = []byte("key1234567890123")
	IV = []byte("12345678")
	IV128 = []byte("1234567890123456")
}

func TestCrypto(t *testing.T) {
	assert := assert.New(t)

	// test start
	L := lua.NewState()
	defer L.Close()
	gluacrypto.Preload(L)

	script := `
		return require('crypto')
	`
	assert.NoError(L.DoString(script))

	c := L.Get(1)
	assert.NotNil(c)
}
