package hardinfo

import (
	"fmt"
	"testing"
)

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

func TestStrLen(t *testing.T) {
	strs := []string{"我是张三；", "hello word："}
	for _, str := range strs {
		t.Log("原始字符串=", str, "长度=", len(str))
	}
}

func TestWlanInfo(t *testing.T) {
	rst, err := WlanInfo("HUAWEI-Main")
	if err != nil {
		t.Error(err.Error())
	} else {
		for k, v := range rst {
			fmt.Println(k, ":", v)
		}
	}
}
