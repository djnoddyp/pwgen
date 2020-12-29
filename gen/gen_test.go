package gen

import (
	"testing"
)

func TestFilterByLength(t *testing.T)  {
	names := []string{"bobby", "frank", "jim", "patrick", "alex"}
	result := filterByLength(names, 2, 4)
	len := len(result) 
	if len != 2 {
		t.Fatalf("Expected length=2 but was length=%v", len)
	}
}