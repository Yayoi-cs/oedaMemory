package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage :go run main.go <file_path>")
		os.Exit(1)
	}
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error :::: ", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	sourceCode := ""
	for scanner.Scan() {
		sourceCode += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error :::: ", err)
		os.Exit(1)
	}
	functions := extractFunctions(sourceCode)
	for functionName, variables := range functions {
		fmt.Println("Function:", functionName)
		//fmt.Println("Variables:")
		maxLength := 26
		fmt.Printf("%-*s\n", maxLength, strings.Repeat("_", maxLength+1))
		for _, variable := range variables {
			parts := strings.SplitN(variable, " ", 2)
			var bytes int
			switch parts[0] {
			case "int", "float", "long":
				bytes = 4
			case "char":
				bytes = 1
			case "short":
				bytes = 2
			case "double":
				bytes = 8
			default:
				bytes = 0
				fmt.Println("Exit")
				os.Exit(1)
			}
			var name string
			if strings.Contains(parts[1], "=") {
				name = strings.SplitN(parts[1], "=", 2)[0]
			} else {
				name = strings.SplitN(parts[1], ";", 2)[0]
			}
			fmt.Printf("%-*s : ", 8, name)
			fmt.Printf("%-*s(%d bytes)|\n", 6, parts[0], bytes)
			//fmt.Printf("%-*s|\n", maxLength, variable)
		}
		fmt.Printf("%-*s\n", maxLength, strings.Repeat("¯", maxLength+1))
	}
}
func extractFunctions(sourceCode string) map[string][]string {
	functions := make(map[string][]string)
	re := regexp.MustCompile(`\w+\s+(\w+)\s*\([^)]*\)\s*{([^}]*)}`)
	matches := re.FindAllStringSubmatch(sourceCode, -1)
	for _, match := range matches {
		functionName := match[1]
		functionCode := match[2]
		variables := extractVariables(functionCode)
		functions[functionName] = variables
	}
	return functions
}

func extractVariables(functionCode string) []string {
	variables := make([]string, 0)
	re := regexp.MustCompile(`\b(int|char|double|float|short|long)\b\s+\w+\s*;`)
	re2 := regexp.MustCompile(`\b(int|char|double|float|short|long)\b\s+\w+\s*=\s*\S+\s*;`)
	matches := re.FindAllString(functionCode, -1)
	matches2 := re2.FindAllString(functionCode, -1)
	for _, match := range matches {
		variables = append(variables, match)
	}
	for _, match := range matches2 {
		variables = append(variables, match)
	}
	return variables
}
