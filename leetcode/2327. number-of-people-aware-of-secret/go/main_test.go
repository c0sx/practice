package leetcode2327

import "testing"

func TestPeopleAwareOfSecret(t *testing.T) {
	type args struct {
		n      int
		delay  int
		forget int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				n:      6,
				delay:  2,
				forget: 4,
			},
			want: 5,
		},
		// {
		// 	name: "test case 2",
		// 	args: args{
		// 		n:      4,
		// 		delay:  1,
		// 		forget: 3,
		// 	},
		// 	want: 6,
		// },
		// {
		// 	name: "test case 3",
		// 	args: args{
		// 		n:      684,
		// 		delay:  18,
		// 		forget: 496,
		// 	},
		// 	want: 653668527,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := peopleAwareOfSecret(tt.args.n, tt.args.delay, tt.args.forget); got != tt.want {
				t.Errorf("peopleAwareOfSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}
