package models

import "strconv"

// 坐标点
type Point struct {
	// 唯一值 （x,y）
	Key string
	// 横坐标
	X int
	// 纵坐标
	Y int
	// 展示符
	View string
}

// 设置点唯一值
func pointAsKey(x, y int) (key string) {
	key = strconv.Itoa(x) + "," + strconv.Itoa(y)
	return key
}

