package models

const (
	// SKIP skip set attr. used by SetStringTiles
	SKIP =" ❄ " //byte('*')
	// PLAIN  point can be arrived to
	PLAIN = " ☐ " //byte('☐')
	// BLOCK  point can not be arrived to
	BLOCK =" ■ " // //byte('x')
	// START  the start point
	START =" ☆ " //byte('s')
	// END  the end point
	END =" ★ " //byte('e')
	// PATH  not contains start and end.
	PATH =" ◉ " //byte('o')

	X_LEN=10
	Y_LEN=10
	X_Y_LEN=14//.142135623731
)




