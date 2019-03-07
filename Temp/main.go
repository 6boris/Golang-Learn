package main

func main() {

}

func getRandStr(len int, format string) string {

	char := ""
	switch format {
	case "ALL":
		char = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
		break
	case "CHAR":
		char = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
		break
	case "NUMBAER":
		char = "0123456789"
		break
	default:
		char = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
		break
	}
	return char
}
