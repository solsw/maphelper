package maphelper

import (
	"reflect"
	"sort"
	"testing"

	"github.com/solsw/generichelper"
)

func getm1() map[int]string {
	m1 := make(map[int]string)
	m1[1] = "one"
	m1[2] = "two"
	m1[3] = "three"
	return m1
}

func TestTuples(t *testing.T) {
	type args struct {
		m map[int]string
	}
	tests := []struct {
		name string
		args args
		want []generichelper.Tuple2[int, string]
	}{
		{name: "nil",
			args: args{m: nil},
			want: nil,
		},
		{name: "1",
			args: args{m: getm1()},
			want: []generichelper.Tuple2[int, string]{{Item1: 1, Item2: "one"}, {Item1: 2, Item2: "two"}, {Item1: 3, Item2: "three"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Tuples(tt.args.m)
			sort.Slice(got, func(i, j int) bool { return got[i].Item1 < got[j].Item1 })
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tuples() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromTuples(t *testing.T) {
	type args struct {
		tt []generichelper.Tuple2[int, string]
	}
	tests := []struct {
		name string
		args args
		want map[int]string
	}{
		{name: "nil",
			args: args{tt: nil},
			want: nil,
		},
		{name: "1",
			args: args{tt: []generichelper.Tuple2[int, string]{{Item1: 1, Item2: "one"}, {Item1: 2, Item2: "two"}, {Item1: 3, Item2: "three"}}},
			want: getm1(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromTuples(tt.args.tt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromTuples() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuplesRoundTrip(t *testing.T) {
	m := getm1()
	if got := NewFromTuples(Tuples(m)); !reflect.DeepEqual(got, m) {
		t.Errorf("NewFromTuples(Tuples(m)) = %v, want %v", got, m)
	}
}
