package Utils

const (
	backgroundImageSize = 100
	//screenWidth         = backgroundImageSize * 7
	//screenHeight        = backgroundImageSize * 7 //backgroundElem
	stageBlocSize = 32 //size if the image 32x32
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
	w9
	w10 // a
	w11
	w12
	w13
	w14
	w15
	w16
	w17
	w18
	w19
	w20
	w21
	w22
	w23
	w24
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
