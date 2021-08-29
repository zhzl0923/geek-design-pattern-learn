package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()
	if instance1 != instance2 {
		t.Error("test  get instance fail")
	}
}

func BenchmarkGetInstance(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = GetInstance()
	}
}
