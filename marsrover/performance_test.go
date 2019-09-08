package main

import
	"testing"

func Test_rover_moveForward(t *testing.T) {
	type fields struct {
		xPosition int
		yPosition int
		direction string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		want1  int
	}{
		{
			name: "MoveTowardsY-axis",
			fields: fields{
				xPosition: 1,
				yPosition: 2,
				direction: "W",
			},
			want: -1, want1: 0,
		},
		{
			name: "MoveTowardsX-axis",
			fields: fields{
				xPosition: 0,
				yPosition: 2,
				direction: "S",
			},
			want: -1, want1: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pos := &Rover{
				Direction: tt.fields.direction,
			}
			got, got1 := pos.moveForward()
			if got != tt.want {
				t.Errorf("moveForward() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("moveForward() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_moveRight(t *testing.T) {
	type fields struct {
		direction string
	}
	tests := []struct {
		name string
		fields fields
		want string
	}{
		{
			name: "TurnFirstRight",
			fields: fields{
				direction: "N",
			},
			want: "E",
		},

		{
			name: "TurnsSecondRight",
			fields: fields{
				direction: "E",
			},
			want: "S",
		},
		{
			name: "TurnsFourthRight",
			fields: fields{
				direction: "W",
			},
			want: "N",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pos := &Rover{
				Direction: tt.fields.direction,
			}
			if got := pos.moveRight(); got != tt.want {
				t.Errorf("moveLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}



func Test_moveLeft(t *testing.T) {
	type fields struct {
		direction string
	}
	tests := []struct {
		name string
		fields fields
		want string
	}{
		{
			name: "TurnFirstLeft",
			fields: fields{
				direction: "N",
			},
			want: "W",
		},

		{
			name: "TurnsSecondLeft",
			fields: fields{
				direction: "W",
			},
			want: "S",
		},
		{
			name: "TurnsThirdLeft",
			fields: fields{
				direction: "S",
			},
			want: "E",
		},
		{
			name: "TurnsFourthLeft",
			fields: fields{
				direction: "E",
			},
			want: "N",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pos := &Rover{
				Direction: tt.fields.direction,
			}
			if got := pos.moveLeft(); got != tt.want {
				t.Errorf("moveLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}


