package funcapi

//
// checks whether a codeset contains the supplied value
//
func CodesetContains(codeset map[string]struct{}, value interface{}) bool {

	vstr, ok := value.(string)
	if !ok {
		return ok
	}
	_, ok = codeset[vstr]
	return ok
}

var empty = struct{}{} // helper to save typing, also requires zero space in map
// codeset as lookup map, faster than iterating array
var AUCodeSetsYearLevelCodeTypeValues = map[string]struct{}{
	"0":        empty,
	"1":        empty,
	"2":        empty,
	"3":        empty,
	"4":        empty,
	"5":        empty,
	"6":        empty,
	"7":        empty,
	"8":        empty,
	"9":        empty,
	"10":       empty,
	"11":       empty,
	"12":       empty,
	"13":       empty,
	"K":        empty,
	"P":        empty,
	"K3":       empty,
	"K4":       empty,
	"CC":       empty,
	"PS":       empty,
	"UG":       empty,
	"11MINUS":  empty,
	"12PLUS":   empty,
	"UGJunSec": empty,
	"UGPri":    empty,
	"UGSec":    empty,
	"UGSnrSec": empty,
}
