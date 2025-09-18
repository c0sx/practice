package leetcode3408

import "testing"

func TestTaskManager(t *testing.T) {
	obj := Constructor([][]int{{1, 101, 10}, {2, 102, 20}, {3, 103, 15}})
	obj.Add(4, 104, 5)
	obj.Edit(102, 8)

	result1 := obj.ExecTop()
	if result1 != 3 {
		t.Errorf("expected %d, but got %d", 3, result1)
	}

	obj.Rmv(101)
	obj.Add(5, 105, 15)

	param_4 := obj.ExecTop()
	if param_4 != 5 {
		t.Errorf("expected %d, but got %d", 5, param_4)
	}
}

func TestTaskManager2(t *testing.T) {
	obj := Constructor([][]int{{7, 13, 6}, {3, 11, 41}, {7, 18, 40}})

	obj.Edit(13, 0)
	result0 := obj.ExecTop()
	if result0 != 3 {
		t.Errorf("expected %d, but got %d", 3, result0)
	}

	obj.Edit(13, 19)
	obj.Rmv(18)
	obj.Rmv(13)

	result1 := obj.ExecTop()
	if result1 != -1 {
		t.Errorf("expected %d, but got %d", -1, result1)
	}

	result2 := obj.ExecTop()
	if result2 != -1 {
		t.Errorf("expected %d, but got %d", 3, result2)
	}
}
