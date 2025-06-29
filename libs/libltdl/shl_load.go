package libltdl

// ShlLoadVtable substitutes HP-UX shl_load with Go's plugin API.
var ShlLoadVtable = &Dlvtable{
	Name:        "lt_shl_load",
	ModuleOpen:  dlopenOpen,
	ModuleClose: dlopenClose,
	FindSym:     dlopenSym,
	Priority:    DlloaderAppend,
}
