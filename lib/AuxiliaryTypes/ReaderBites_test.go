package AuxiliaryTypes

import (
	"bufio"
	"bytes"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func Test_bitesBuffer_NextBites(t *testing.T) {
	bites := bytes.NewBuffer(
		[]byte{43,43,21,3,4,5,66,7,8,9,12,13,141,12,12,3,45,66,7,88,9,3,2,1,23,5,3,4,5,6},
		)

	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Bite
	}{
		// TODO: Add test cases.
		{
			name: "base_test",
			fields: fields{
				Buffer: bites,
			},
			args: args{
				n: 40,
			},
			want:[]Bite{Bite(43),Bite(43),Bite(21),Bite(3),Bite(4),Bite(5),Bite(66),Bite(7),Bite(8),Bite(9)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bB := bitesBuffer{
				Buffer: tt.fields.Buffer,
			}
			if got := bB.NextBites(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextBites() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestReaderBites_Read(t *testing.T) {
	genBts := func() *ReaderBites {
		bts := make([]byte, 1000)
		rand.Seed(time.Now().Unix())
		for i := 0; i < 1000; i++ {
			bts[i] = byte(48 + rand.Intn(1))
		}

		return  NewReaderBites(bufio.NewReader(bytes.NewReader(bts)))
	}
	type fields struct {
		buf         *bitesBuffer
		readerSlice []byte
		reader      *bufio.Reader
	}
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		rtb  *ReaderBites
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name : "Base 888",
			rtb: genBts(),
			args:args{
				888,
			},
			want: 888,
			wantErr: false,

		},
		{
			name : "Base 50",
			rtb: genBts(),
			args:args{
				50,
			},
			want: 50,
			wantErr: false,

		},
		{
			name : "Base 100",
			rtb: genBts(),
			args:args{
				100,
			},
			want: 100,
			wantErr: false,

		},
		{
			name : "Base 400",
			rtb: genBts(),
			args:args{
				400,
			},
			want: 400,
			wantErr: false,

		},
		{
			name : "Base 801",
			rtb: genBts(),
			args:args{
				801,
			},
			want: 801,
			wantErr: false,

		},
		{
			name : "Base 1",
			rtb: genBts(),
			args:args{
				1,
			},
			want: 1,
			wantErr: false,

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := tt.rtb.Read(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("Read() got = %v, want %v", len(got), tt.want)
			}
		})
	}
}