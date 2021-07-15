package singleton

import (
	"reflect"
	"testing"
)

func TestGetLazyInstance(t *testing.T) {
	tests := []struct {
		name string
		want *LazySingleton
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLazyInstance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLazyInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}
