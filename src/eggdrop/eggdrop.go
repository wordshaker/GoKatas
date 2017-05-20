package eggdrop

import (
	"math"
)

func EggDrop(buildingHeight int, breakingFloor int) []int{
	checkFloor := int(StartingFloor(buildingHeight))
	building := Build(buildingHeight, breakingFloor)
	checks:= 0
	result:= []int{}
	for i := checkFloor; i < len(building); i =checkFloor+(checkFloor -1) {
		checks++
		result = CheckBuilding(checkFloor, building, checks)
	}
	return result
}

func CheckBuilding(checkFloor int, building []bool, checks int) []int{
	floor := 0
	if (building[checkFloor] == true){
		for i:= checkFloor; i == 0; i--{
			checks++
			if(building[i] == true){
				floor = i - 1
			}			
		}
	}
	return []int{floor, checks}
}

func Build(max int, breakingFloor int) []bool{
		building := make([]bool, max)
		for i := breakingFloor; i < len(building); i ++ {
			building[i] = true
		}
		for y := breakingFloor-1; y > 0 ; y -- {
			building[y] = false
		}
	return building
}

func StartingFloor(buildingHeight int) float64{
	bh := float64(buildingHeight)
	return (-1 + math.Sqrt(1+ (8*bh)))/2
}