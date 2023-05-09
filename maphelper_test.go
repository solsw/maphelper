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

func TestKeys(t *testing.T) {
	type args struct {
		m map[int]string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "1",
			args: args{m: getm1()},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Keys(tt.args.m)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElements(t *testing.T) {
	type args struct {
		m map[int]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "1",
			args: args{m: getm1()},
			want: []string{"one", "three", "two"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Elements(tt.args.m)
			sort.Strings(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Elements() = %v, want %v", got, tt.want)
			}
		})
	}
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
		{name: "1",
			args: args{m: getm1()},
			want: []generichelper.Tuple2[int, string]{{1, "one"}, {2, "two"}, {3, "three"}},
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
