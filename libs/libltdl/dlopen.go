package libltdl

import (
	"plugin"
)

// DlopenVtable exposes a simple loader based on Go's plugin package.
var DlopenVtable = &Dlvtable{
	Name:        "lt_dlopen",
	ModuleOpen:  dlopenOpen,
	ModuleClose: dlopenClose,
	FindSym:     dlopenSym,
	Priority:    DlloaderPrepend,
}

func dlopenOpen(data interface{}, filename string, advise interface{}) interface{} {
	p, err := plugin.Open(filename)
	if err != nil {
		LtSetLastError(err.Error())
		return nil
	}
	return p
}

func dlopenClose(data interface{}, module interface{}) int {
	// Go plugins cannot be unloaded.
	return 0
}

func dlopenSym(data interface{}, module interface{}, name string) interface{} {
	if mod, ok := module.(*plugin.Plugin); ok {
		sym, err := mod.Lookup(name)
		if err != nil {
			LtSetLastError(err.Error())
			return nil
		}
		return sym
	}
	LtSetLastError("invalid module")
	return nil
}
