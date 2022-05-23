package cache

import "testing"

func TestCache(t *testing.T) {
	//s:=map{"a":"dfg"}string
	NewCache().Put("a", "dfg")
	d := NewCache().Keys()
	if d == nil {
		t.Errorf("Error: ")
	}

}
