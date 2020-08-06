package models

import (
	"container/heap"
)

// 开放点列表
type OpenList []*AstarPoint

// 长度
func(self OpenList)Len()int{
	return len(self)
}

// 求价值更低的点
func (self OpenList) Less(i, j int) bool {
	return self[i].f_Value < self[j].f_Value
}

// 换位
func (self OpenList) Swap(i, j int){
	self[i], self[j] = self[j], self[i]
}

// 放入
func (this *OpenList) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*this = append(*this, x.(*AstarPoint))
}

// 取出
func (this *OpenList) Pop() interface{} {
	old := *this
	n := len(old)
	x := old[n-1]
	*this = old[0 : n-1]
	return x
}


type SearchRoad struct {
	TheGraph  *Graph
	start   AstarPoint
	end     AstarPoint
	closeLi map[string]*AstarPoint
	openLi  OpenList
	openSet map[string]*AstarPoint
	TheRoad []*AstarPoint
}

func NewSearchRoad(startx, starty, endx, endy int, m *Graph) *SearchRoad {
	sr := &SearchRoad{}
	// 关联主图
	sr.TheGraph = m
	// 设置起点
	sr.start = *NewAstarPoint(m.SetStart(startx, starty), nil, nil)
	// 设置终点
	sr.end = *NewAstarPoint(m.SetEnd(endx, endy), nil, nil)
	// 路径集
	sr.TheRoad = make([]*AstarPoint, 0)
	// 待用列表
	sr.openSet = make(map[string]*AstarPoint, m.MaxX+m.MaxY)
	// 不可用列表
	sr.closeLi = make(map[string]*AstarPoint, m.MaxX+m.MaxY)

	heap.Init(&sr.openLi)
	heap.Push(&sr.openLi, &sr.start) // 首先把起点加入开放列表
	sr.openSet[pointAsKey(sr.start.X, sr.start.Y)] = &sr.start

	// 将障碍点放入关闭列表
	for k, v := range m.Blocks {
		sr.closeLi[k] = NewAstarPoint(v, nil, nil)
	}

	return sr
}

func (this *SearchRoad) FindoutRoad() bool {
	for len(this.openLi) > 0 {
		// 将节点从开放列表移到关闭列表当中。
		x := heap.Pop(&this.openLi)
		curPoint := x.(*AstarPoint)

		delete(this.openSet, pointAsKey(curPoint.X, curPoint.Y))
		this.closeLi[pointAsKey(curPoint.X, curPoint.Y)] = curPoint

		// 搜索附近的可能的点
		adjacs := this.TheGraph.GetAdjacentPoint(&curPoint.Point)

		for _, p := range adjacs {
			// 当前可能的点 转化 为AStar算法点
			cur_astar_p := NewAstarPoint(p, curPoint, &this.end)
			// 判断 当前算法点是否为终点
			if pointAsKey(cur_astar_p.X, cur_astar_p.Y) == pointAsKey(this.end.X, this.end.Y) {
				// 路径完成，循环单链表，在主图上标记出所有的点
				for cur_astar_p.PrevPoint != nil {
					// 添加到路径集合
					this.TheRoad = append(this.TheRoad, cur_astar_p)
					// 上一个 算法点
					cur_astar_p = cur_astar_p.PrevPoint
					// 标记为路径
					this.TheGraph.SetRoad(cur_astar_p.X,cur_astar_p.Y)
				}
				// 算法完成退出
				return true
			}

			// 判断 关闭列表中是否存在
			_, ok := this.closeLi[pointAsKey(p.X, p.Y)]
			// 不可取用的算法点
			if ok {
				continue
			}
			// 可用算法点
			existAP, ok := this.openSet[pointAsKey(p.X, p.Y)]
			if !ok {
				heap.Push(&this.openLi, cur_astar_p)
				this.openSet[pointAsKey(cur_astar_p.X, cur_astar_p.Y)] = cur_astar_p
			} else {

				oldFVal, oldPrev := existAP.Cal_F_Value(&this.end), existAP.PrevPoint
				existAP.PrevPoint = curPoint
				existAP.Cal_F_Value(&this.end)
				// 找出 新的节点的总价值 与 之前的节点总价值更低值
				if existAP.f_Value > oldFVal {
					existAP.PrevPoint = oldPrev
					existAP.f_Value = oldFVal
				}
			}

		}
	}

	return false
}
