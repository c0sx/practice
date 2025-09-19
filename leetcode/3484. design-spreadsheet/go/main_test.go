package leetcode3484

import "testing"

func TestSpreadsheet(t *testing.T) {
	ss := Constructor(3)

	if val := ss.GetValue("=5+7"); val != 12 {
		t.Errorf("GetValue(=5+7) = %d; want 12", val)
	}

	ss.SetCell("A1", 10)
	if val := ss.GetValue("=A1+6"); val != 16 {
		t.Errorf("GetValue(=A1+6) = %d; want 16", val)
	}

	ss.SetCell("B2", 15)
	if val := ss.GetValue("=A1+B2"); val != 25 {
		t.Errorf("GetValue(=A1+B2) = %d; want 25", val)
	}

	ss.ResetCell("A1")
	if val := ss.GetValue("=A1+B2"); val != 15 {
		t.Errorf("GetValue(=A1+B2) = %d; want 15", val)
	}
}

func TestSpreadsheet2(t *testing.T) {
	ss := Constructor(458)

	if val := ss.GetValue("=O126+10272"); val != 10272 {
		t.Errorf("GetValue(=O126+10272) = %d; want 10272", val)
	}
}

func TestSpreadsheet3(t *testing.T) {
	ss := Constructor(159)

	if val := ss.GetValue("=17904+30717"); val != 48621 {
		t.Errorf("GetValue(=17904+30717) = %d; want 48621", val)
	}

	if val := ss.GetValue("=Z37+38471"); val != 38471 {
		t.Errorf("GetValue(=Z37+38471) = %d; want 38471", val)
	}
}

func TestSpreadsheet4(t *testing.T) {
	ss := Constructor(269)

	if val := ss.GetValue("=Z1+P71"); val != 0 {
		t.Errorf("GetValue(=17904+30717) = %d; want 0", val)
	}

	ss.ResetCell("Z258")

	if val := ss.GetValue("=29678+24203"); val != 53881 {
		t.Errorf("GetValue(=Z37+38471) = %d; want 53881", val)
	}

	ss.SetCell("I182", 32333)
}

func TestSpreadsheet5(t *testing.T) {
	ss := Constructor(832)

	ss.SetCell("Z463", 52348)

	if val := ss.GetValue("=91520+Z463"); val != 143868 {
		t.Errorf("GetValue(=91520+Z463) = %d; want 143868", val)
	}

	ss.SetCell("F91", 32639)

	if val := ss.GetValue("=43503+29118"); val != 72621 {
		t.Errorf("GetValue(=43503+29118) = %d; want 72621", val)
	}

	ss.ResetCell("Z463")
}

func TestSpreadsheet6(t *testing.T) {
	ss := Constructor(530)

	if val := ss.GetValue("=29483+15079"); val != 44562 {
		t.Errorf("GetValue(=29483+15079) = %d; want 44562", val)
	}

	if val := ss.GetValue("=A333+A135"); val != 0 {
		t.Errorf("GetValue(=A333+A135) = %d; want 0", val)
	}

	if val := ss.GetValue("=J215+P337"); val != 0 {
		t.Errorf("GetValue(=J215+P337) = %d; want 0", val)
	}

	ss.ResetCell("F241")
	ss.SetCell("Y104", 2018)
	ss.SetCell("O71", 2353)

	if val := ss.GetValue("=Y104+O71"); val != 4371 {
		t.Errorf("GetValue(=Y104+O71) = %d; want 4371", val)
	}

	if val := ss.GetValue("=73100+39834"); val != 112934 {
		t.Errorf("GetValue(=73100+39834) = %d; want 112934", val)
	}

	ss.SetCell("Y118", 58058)

	if val := ss.GetValue("=O71+Y104"); val != 4371 {
		t.Errorf("GetValue(=O71+Y104) = %d; want 4371", val)
	}

	ss.ResetCell("O71")
	ss.ResetCell("Y118")

	if val := ss.GetValue("=F254+J85"); val != 0 {
		t.Errorf("GetValue(=F254+J85) = %d; want 0", val)
	}
}

func TestSpreadsheet7(t *testing.T) {
	ss := Constructor(1000)

	ss.SetCell("A1000", 100000)

	if val := ss.GetValue("=A1000+0"); val != 100000 {
		t.Errorf("GetValue(=A1000+0) = %d; want 100000", val)
	}
}
