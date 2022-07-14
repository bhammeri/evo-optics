package utils

import (
	"testing"
)

func TestVector_ScaleToUnitVector(t *testing.T) {
	type fields struct {
		Origin Point
		End    Point
	}
	type want struct {
		vector Vector
		equal  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   want
	}{
		{
			name:   "1",
			fields: fields{Origin: Point{300, 300}, End: Point{400, 400}},
			want: want{
				Vector{Origin: Point{300, 300}, End: Point{300.7071, 300.7071}},
				false,
			},
		},
		{
			name:   "1",
			fields: fields{Origin: Point{0, 0}, End: Point{1, 1}},
			want: want{
				Vector{Origin: Point{0, 0}, End: Point{0.7071067811, 0.7071067811}},
				true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vector := &Vector{
				Origin: tt.fields.Origin,
				End:    tt.fields.End,
			}
			vector.ScaleToUnitVector()
			if got := vector; got.EqualTo(&tt.want.vector) != tt.want.equal {
				t.Errorf("Vector.ScaleToUnitVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_ScaleToUnitVector_NotEqual(t *testing.T) {
	type fields struct {
		Origin Point
		End    Point
	}
	tests := []struct {
		name   string
		fields fields
		want   Vector
	}{
		{
			name:   "1",
			fields: fields{Origin: Point{300, 300}, End: Point{400, 400}},
			want:   Vector{Origin: Point{300, 300}, End: Point{300.7071, 300.7071}},
		},
		{
			name:   "1",
			fields: fields{Origin: Point{0, 0}, End: Point{1, 1}},
			want:   Vector{Origin: Point{0, 0}, End: Point{0.7071, 0.7071}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vector := &Vector{
				Origin: tt.fields.Origin,
				End:    tt.fields.End,
			}
			vector.ScaleToUnitVector()
			if got := vector; got.EqualTo(&tt.want) {
				t.Errorf("Vector.ScaleToUnitVector() = %v, want %v", got, tt.want)
			}
		})
	}
}
