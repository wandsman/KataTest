package main

import (
	"errors"
	"testing"
)

func Test_isValid(t *testing.T) {
	type args struct {
		tmp []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		// TODO: Add test cases.
		{
			name: "simple case arabic",
			args: args{
				tmp: []string{"2", "+", "2"},
			},
			wantErr: nil,
		},
		{
			name: "simple case roma",
			args: args{
				tmp: []string{"I", "+", "I"},
			},
			wantErr: nil,
		},
		{
			name: "error first arg",
			args: args{
				tmp: []string{"2", "+", "I"},
			},
			wantErr: errors.New("different number systems are used simultaneously"),
		},
		{
			name: "error second arg",
			args: args{
				tmp: []string{"I", "+", "2"},
			},
			wantErr: errors.New("different number systems are used simultaneously"),
		},
		{
			name: "error action",
			args: args{
				tmp: []string{"2", "%", "2"},
			},
			wantErr: errors.New("this character is not valid %"),
		},
		{
			name: "error too mush arg",
			args: args{
				tmp: []string{"1", "-", "2", "+", "3"},
			},
			wantErr: errors.New("too many arguments"),
		},
		{
			name: "yoy",
			args: args{
				tmp: []string{"X", "*", "X"},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := isValid(tt.args.tmp); err != nil && err.Error() != tt.wantErr.Error() {

				t.Errorf("isValid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_makeAction(t *testing.T) {
	type args struct {
		src []string
	}
	var tests = []struct {
		name    string
		args    args
		wantRes string
		wantErr error
	}{
		// TODO: Add test cases.
		{
			name: "simple test arabic",
			args: args{
				src: []string{"2", "+", "2"},
			},
			wantRes: "4",
			wantErr: nil,
		},
		{
			name: "simple test rom",
			args: args{
				src: []string{"II", "+", "II"},
			},
			wantRes: "IV",
			wantErr: nil,
		},
		{
			name: "divide zero",
			args: args{
				src: []string{"4", "/", "0"},
			},
			wantRes: "",
			wantErr: errors.New("division by zero"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := makeAction(tt.args.src)
			if err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("makeAction() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if gotRes != tt.wantRes {
				t.Errorf("makeAction() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_makeRom(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
		wantErr error
	}{
		// TODO: Add test cases.
		{
			name: "simple case 1",
			args: args{
				num: 2,
			},
			wantRes: "II",
			wantErr: nil,
		},
		{
			name: "error zero",
			args: args{
				num: 0,
			},
			wantRes: "",
			wantErr: errors.New("цифры 0 нет в римской системе"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := makeRom(tt.args.num)
			if (err != nil) && err.Error() != tt.wantErr.Error() {
				t.Errorf("makeRom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != tt.wantRes {
				t.Errorf("makeRom() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
