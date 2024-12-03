package nocancer

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	f, err := os.Open("example2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scnr := bufio.NewScanner(f)
	var task1 int

	var t2String string
	fmt.Println("calculating t1")
	for scnr.Scan() {
		t := scnr.Text()
		t1String := strings.Clone(t)
		t2String += t
		for {
			ind := strings.Index(t1String, "mul(")
			if ind == -1 {
				break
			}
			task1 += calcMultiply(t1String, ind)
			t1String = t1String[ind+4:] // 4 == mul(
		}
	}

	fmt.Println("task1: ", task1) // 167090022
}

func calcMultiply(str string, ind int) int {
	workingString := str[ind:]
	after, _ := strings.CutPrefix(workingString, "mul(")
	digit1, digit2 := "", ""
	first := true
	for _, char := range after {
		if len(digit1) > 3 || len(digit2) > 3 {
			return 0
		}
		if unicode.IsDigit(char) && first {
			digit1 += string(char)
			continue
		}
		if unicode.IsDigit(char) && !first {
			digit2 += string(char)
			continue
		}
		if !unicode.IsDigit(char) && first && digit1 == "" {
			return 0
		}

		if char == ',' && first {
			first = false
			continue
		}

		if !unicode.IsDigit(char) && !first {
			if char == ')' && digit2 != "" {
				int1, _ := strconv.Atoi(digit1)
				int2, _ := strconv.Atoi(digit2)
				return int1 * int2
			}
			return 0
		}

	}
	return 0
}
