package integertoroman

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 1",
			args: args{num: 3},
			want: "III",
		},
		{
			name: "test case 2",
			args: args{num: 4},
			want: "IV",
		},
		{
			name: "test case 3",
			args: args{num: 9},
			want: "IX",
		},
		{
			name: "test case 4",
			args: args{num: 58},
			want: "LVIII",
		},
		{
			name: "test case 5",
			args: args{num: 1994},
			want: "MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intToRoman2(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 1",
			args: args{num: 3},
			want: "III",
		},
		{
			name: "test case 2",
			args: args{num: 4},
			want: "IV",
		},
		{
			name: "test case 3",
			args: args{num: 9},
			want: "IX",
		},
		{
			name: "test case 4",
			args: args{num: 58},
			want: "LVIII",
		},
		{
			name: "test case 5",
			args: args{num: 1994},
			want: "MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman2(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
