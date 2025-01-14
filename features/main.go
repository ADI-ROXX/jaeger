package features

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func beautify(data string) string {
	outputString := string(data)
	startIndex := strings.Index(outputString, "--feature-gates")
	if startIndex == -1 {
		fmt.Println("The string '--feature-gates' was not found.")
		os.Exit(1)
	}
	relativeLineBreak := strings.Index(outputString[startIndex:], "\n")
	var featureLine string
	if relativeLineBreak == -1 {
		featureLine = outputString[startIndex:]
	} else {
		featureLine = outputString[startIndex : startIndex+relativeLineBreak]
	}
	openParen := strings.Index(featureLine, "(")
	closeParen := strings.Index(featureLine, ")")
	if openParen == -1 || closeParen == -1 || closeParen <= openParen {
		fmt.Println("Could not find parentheses surrounding the feature gates data.")
		os.Exit(1)
	}
	inside := featureLine[openParen+1 : closeParen]
	splitData := strings.Split(inside, ",")
	combined_split := "A list of feature gate identifiers. Prefix with '-' to disable the feature. '+' or no prefix will enable the feature.\n"
	for _, item := range splitData {
		cleanItem := strings.TrimSpace(item)
		combined_split = combined_split + "\t" + cleanItem + "\n"
	}
	return combined_split
}
func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	outputString := string(data)
	startIndex := strings.Index(outputString, "--feature-gates")
	if startIndex == -1 {
		fmt.Print(outputString)
		return
	}
	relativeLineBreak := strings.Index(outputString[startIndex:], "\n")
	if relativeLineBreak == -1 {
		relativeLineBreak = len(outputString) - startIndex
	}
	lineBreakIndex := startIndex + relativeLineBreak
	extracted := outputString[startIndex:lineBreakIndex]
	// newString := outputString[:startIndex] + beautify(extracted) + outputString[lineBreakIndex:]
	fmt.Println(beautify(extracted))
}