# A-Star(A*)算法

#### 前言

##### 简介

A-Star(A*)算法作为Dijkstra算法的扩展，在寻路和图的遍历过程中具有一定的高效性。

##### 应用场景

静态图搜索
#### 算法实现

##### 核心价值（代价）计算

*f*(*i*)=*g*(*i*)+*h*(*i*)；

当前节点的价值估值 *f*(*i*) =起始点到该节点的距离 *g*(*i*) +当前节点距离终点的距离 *h*(*i*) (启发函数)

启发函数一般选用曼哈顿距离。

#### 概念

| 结构       | 名          | 描述                                       |
| ---------- | ----------- | ------------------------------------------ |
| Graph      | 结构图      | 整体图，主图                               |
| SearchRoad | 规划路径    | 规划路径                                   |
| Point      | 坐标点      | 主图中的坐标点                             |
| AstarPoint | AStar算法点 | 用于AStar计算的坐标点（继承：组合了Point） |

##### 常量

+ 横向移动一次代价为10

+ 纵向移动一次代价为10

+ 斜向移动一次代价约为14

```go
	X_LEN=10
	Y_LEN=10
	X_Y_LEN=14 // 14.142135623731
```

##### 价值计算代码

```Go
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
	fmt.Printf("%d,%d====> %d %d %d \n",this.X,this.Y,this.g_Value,this.h_Value,this.f_Value)
	return this.f_Value
}
```






