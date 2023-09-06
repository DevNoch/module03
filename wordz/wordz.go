package wordz

import (
	"crypto/rand"
	"math/big"
)

var Hello = "Hello it`s package wordz"
var Prefix = "Random wordz"
var Wordz = []string{"One", "Two", "Three", "Four"}

func Random() string {
	max := len(Wordz)
	r, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return get(r.Int64())
}

func get(index int64) string {
	return Prefix + Wordz[index]
}

func init(){
fmt.Println("Func init in package wordz")
}

func init() {
	fmt.Println("Function init in package wordz")
}
