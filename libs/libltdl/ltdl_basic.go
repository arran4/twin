package libltdl

var initialized bool

// DlInit initializes the loader system.
func DlInit() int {
	if initialized {
		return 0
	}
	initialized = true
	// register the built-in dlopen loader
	DlloaderAdd(DlopenVtable)
	return 0
}

// DlExit shuts down the loader system.
func DlExit() int {
	if !initialized {
		LtSetLastError("library already shutdown")
		return 1
	}
	initialized = false
	loaders = nil
	userSearchPath = nil
	return 0
}
