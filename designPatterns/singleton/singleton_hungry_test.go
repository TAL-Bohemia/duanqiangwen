package singleton

import (
	"reflect"
	"testing"
)

func TestGetHungryInstance(t *testing.T) {
	tests := []struct {
		name string
		want *HungrySingleton
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHungryInstance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHungryInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}
