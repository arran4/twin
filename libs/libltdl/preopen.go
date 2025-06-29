package libltdl

// PreloadSym represents a preloaded symbol.
type PreloadSym struct {
	Name    string
	Address interface{}
}

var (
	defaultPreloaded []PreloadSym
	preloadedLists   [][]PreloadSym
)

// DlPreloadDefault records preloaded symbols to be used by default.
func DlPreloadDefault(preloaded []PreloadSym) int {
	defaultPreloaded = preloaded
	return 0
}

// DlPreload adds preloaded symbols. Passing nil resets to the defaults.
func DlPreload(preloaded []PreloadSym) int {
	if preloaded == nil {
		preloadedLists = nil
		if defaultPreloaded != nil {
			preloadedLists = append(preloadedLists, defaultPreloaded)
		}
		return 0
	}
	preloadedLists = append(preloadedLists, preloaded)
	return 0
}

// DlPreloadOpen invokes cb for each module from originator.
func DlPreloadOpen(originator string, cb func(*DlHandle) int) int {
	var errors int
	for _, list := range preloadedLists {
		if len(list) == 0 {
			continue
		}
		if originator != "" && list[0].Name != originator {
			continue
		}
		for i := 1; i < len(list); i++ {
			if list[i].Address == nil && list[i].Name != "" {
				h := DlOpen(list[i].Name)
				if h == nil {
					errors++
				} else if cb != nil {
					errors += cb(h)
				}
			}
		}
	}
	if errors > 0 {
		return errors
	}
	return 0
}
