package AuxiliaryTypes

import (
	"reflect"
	"testing"
)

func TestChar_ToByte(t *testing.T) {
	tests := []struct {
		name string
		c    Char
		want []byte
	}{
		// TODO: Add test cases.
		{
			name: "base_test",
			c: Char('u'),
			want: []byte{117},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ToByte(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChar_ToStringByte(t *testing.T) {
	tests := []struct {
		name string
		c    Char
		want string
	}{
		// TODO: Add test cases.
		{
			name : "Base_test",
			c: Char(117),
			want: "01110101",
		},
		{
			name : "test_1",
			c: Char(22),
			want: "00010110",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ToStringByte(); got != tt.want {
				t.Errorf("ToStringByte() = %v, want %v", got, tt.want)
			}
		})
	}
}