package services

import (
	"reflect"
	"repartners/internal/pack"
	"repartners/internal/pack/repository"
	"testing"
)

func TestCalculate(t *testing.T) {
	mr := repository.NewMemoryRepo()

	packs, _ := mr.GetAll()

	type args struct {
		items int
		packs []pack.PackSize
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{
			name: "1",
			args: args{
				items: 1,
				packs: packs,
			},
			want: map[int]int{250: 1},
		},
		{
			name: "250",
			args: args{
				items: 250,
				packs: packs,
			},
			want: map[int]int{250: 1},
		},
		{
			name: "251",
			args: args{
				items: 251,
				packs: packs,
			},
			want: map[int]int{500: 1},
		},
		{
			name: "501",
			args: args{
				items: 501,
				packs: packs,
			},
			want: map[int]int{500: 1, 250: 1},
		},
		{
			name: "12001",
			args: args{
				items: 12001,
				packs: packs,
			},
			want: map[int]int{5000: 2, 2000: 1, 250: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calculate(tt.args.items, tt.args.packs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
