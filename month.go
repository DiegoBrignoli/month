package month  //tutti e 4

type Month int

const (
	Gennaio Month = 1 + iota
	Febbraio
	Marzo
	Aprile
	Maggio
	Giugno
	Luglio
	Agosto
	Settembre
	Ottobre
	Novembre
	Dicembre
)

var longMonthNames = []string{
	"Gennaio",
	"Febbraio",
	"Marzo",
	"Aprile",
	"Maggio",
	"Giugno",
	"Luglio",
	"Agosto",
	"Settembre",
	"Ottobre",
	"Novembre",
	"Dicembre",
}

// String returns the Italian name of the month ("Gennaio", "Febbraio", ...).
func (m Month) String() string {
	if Gennaio <= m && m <= Dicembre {
		return longMonthNames[m-1]
	}
	buf := make([]byte, 20)
	n := fmtInt(buf, uint64(m))
	return "%!Month(" + string(buf[n:]) + ")"
}

// fmtInt formats v into the tail of buf.
// It returns the index where the output begins.
func fmtInt(buf []byte, v uint64) int {
	w := len(buf)
	if v == 0 {
		w--
		buf[w] = '0'
	} else {
		for v > 0 {
			w--
			buf[w] = byte(v%10) + '0'
			v /= 10
		}
	}
	return w
}

