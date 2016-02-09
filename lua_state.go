package lunargo

// #include "lua.h"
// #include "lualib.h"
// #include "lauxlib.h"
// #include <stdlib.h>
import "C"
import "unsafe"

const (
	// LuaMultRet ...
	LuaMultRet = -1
)

// LuaState ...
type LuaState struct {
	lua *C.lua_State
}

// NewLuaState ...
func NewLuaState() *LuaState {
	lua := C.luaL_newstate()
	C.luaL_openlibs(lua)
	return &LuaState{lua}
}

// DoString ...
func (state *LuaState) DoString(script string) (string, error) {
	code := C.CString(script)
	defer C.free(unsafe.Pointer(code))

	C.luaL_loadstring(state.lua, code)
	C.lua_pcallk(state.lua, 0, LuaMultRet, 0, 0, nil)
	return "", nil
}
