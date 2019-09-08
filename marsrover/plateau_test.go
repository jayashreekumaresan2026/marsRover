package main

import "testing"

func Test_maxGridPlanet_isLesser(t *testing.T) {
	type fields struct {
		xMax int
		yMax int
	}
	type args struct {
		pos *Rover
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:"true case",
			fields: fields{
				xMax: 0,
				yMax: 2,
			},
			args:args{},
			want:true,
		},
		{
			name:"false case",
			fields: fields{
				xMax: 6,
				yMax: 2,
			},
			want:true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			max := &MaxGridPlanet{
				XMax: tt.fields.xMax,
				YMax: tt.fields.yMax,
			}
			if got := max.isLesser(tt.args.pos); got != tt.want {
				t.Errorf("isLesser() = %v, want %v", got, tt.want)
			}
		})
	}
}