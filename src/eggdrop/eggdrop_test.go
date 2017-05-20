package eggdrop

import (
	"testing"
	"strconv"
)

func TestBestCase(t *testing.T){
	actual := EggDrop(100, 14)
	if actual[0] != 14 && actual[1] != 1 {
		t.Error("Floor :" + strconv.Itoa(actual[0]) + 
				"Number Of Drops :" + strconv.Itoa(actual[1]))
	}
}

func TestFloorOne(t *testing.T){
	actual := EggDrop(100, 1)
	if actual[0] != 1 && actual[1] != 2{
		t.Error("Floor :" + strconv.Itoa(actual[0]) + 
				"Number Of Drops :" + strconv.Itoa(actual[1]))
	}
}

func TestFloorTwo(t *testing.T){
	actual := EggDrop(100, 2)
	if actual[0] != 2 && actual[1] != 3{
		t.Error("Floor :" + strconv.Itoa(actual[0]) + 
				"Number Of Drops :" + strconv.Itoa(actual[1]))
	}
}

func TestFloorThree(t *testing.T){
	actual := EggDrop(100, 3)
	if actual[0] != 3 && actual[1] != 4{
		t.Error("Floor :" + strconv.Itoa(actual[0]) + 
				"Number Of Drops :" + strconv.Itoa(actual[1]))
	}
}

func TestFloorFour(t *testing.T){
	actual := EggDrop(100, 4)
	if actual[0] != 4 && actual[1] != 5{
		t.Error("Floor :" + strconv.Itoa(actual[0]) + 
				"Number Of Drops :" + strconv.Itoa(actual[1]))
	}
}


func TestFloorFive(t *testing.T){
	actual := EggDrop(100, 5)
	if actual[0] != 5 && actual[1] != 6{
		t.Error("Floor :" + strconv.Itoa(actual[0]) + 
				"Number Of Drops :" + strconv.Itoa(actual[1]))
	}
}