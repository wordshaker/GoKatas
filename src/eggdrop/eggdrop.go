package eggdrop

import (
)

func EggDrop(buildingHeight int, breakingFloor int) int{
	building := Build(buildingHeight, breakingFloor)
	floor := 0
	for i := 1; i < len(building); i ++ {
		if(building[i] == true){
			floor = i - 1
		}
	}
	return floor
}

func Build(max int, breakingFloor int) []bool{
		building := make([]bool, max)
		for i := breakingFloor; i < len(building); i ++ {
			building[i] = true
		}
		for y := breakingFloor; y > 0 ; y -- {
			building[y] = false
		}
	return building
}
