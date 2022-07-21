package rpgforge

import (
	"testing"
)

func TestThrowDices(t *testing.T) {
	for x := 1; x < 100; x++ {
		roll := ThrowDices("3d6")
		sum := roll.Sum()
		t.Log(sum)
		if (sum < 3) || (sum > 18) {
			t.Errorf("Wrong result 3d6 roll = %d (%d)", sum, roll.result)
		}
	}
}
