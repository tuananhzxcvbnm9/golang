package basic

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "both positive", a: 3, b: 5, want: 8},
		{name: "with zero", a: 0, b: 7, want: 7},
		{name: "negative", a: -2, b: 1, want: -1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Sum(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("Sum(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{n: 2, want: true},
		{n: 7, want: false},
		{n: 0, want: true},
		{n: -4, want: true},
	}

	for _, tc := range tests {
		got := IsEven(tc.n)
		if got != tc.want {
			t.Fatalf("IsEven(%d) = %v, want %v", tc.n, got, tc.want)
		}
	}
}
