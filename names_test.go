package rpgforge

import (
  "fmt"
	"testing"
)

func TestGetLastName(t *testing.T) {
	name := GetLastName()
  fmt.Printf("Got %s names \n", name)
	if len(name) < 1 {
		t.Errorf("Name empty")
	}
}
