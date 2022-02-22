package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type cint int

var SidesTriangle cint = 3
var SidesSquare cint = 4
var SidesCircle cint = 0

// var SidesDummy cint = 77
// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum cint) float64 {
	if sidesNum == 4 {
		return math.Pow(sideLen, 2.0)
	} else if sidesNum == 3 {
		return math.Sqrt(3.0) / 4 * math.Pow(sideLen, 2.0)
	} else if sidesNum == 0 {
		return math.Pi * math.Pow(sideLen, 2.0)
	} else {
		return 0
	}
}
