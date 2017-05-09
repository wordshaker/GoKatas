package eggdrop

import (
	"testing"
	"strconv"
	"math/rand"
)

func Test100(t *testing.T){
	var actual = EggDrop(100, 10)
	if actual < 0{
		t.Error("Fail.Actual :" + strconv.Itoa(actual))
	}
}

func RandomiseBreak(buildingHeight int) int {
	return rand.Intn(buildingHeight)
}