package hardinfo

import "testing"

func TestHardinfo(t *testing.T) {
	PrintHardInfo()
}

func TestBordInfo(t *testing.T) {
	bi := GetBoardInfo()
	for _, d := range bi {
		t.Logf("%+v", d)
	}
}

func TestGetDiskDriverInfo(t *testing.T) {
	bi := GetDiskDriverInfo()
	for _, d := range bi {
		t.Logf("%+v", d)
	}
}
