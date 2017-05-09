package eggdrop

import (
	"testing"
	"strconv"
)

func Test100(t *testing.T){
	var actual = EggDrop(100)
	if actual < 0{
		t.Error("Fail.Actual :" + strconv.Itoa(actual))
	}
}

