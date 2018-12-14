package v2

import (
	"fmt"
	"testing"
)

func TestPreKMP(t *testing.T) {
	ret := preKMP("cocacola")
	if len(ret) > 5 && (ret[0] != -1 || ret[2] != -1 || ret[4] != -1) {
		t.Error("PreKMP error:", ret)
	}

	fmt.Println("ORI= cocacola")
	fmt.Println("MP=", preMP("cocacola"))
	fmt.Println("KMP=", preKMP("cocacola"))

	fmt.Println("ORI= ABCDABD")
	fmt.Println("MP=", preMP("ABCDABD"))
	fmt.Println("KMP=", preKMP("ABCDABD"))

	fmt.Println("ORI= abcabcacab")
	fmt.Println("MP=", preMP("abcabcacab"))
	fmt.Println("KMP=", preKMP("abcabcacab"))
}

func TestKMP(t *testing.T) {
	if ret := KMP("co", "cococp"); len(ret) != 0 {
		t.Error("Input error detect failed")
	}

	ret := KMP("cocacola", "co")
	if len(ret) == 0 {
		t.Error("No result on KMP")
		return
	}

	if len(ret) != 2 && (ret[0] != 0 || ret[1] != 4) {
		t.Error("Result is wrong on KMP 1st test:", ret)
		return
	}

	if r2 := KMP("ABC ABCDAB ABCDABCDABDE", "ABCDABD"); len(r2) > 0 && r2[0] != 15 {
		t.Error("wrong on 2nd KMP test:", r2)
	}

	if r3 := KMP("ACBDA", "ABCD"); len(r3) != 0 {
		t.Error("wrong on 3rd KMP", r3)
	}
}

func TestStrstr(t *testing.T) {
	if index := Strstr("hellow world", "low"); index != 3 {
		t.Error("wrong on 1st Strstr test:", index)
	}

	if index := Strstr("ABC ABCDAB ABCDABCDABDE", "ABCDABD"); index != 15 {
		t.Error("wrong on 2nd Strstr test:", index)
	}
}

func TestStrchr(t *testing.T) {
	if index := Strchr("hellow world", "w"); index != 7 {
		t.Error("wrong on 1st Strchr test:", index)
	}

	if index := Strchr("ABC ABCDAB ABCDABCDABDE", "CDAB"); index != 17 {
		t.Error("wring on 2nd Strchr test:", index)
	}

}
