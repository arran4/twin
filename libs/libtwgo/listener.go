package libtw

import "sync"

type TwEventCommon struct {
	W    uint32
	Code uint32
}

type TwMsg struct {
	Type  uint32
	Event TwEventCommon
}

type TwListener struct {
	TwD      *TwDisplay
	Type     uint32
	Event    TwEventCommon
	Listener func(ev *TwMsg, arg interface{})
	Arg      interface{}
}

func keyForListener(t uint32, ev TwEventCommon) uint64 {
	return uint64(t)<<32 | uint64(ev.W)<<16 | uint64(ev.Code)
}

func (d *TwDisplay) insertListener(L *TwListener) {
	d.init()
	key := keyForListener(L.Type, L.Event)
	d.listeners[key] = append(d.listeners[key], L)
	L.TwD = d
}

func (d *TwDisplay) removeListener(L *TwListener) {
	key := keyForListener(L.Type, L.Event)
	list := d.listeners[key]
	for i, l := range list {
		if l == L {
			list = append(list[:i], list[i+1:]...)
			break
		}
	}
	if len(list) == 0 {
		delete(d.listeners, key)
	} else {
		d.listeners[key] = list
	}
	L.TwD = nil
}

var displayMu sync.Mutex

func Tw_InsertListener(d *TwDisplay, L *TwListener) {
	if d == nil || L == nil {
		return
	}
	displayMu.Lock()
	d.insertListener(L)
	displayMu.Unlock()
}

func Tw_RemoveListener(d *TwDisplay, L *TwListener) {
	if d == nil || L == nil {
		return
	}
	displayMu.Lock()
	if L.TwD == d {
		d.removeListener(L)
	}
	displayMu.Unlock()
}

func Tw_DeleteListener(d *TwDisplay, L *TwListener) {
	if d == nil || L == nil {
		return
	}
	displayMu.Lock()
	if L.TwD == d {
		d.removeListener(L)
	}
	displayMu.Unlock()
}

func Tw_SetDefaultListener(d *TwDisplay, fn func(*TwMsg, interface{}), arg interface{}) {
	if d == nil {
		return
	}
	displayMu.Lock()
	d.DefaultListener = fn
	d.DefaultArg = arg
	displayMu.Unlock()
}

func Tw_CreateListener(d *TwDisplay, typ uint32, ev TwEventCommon, fn func(*TwMsg, interface{}), arg interface{}) *TwListener {
	L := &TwListener{Type: typ, Event: ev, Listener: fn, Arg: arg}
	return L
}

func Tw_DispatchMsg(d *TwDisplay, msg *TwMsg) bool {
	if d == nil || msg == nil {
		return false
	}
	displayMu.Lock()
	key := keyForListener(msg.Type, msg.Event)
	if list := d.listeners[key]; len(list) > 0 {
		for _, l := range list {
			l.Listener(msg, l.Arg)
		}
		displayMu.Unlock()
		return true
	}
	fn := d.DefaultListener
	arg := d.DefaultArg
	displayMu.Unlock()
	if fn != nil {
		fn(msg, arg)
		return true
	}
	return false
}
