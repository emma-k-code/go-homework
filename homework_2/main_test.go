package main

import (
	"testing"
)

// 使用 VS Code 產生的測試樣板
func TestGetMachineGame(t *testing.T) {
	type args struct {
		MID string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "有資料",
			args: args{MID: "1"},
			want: 1,
		},
		{
			name: "無資料",
			args: args{MID: "5"},
			want: 0,
		},
		{
			name: "換分比例異常",
			args: args{MID: "2"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMachineGame(tt.args.MID); got != tt.want {
				t.Errorf("%s, GetMachineGame(%v) = %v, want %v", tt.name, tt.args.MID, got, tt.want)
			}
		})
	}
}
