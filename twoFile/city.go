package twoFile

import (
	"GoTestLiberally/wordz"
	"strings"
)

var Prefix = "Random wordz: "
var Citys = []string{"Ternopil", "Zbarazh", "Detroit", "New York"}

//отримати рузельтат Рандом() . random wortdzTwo .
// відрізати від того результату префікс .Two.
// методом перебору знайти з елементом якого індексу з масиву Вордз співпадає .

func City() string {
	var r int
	var s = wordz.Random()
	stripped := strings.SplitAfter(s, Prefix)[1]

	for i := 0; i < len(Citys); i++ {
		if stripped == wordz.Wordz[i] {
			r = i
		}
	}
	result := Citys[r]
	return result
}
