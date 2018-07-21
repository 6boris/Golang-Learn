package ConvertToBin

import (
	"fmt"
	"testing"
)

func TestConvertToBin(t *testing.T) {
	fmt.Println("5: ", ConvertToBin(5))
	fmt.Println("7: ", ConvertToBin(7))
	fmt.Println("13: ", ConvertToBin(13))
	fmt.Println("123123123: ", ConvertToBin(123123123))
	fmt.Println("0: ", ConvertToBin(0))
	fmt.Println("-1: ", ConvertToBin(-1))
}
