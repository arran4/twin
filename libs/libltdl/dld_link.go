package libltdl

// DldLinkVtable mimics the old dld loader but uses Go's plugin loader.
var DldLinkVtable = &Dlvtable{
	Name:        "lt_dld_link",
	ModuleOpen:  dlopenOpen,
	ModuleClose: dlopenClose,
	FindSym:     dlopenSym,
	Priority:    DlloaderAppend,
}
