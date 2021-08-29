package singleton

import (
	"testing"
)

func TestGetlazyInstance(t *testing.T) {
	ins1 := GetLazyInstance()
	ins2 := GetLazyInstance()
	if ins1 != ins2 {
		t.Error("test get lazy instance fail")
	}
}

func BenchmarkGetLazyInstance(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = GetLazyInstance()
	}
}
