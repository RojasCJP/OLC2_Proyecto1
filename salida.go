/*----HEADER----*/
package main

import (
	"fmt"
)

var t0, t1 float64
var P, H float64
var stack [30101999]float64
var heap [30101999]float64

func main() {
	if t0 == 10 {
		goto L1
	}
	goto L2
L1:
	fmt.Printf("%c", t0)
L2:
}
