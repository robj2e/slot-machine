package helpers

import "testing"

func TestWinLineMinusOne(t *testing.T) {
	type args struct {
		i          int
		reelLength int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "regular",
			args: args{i: 3, reelLength: 6},
			want: 2,
		},
		{
			name: "reverts back to end of slice",
			args: args{i: 0, reelLength: 6},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WinLineMinusOne(tt.args.i, tt.args.reelLength); got != tt.want {
				t.Errorf("WinLineMinusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWinLinePlusOne(t *testing.T) {
	type args struct {
		i          int
		reelLength int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "regular",
			args: args{i: 3, reelLength: 6},
			want: 4,
		},
		{
			name: "reverts back to beginning of slice",
			args: args{i: 5, reelLength: 6},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WinLinePlusOne(tt.args.i, tt.args.reelLength); got != tt.want {
				t.Errorf("WinLinePlusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayAllSameValue(t *testing.T) {
	type args struct {
		sortedStrs []string
	}
	tests := []struct {
		name    string
		args    args
		wantB   bool
		wantVal string
	}{
		{
			name:    "Check loser generally",
			args:    args{sortedStrs: []string{"A", "A", "J", "X"}},
			wantB:   false,
			wantVal: "No Win",
		},
		{
			name:    "Check winner Aces",
			args:    args{sortedStrs: []string{"A", "A", "A", "A"}},
			wantB:   true,
			wantVal: "A",
		},
		{
			name:    "Check loser all Jokers",
			args:    args{sortedStrs: []string{"X", "X", "X", "X"}},
			wantB:   false,
			wantVal: "No Win",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotB, gotVal := ArrayAllSameValue(tt.args.sortedStrs)
			if gotB != tt.wantB {
				t.Errorf("ArrayAllSameValue() gotB = %v, want %v", gotB, tt.wantB)
			}
			if gotVal != tt.wantVal {
				t.Errorf("ArrayAllSameValue() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
		})
	}
}

func TestStringPromptIntReturn(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringPromptIntReturn(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringPromptIntReturn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringPromptIntReturn() = %v, want %v", got, tt.want)
			}
		})
	}
}
