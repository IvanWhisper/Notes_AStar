package models

import (
	"math"
)

// AStar算法点
type AstarPoint struct {
	// 继承
	Point
	// 上一节点
	PrevPoint *AstarPoint
	// 距离值
	g_Value int
	// 启发值
	h_Value int
	// 总价值（越小越好）
	f_Value int
}

func NewAstarPoint(p *Point, prev *AstarPoint, end *AstarPoint) (ap *AstarPoint) {
	ap = &AstarPoint{*p, prev, 0, 0, 0}
	if end != nil {
		ap.Cal_F_Value(end)
	}
	return ap
}

// 计算 距离
func (this *AstarPoint) Cal_G_Value() int {
	if this.PrevPoint != nil {
		// 当前点与上一点的横向距离
		deltaX := math.Abs(float64(this.PrevPoint.X - this.X))
		// 当前点与上一点的纵向距离
		deltaY := math.Abs(float64(this.PrevPoint.Y - this.Y))

		if deltaX == 1 && deltaY == 0 {
			// 横向 增量 1
			this.g_Value = this.PrevPoint.g_Value + X_LEN
		} else if deltaX == 0 && deltaY == 1 {
			// 纵向 增量 1
			this.g_Value = this.PrevPoint.g_Value + Y_LEN
		} else if deltaX == 1 && deltaY == 1 {
			// 斜向 增量 1
			this.g_Value = this.PrevPoint.g_Value + X_Y_LEN
		} else {
			panic("PrevPoint is invalid!")
		}
	}
	return this.g_Value
}

// 计算 曼哈顿距离
func (this *AstarPoint) Cal_H_Value(end *AstarPoint) int {
	this.h_Value = int(math.Abs(float64(end.X-this.X))*X_LEN + math.Abs(float64(end.Y-this.Y)))*Y_LEN
	return this.h_Value
}

// 计算 总价值
func (this *AstarPoint) Cal_F_Value(end *AstarPoint) int {
	this.f_Value = this.Cal_G_Value() + this.Cal_H_Value(end)
	// fmt.Printf("%d,%d====> %d %d %d \n",this.X,this.Y,this.g_Value,this.h_Value,this.f_Value)
	return this.f_Value
}