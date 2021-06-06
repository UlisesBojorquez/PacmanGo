package pacman

const (
	backgroundImageSize = 100
	stageBlocSize       = 32 //size if the image 32x32
)

/*USED FOR THE IMAGES*/
type elem int

const (
	w0 elem = iota //value is 0 and increse for the other elem
	w1
	w2
	w3
	w4
	w5
	w6
	w7
	w8
	w9         //ending of the walls
	w10        // a
	w11        // b
	w12        // c
	w13        // d
	w14        // e
	w15        // f
	w16        // g
	w17        // h
	w18        // i
	w19        // j
	w20        // k
	w21        // l
	w22        // m
	w23        // n
	w24        // o
	playerElem // p
	bigDotElem // q
	dotElem    // r
	empty      // s
	blinkyElem // t
	clydeElem  // u
	inkyElem   // v
	pinkyElem  // w
	backgroundElem
)

/*USED FOT THE INPUT*/
type input int

const (
	_ input = iota
	up
	right
	down
	left
	rKey
)
