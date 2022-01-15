package config

import (
	"reflect"
	"testing"
)

func TestReadEnv(t *testing.T) {
	var tests []struct {
		name string
		want *AppConfig
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadEnv(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadEnv() = %v, want %v", got, tt.want)
			}

		})
	}
}
