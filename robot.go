package toy_robot

import (
	"fmt"
)


const (
	TABLE_DIM_X	= 5
	TABLE_DIM_Y = 5
)


// direction type related defs
var directionEnums []string;

type DirectionType uint8

func (d DirectionType) String() string {
	return directionEnums[int(d)]
}

func ciota(s string) DirectionType {
	directionEnums = append(directionEnums, s)
	return DirectionType(len(directionEnums) - 1)
}

var (
	NORTH = ciota("NORTH") 
	EAST = ciota("EAST")
	SOUTH = ciota("SOUTH")
	WEST = ciota("WEST")

	NODIR = ciota("NODIR")
)

func stringToDirectionType(dir string) DirectionType {
	if dir == directionEnums[NORTH] {
		return NORTH
	} else if dir == directionEnums[EAST] {
		return EAST
	} else if dir == directionEnums[SOUTH] {
		return SOUTH
	} else if dir == directionEnums[WEST] {
		return WEST
	} else {
		return NODIR
	} 
}

// easy mechanism for moving robot using
// pre calculated increments on axis based on direction
type MovePair struct {
	x, y interface{}
}

var movePairArray [4]MovePair


type Robot struct {
	X			int
	Y			int
	Direction 	DirectionType
}

func init() {
	movePairArray[NORTH] = MovePair{0, +1}
	movePairArray[EAST] = MovePair{+1, 0}
	movePairArray[SOUTH] = MovePair{0, -1}
	movePairArray[WEST] = MovePair{-1, 0}
}

// Craete and initialize Robot 
func newRobot() Robot {
	r := Robot{-1, -1, NODIR}
	return r
}

// Validate given x,y to be on the table
func (r *Robot) validateXY(x int, y int) bool {
	return x > -1 && x < TABLE_DIM_X && y > -1 && y < TABLE_DIM_Y;
}

// PLACE will put the toy Robot on the table in position X,Y 
// and facing NORTH, SOUTH, EAST or WEST. 
func (r *Robot) Place(x int, y int, dir string) {
	if direction := stringToDirectionType(dir); r.validateXY(x, y) && direction != NODIR {
		r.X = x
		r.Y = y
		r.Direction = direction
	}
}

// MOVE will move the toy Robot one unit forward in the 
// direction it is currently facing.
func (r *Robot) Move() {
	if r.Direction == NODIR {
		return
	}

	nextX := r.X + movePairArray[r.Direction].x.(int)
	nextY := r.Y + movePairArray[r.Direction].y.(int)

	if r.validateXY(nextX, nextY) {
		r.X = nextX;
		r.Y = nextY;
	}
}

// LEFT will rotate the Robot 90 degrees in the specified 
// direction without changing the position of the Robot.
func (r *Robot) Left() {
	r.Direction--
	r.Direction %= 4;
}

// RIGHT will rotate the Robot 90 degrees in the specified 
// direction without changing the position of the Robot.
func (r *Robot) Right() {
	r.Direction++
	r.Direction %= 4;
}

// REPORT will announce the X,Y and orientation of the Robot.
func (r *Robot) Report() {
	if r.validateXY(r.X, r.Y) {
		fmt.Printf("%d,%d,%s\n", r.X, r.Y, r.Direction.String())
	}
}