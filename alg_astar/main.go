package main

import(
	"alg_astar/models"
	"fmt"
)

func main(){

	g:=models.New(30,30)

	g.SetBlock(4,1)
	g.SetBlock(4,2)
	g.SetBlock(3,4)
	g.SetBlock(4,3)
	g.SetBlock(4,4)
	g.SetBlock(4,6)
	g.SetBlock(4,5)
	g.SetBlock(4,7)
	g.SetBlock(4,8)

	g.SetBlock(12,4)
	g.SetBlock(12,5)
	g.SetBlock(12,6)
	g.SetBlock(12,7)
	g.SetBlock(12,8)
	g.SetBlock(12,9)
	g.SetBlock(12,10)
	g.SetBlock(12,11)
	g.SetBlock(12,12)
	g.SetBlock(12,13)
	g.SetBlock(12,14)
	g.SetBlock(12,15)
	g.SetBlock(12,16)

	g.SetBlock(12,18)
	g.SetBlock(12,19)
	g.SetBlock(12,20)
	g.SetBlock(12,21)



	searchRoad := models.NewSearchRoad(1, 3, 18, 20, g)

	g.Print()
	if searchRoad.FindoutRoad() {
		fmt.Println("计算结果：")
		g.Print()
	} else {
		fmt.Println("无结果")
	}


}
