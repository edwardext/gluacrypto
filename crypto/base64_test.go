package gluacrypto_crypto_test

import (
	"encoding/base64"
	"testing"

	"github.com/edwardext/gluacrypto"
	"github.com/edwardext/gluacrypto/luautil"
	"github.com/stretchr/testify/assert"
	lua "github.com/yuin/gopher-lua"
)

func TestBase64Encode(t *testing.T) {
	assert := assert.New(t)

	// test start
	L := lua.NewState()
	defer L.Close()
	gluacrypto.Preload(L)

	script := `
		crypto = require('crypto')
		return crypto.base64_encode('` + string(Data) + `')
	`
	assert.NoError(L.DoString(script))

	val := luautil.GetValue(L, 1)
	err := luautil.GetValue(L, 2)
	assert.Nil(err)
	assert.Equal(base64.StdEncoding.EncodeToString(Data), val)
}

func TestBase64Decode(t *testing.T) {
	assert := assert.New(t)

	// test start
	L := lua.NewState()
	defer L.Close()
	gluacrypto.Preload(L)

	b64data := base64.StdEncoding.EncodeToString(Data)

	script := `
		crypto = require('crypto')
		return crypto.base64_decode('` + b64data + `')
	`
	assert.NoError(L.DoString(script))

	val := luautil.GetValue(L, 1)
	err := luautil.GetValue(L, 2)
	assert.Nil(err)
	assert.Equal(string(Data), val)
}
