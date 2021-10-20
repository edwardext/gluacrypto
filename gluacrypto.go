package gluacrypto

import (
	crypto "github.com/edwardext/gluacrypto/crypto"
	lua "github.com/yuin/gopher-lua"
)

func Preload(L *lua.LState) {
	L.PreloadModule("crypto", crypto.Loader)
}
