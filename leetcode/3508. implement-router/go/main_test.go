package leetcode3508

import "testing"

func TestRouter(t *testing.T) {
	router := Constructor(3)

	if !router.AddPacket(1, 4, 90) {
		t.Errorf("AddPacket(1, 4, 90) = false, want true")
	}

	if !router.AddPacket(2, 5, 90) {
		t.Errorf("AddPacket(2, 5, 90) = false, want true")
	}

	if router.AddPacket(1, 4, 90) {
		t.Errorf("AddPacket(1, 4, 90) = true, want false")
	}

	if !router.AddPacket(3, 5, 95) {
		t.Errorf("AddPacket(3, 5, 95) = false, want true")
	}

	if !router.AddPacket(4, 5, 105) {
		t.Errorf("AddPacket(4, 5, 105) = false, want true")
	}

	packet := router.ForwardPacket()
	if len(packet) != 3 || packet[0] != 2 || packet[1] != 5 || packet[2] != 90 {
		t.Errorf("ForwardPacket() = %v, want [2, 5, 90]", packet)
	}

	if !router.AddPacket(5, 2, 110) {
		t.Errorf("AddPacket(5, 2, 110) = false, want true")
	}

	count := router.GetCount(5, 100, 110)
	if count != 1 {
		t.Errorf("GetCount(5, 100, 110) = %d, want 1", count)
	}
}

func TestRouter2(t *testing.T) {
	router := Constructor(2)

	if !router.AddPacket(7, 4, 90) {
		t.Errorf("AddPacket(7, 4, 90) = false, want true")
	}

	result := router.ForwardPacket()
	if len(result) != 3 || result[0] != 7 || result[1] != 4 || result[2] != 90 {
		t.Errorf("ForwardPacket() = %v, want [7, 4, 90]", result)
	}

	result2 := router.ForwardPacket()
	if len(result2) != 0 {
		t.Errorf("ForwardPacket() = %v, want []", result2)
	}
}

func TestRouter3(t *testing.T) {
	router := Constructor(4)

	if !router.AddPacket(3, 4, 1) {
		t.Errorf("AddPacket(3,4,1) = false, want true")
	}

	if !router.AddPacket(5, 2, 1) {
		t.Errorf("AddPacket(5,2,1) = false, want true")
	}
}

func TestRouter4(t *testing.T) {
	router := Constructor(52)

	if !router.AddPacket(40, 32, 3) {
		t.Errorf("AddPacket(40, 32, 3) = false, want true")
	}

	if !router.AddPacket(14, 37, 5) {
		t.Errorf("AddPacket(14, 37, 5) = false, want true")
	}

	if !router.AddPacket(29, 20, 5) {
		t.Errorf("AddPacket(29, 20, 5) = false, want true")
	}

	count := router.GetCount(37, 3, 5)
	if count != 1 {
		t.Errorf("GetCount(37, 3, 5) = %d, want 1", count)
	}

	count = router.GetCount(32, 2, 3)
	if count != 1 {
		t.Errorf("GetCount(32, 2, 3) = %d, want 1", count)
	}

	count = router.GetCount(32, 2, 3)
	if count != 1 {
		t.Errorf("GetCount(32, 2, 3) = %d, want 1", count)
	}

	if !router.AddPacket(8, 16, 5) {
		t.Errorf("AddPacket(8, 16, 5) = false, want true")
	}

	if !router.AddPacket(21, 2, 5) {
		t.Errorf("AddPacket(21, 2, 5) = false, want true")
	}

	count = router.GetCount(2, 5, 5)
	if count != 1 {
		t.Errorf("GetCount(2, 5, 5) = %d, want 1", count)
	}

	count = router.GetCount(20, 1, 5)
	if count != 1 {
		t.Errorf("GetCount(20, 1, 5) = %d, want 1", count)
	}

	if !router.AddPacket(2, 44, 5) {
		t.Errorf("AddPacket(2, 44, 5) = false, want true")
	}

	count = router.GetCount(37, 1, 1)
	if count != 0 {
		t.Errorf("GetCount(37, 1, 1) = %d, want 0", count)
	}

	if !router.AddPacket(38, 25, 5) {
		t.Errorf("AddPacket(38, 25, 5) = false, want true")
	}

	count = router.GetCount(32, 5, 5)
	if count != 0 {
		t.Errorf("GetCount(32, 5, 5) = %d, want 0", count)
	}

	count = router.GetCount(2, 5, 5)
	if count != 1 {
		t.Errorf("GetCount(2, 5, 5) = %d, want 1", count)
	}

	packet := router.ForwardPacket()
	if len(packet) != 3 || packet[0] != 40 || packet[1] != 32 || packet[2] != 3 {
		t.Errorf("ForwardPacket() = %v, want [40, 32, 3]", packet)
	}

	if !router.AddPacket(32, 25, 5) {
		t.Errorf("AddPacket(32,25,5) = false, want true")
	}

	if !router.AddPacket(17, 42, 9) {
		t.Errorf("AddPacket(17,42,9) = false, want true")
	}

	count = router.GetCount(42, 8, 9)
	if count != 1 {
		t.Errorf("GetCount(42,8,9) = %d, want 1", count)
	}

	if !router.AddPacket(9, 5, 9) {
		t.Errorf("AddPacket(9, 5, 9) = false, want true")
	}

	count = router.GetCount(42, 8, 8)
	if count != 0 {
		t.Errorf("GetCount(42, 8, 8) = %d, want 0", count)
	}

	count = router.GetCount(16, 1, 8)
	if count != 1 {
		t.Errorf("GetCount(16, 1, 8) = %d, want 1", count)
	}

	if !router.AddPacket(12, 1, 9) {
		t.Errorf("AddPacket(12, 1, 9) = false, want true")
	}

	count = router.GetCount(42, 6, 7)
	if count != 0 {
		t.Errorf("GetCount(42, 6, 7) = %d, want 0", count)
	}

	if !router.AddPacket(19, 11, 9) {
		t.Errorf("AddPacket(19, 11, 9) = false, want true")
	}

	count = router.GetCount(1, 5, 5)
	if count != 0 {
		t.Errorf("GetCount(1, 5, 5) = %d, want 0", count)
	}

	packet = router.ForwardPacket()
	if len(packet) != 3 || packet[0] != 14 || packet[1] != 37 || packet[2] != 5 {
		t.Errorf("ForwardPacket() = %v, want [14, 37, 5]", packet)
	}

	count = router.GetCount(16, 8, 8)
	if count != 0 {
		t.Errorf("GetCount(16, 8, 8) = %d, want 0", count)
	}

	if !router.AddPacket(7, 9, 9) {
		t.Errorf("AddPacket(7, 9, 9) = false, want true")
	}

	if !router.AddPacket(39, 36, 9) {
		t.Errorf("AddPacket(39, 36, 9) = false, want true")
	}

	if !router.AddPacket(9, 31, 9) {
		t.Errorf("AddPacket(9, 31, 9) = false, want true")
	}
}

func TestRouter5(t *testing.T) {
	router := Constructor(31)

	if !router.AddPacket(4, 2, 1) {
		t.Errorf("AddPacket(4, 2, 1) = false, want true")
	}

	if !router.AddPacket(2, 1, 1) {
		t.Errorf("AddPacket(2, 1, 1) = false, want true")
	}

	if !router.AddPacket(2, 4, 1) {
		t.Errorf("AddPacket(2, 4, 1) = false, want true")
	}

	if !router.AddPacket(3, 1, 1) {
		t.Errorf("AddPacket(3, 1, 1) = false, want true")
	}

	packet := router.ForwardPacket()
	if len(packet) != 3 || packet[0] != 4 || packet[1] != 2 || packet[2] != 1 {
		t.Errorf("ForwardPacket() = %v, want [4, 2, 1]", packet)
	}

	if !router.AddPacket(1, 2, 2) {
		t.Errorf("AddPacket(1, 2, 2) = false, want true")
	}

	count := router.GetCount(2, 1, 1)
	if count != 0 {
		t.Errorf("GetCount(2, 1, 1) = %d, want 0", count)
	}

	count = router.GetCount(2, 2, 2)
	if count != 1 {
		t.Errorf("GetCount(2, 2, 2) = %d, want 1", count)
	}

	if !router.AddPacket(3, 4, 5) {
		t.Errorf("AddPacket(3, 4, 5) = false, want true")
	}

	count = router.GetCount(2, 1, 2)
	if count != 1 {
		t.Errorf("GetCount(2, 1, 2) = %d, want 1", count)
	}

	if !router.AddPacket(2, 4, 5) {
		t.Errorf("AddPacket(2, 4, 5) = false, want true")
	}

	count = router.GetCount(1, 3, 5)
	if count != 0 {
		t.Errorf("GetCount(1, 3, 5) = %d, want 0", count)
	}

	if !router.AddPacket(4, 1, 9) {
		t.Errorf("AddPacket(4, 1, 9) = false, want true")
	}

	if !router.AddPacket(3, 4, 9) {
		t.Errorf("AddPacket(3, 4, 9) = false, want true")
	}

	count = router.GetCount(1, 7, 9)
	if count != 1 {
		t.Errorf("GetCount(1, 7, 9) = %d, want 1", count)
	}

	count = router.GetCount(1, 6, 9)
	if count != 1 {
		t.Errorf("GetCount(1, 6, 9) = %d, want 1", count)
	}

	if !router.AddPacket(3, 1, 9) {
		t.Errorf("AddPacket(3, 1, 9) = false, want true")
	}

	packet = router.ForwardPacket()
	if len(packet) != 3 || packet[0] != 2 || packet[1] != 1 || packet[2] != 1 {
		t.Errorf("ForwardPacket() = %v, want [2, 1, 1]", packet)
	}

	packet = router.ForwardPacket()
	if len(packet) != 3 || packet[0] != 2 || packet[1] != 4 || packet[2] != 1 {
		t.Errorf("ForwardPacket() = %v, want [2, 4, 1]", packet)
	}

	count = router.GetCount(4, 5, 7)
	if count != 2 {
		t.Errorf("GetCount(4, 5, 7) = %d, want 2", count)
	}
}
