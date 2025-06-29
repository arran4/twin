package libtw

import "time"

// Timevalue mirrors the C timevalue struct.
type Timevalue struct {
	Seconds  int64
	Fraction int64
}

const (
	THOUSAND int64 = 1000
	NanoSEC  int64 = 1
	MicroSEC       = THOUSAND * NanoSEC
	MilliSEC       = THOUSAND * MicroSEC
	FullSEC        = THOUSAND * MilliSEC
)

func NormalizeTime(t *Timevalue) {
	if t.Fraction >= FullSEC {
		delta := t.Fraction / FullSEC
		t.Seconds += delta
		t.Fraction -= delta * FullSEC
	}
}

func InstantNow(now *Timevalue) *Timevalue {
	n := time.Now()
	now.Seconds = n.Unix()
	now.Fraction = int64(n.Nanosecond())
	return now
}

func IncrTime(t, incr *Timevalue) *Timevalue {
	t.Seconds += incr.Seconds
	t.Fraction += incr.Fraction
	NormalizeTime(t)
	return t
}

func DecrTime(t, decr *Timevalue) *Timevalue {
	t.Seconds -= decr.Seconds
	if t.Fraction >= decr.Fraction {
		t.Fraction -= decr.Fraction
	} else {
		t.Seconds--
		t.Fraction += FullSEC - decr.Fraction
	}
	return t
}

func SubTime(result, t, decr *Timevalue) *Timevalue {
	*result = *t
	return DecrTime(result, decr)
}

func CmpTime(t1, t2 *Timevalue) int {
	if t1.Seconds > t2.Seconds {
		return 1
	} else if t1.Seconds < t2.Seconds {
		return -1
	}
	if t1.Fraction > t2.Fraction {
		return 1
	} else if t1.Fraction < t2.Fraction {
		return -1
	}
	return 0
}
