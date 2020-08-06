package models

// 结构图
type Graph struct {
	MaxX,MaxY int
	Start,End *Point
	Points [][]Point
	Blocks map[string]*Point
}

// 创建结构图
func New(x,y int) *Graph {
	g:=&Graph{}
	g.MaxX=x
	g.MaxY=y
	g.Points=make([][]Point,0)
	g.Blocks=make(map[string]*Point,0)
	for i:=0;i<y;i++ {
		ps:=make([]Point,0)
		for j:=0;j<x;j++ {
			ps=append(ps, Point{Key:pointAsKey(j,i),X:j,Y:i,View:PLAIN})
		}
		g.Points=append(g.Points,ps)
	}
	return g
}

// 设置起点
func(m *Graph) SetStart(x,y int) *Point {
	m.Set(x,y,START)
	return &m.Points[y][x]
}

// 设置终点
func(m *Graph) SetEnd(x,y int) *Point {
	m.Set(x,y,END)
	return &m.Points[y][x]
}

// 设置路径
func(m *Graph) SetRoad(x,y int) *Point {
	m.Set(x,y,PATH)
	return &m.Points[y][x]
}

// 设置障碍点
func(m *Graph) SetBlock(x,y int) {
	m.Set(x,y,BLOCK)
	m.Blocks[m.Points[y][x].Key]=&m.Points[y][x]
}

// 设置标识
func(m *Graph) Set(x,y int,view string) {
	m.Points[y][x].View=view
}

// 批量设置标识
func(m *Graph) Sets(ps []Point) {
	for _, p := range ps {
		m.Set(p.X,p.Y,p.View)
	}
}

// 获取附近的可能节点
func(m *Graph) GetAdjacentPoint(curPoint *Point) []*Point {
	adjacents:=make([]*Point,0)
	//    上
	if x, y := curPoint.X, curPoint.Y-1; x >= 0 && x < m.MaxX && y >= 0 && y < m.MaxY {
		adjacents = append(adjacents, &m.Points[y][x])
	}
	// 左 上
	if x, y := curPoint.X+1, curPoint.Y-1; x >= 0 && x < m.MaxX && y >= 0 && y < m.MaxY {
		adjacents = append(adjacents, &m.Points[y][x])
	}
	// 左
	if x, y := curPoint.X+1, curPoint.Y; x >= 0 && x < m.MaxX && y >= 0 && y < m.MaxY {
		adjacents = append(adjacents, &m.Points[y][x])
	}
	// 左 下
	if x, y := curPoint.X+1, curPoint.Y+1; x >= 0 && x < m.MaxX && y >= 0 && y < m.MaxY {
		adjacents = append(adjacents, &m.Points[y][x])
	}
	//    下
	if x, y := curPoint.X, curPoint.Y+1; x >= 0 && x < m.MaxX && y >= 0 && y < m.MaxY {
		adjacents = append(adjacents, &m.Points[y][x])
	}
	// 右 下
	if x, y := curPoint.X-1, curPoint.Y+1; x >= 0 && x < m.MaxX && y >= 0 && y < m.MaxY {
		adjacents = append(adjacents, &m.Points[y][x])
	}
	// 右
	if x, y := curPoint.X-1, curPoint.Y; x >= 0 && x < m.MaxX && y >= 0 && y < m.MaxY {
		adjacents = append(adjacents, &m.Points[y][x])
	}
	// 右 上
	if x, y := curPoint.X-1, curPoint.Y-1; x >= 0 && x < m.MaxX && y >= 0 && y < m.MaxY {
		adjacents = append(adjacents, &m.Points[y][x])
	}

	return adjacents
}

// 输出图像
func (m *Graph) Print(){
	for _,ps := range m.Points {
		print("|")
		for _, p := range ps {
			print(p.View)
		}
		println("|")
	}
}
