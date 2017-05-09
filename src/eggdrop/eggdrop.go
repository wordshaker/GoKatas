package eggdrop

import (
	"math/rand"
)

func EggDrop(buildingHeight int) int{
	building := Build(buildingHeight)
	floor := 0
	for i := 1; i < len(building); i ++ {
		if(building[i] == true){
			floor = i - 1
		}
	}
	return floor
}

func Build(max int) []bool{
		building := make([]bool, max)
		breaks := rand.Intn(max)
		for i := breaks; i < len(building); i ++ {
			building[i] = true
		}
		for y := breaks; y > 0 ; y -- {
			building[y] = false
		}
	return building
}