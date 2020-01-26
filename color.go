package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	PrintColor = "\033[38;5;%dm%s\033[39;49m"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected argument(you can set it's color and colored indexes or letters)")
		os.Exit(1)
	} else {
		str := os.Args[1]
		var indexArr []int
		var argColor *string
		indexOfColoredLetter := ""

		if len(os.Args) >= 3 && os.Args[2][:7] == "--color" {
			argCmd := flag.NewFlagSet(str, flag.ExitOnError)
			argColor = argCmd.String("color", "white", "color in string")
			argCmd.Parse(os.Args[2:])

			if len(os.Args) >= 4 {
				indexOfColoredLetter = argCmd.Arg(0)
			}

			if indexOfColoredLetter == "" {
				for i := range str {
					indexArr = append(indexArr, i)
				}
			} else {
				indexArr = IndexOfColoredLetter(str, indexOfColoredLetter)
				if indexArr == nil {
					os.Exit(1)
				}
			}

			//fmt.Println(*argColor)
			//fmt.Println(indexOfColoredLetter)
			//fmt.Println(indexArr)
			//return
		}

		// "\n" handling
		splittwo := string(byte(92)) + string(byte(110))
		words := strings.Split(str, splittwo)

		// read from file
		fileName := "standard.txt"

		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		defer file.Close()

		rawBytes, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}

		lines := strings.Split(string(rawBytes), "\n")

		//print string to terminal
		for _, word := range words {
			for h := 0; h < 9; h++ {
				for indexLetter, l := range []byte(word) {
					flagPrint := false
					if indexArr != nil {
						for _, indx := range indexArr {
							if indx == indexLetter {
								flagPrint = true
							}
						}
						for i, line := range lines {
							if i == (int(l)-32)*9+h {

								if flagPrint == true {

									fmt.Printf(PrintColor, Color(*argColor), line)
								} else {
									fmt.Printf(PrintColor, 7, line)
								}
							}
						}
					}
				}
				fmt.Println()
			}
		}
		if Color(*argColor) == 7 {
			fmt.Printf("Please use this colors:\n")
			fmt.Printf(PrintColor, 1, "red\n")
			fmt.Printf(PrintColor, 2, "green\n")
			fmt.Printf(PrintColor, 3, "yellow\n")
			fmt.Printf(PrintColor, 4, "blue\n")
			fmt.Printf(PrintColor, 5, "purple\n")
			fmt.Printf(PrintColor, 6, "magenta\n")
			fmt.Printf(PrintColor, 7, "white\n")
			fmt.Printf(PrintColor, 130, "orange\n")
		}
	}
}

func IndexOfColoredLetter(s string, colorStr string) []int {
	flagWord := false
	var indexArr []int
	for i, letter := range []byte(s) {
		if letter == colorStr[0] {
			lenColor := 0
			j := i
			for _, l := range []byte(colorStr) {
				if s[j] == l {
					lenColor++
				}
				j++
			}
			if lenColor == len(colorStr) {
				flagWord = true
				for k := i; k < i+len(colorStr); k++ {
					indexArr = append(indexArr, k)
				}
				return indexArr
			}
		}
	}

	if flagWord == false {
		indexStart := 0
		indexEnd := len(s) - 1
		index, err := strconv.Atoi(colorStr)
		if err == nil {
			indexArr = append(indexArr, index)
		} else {
			flagFormat := false

			for i, l := range colorStr {
				if l == ':' {
					flagFormat = true
					if i-1 < 0 && i+1 < len(colorStr) {
						indexEnd1, err2 := strconv.Atoi(colorStr[i+1:])
						indexEnd = indexEnd1
						if err2 != nil {
							fmt.Printf("Parse error: Please provide with last index of letters to be colored in format: \":number\"")
							return indexArr
						}
					} else if i-1 >= 0 && i+1 >= len(colorStr) {
						indexStart1, err1 := strconv.Atoi(colorStr[:i])
						indexStart = indexStart1
						if err1 != nil {
							fmt.Printf("Parse error: Please provide with first index of letters to be colored in format: \"number:\"\n")
							return indexArr
						}
					} else if i-1 >= 0 && i+1 < len(colorStr) {
						indexStart1, err1 := strconv.Atoi(colorStr[:i])
						indexEnd1, err2 := strconv.Atoi(colorStr[i+1:])
						indexStart = indexStart1
						indexEnd = indexEnd1
						if err1 != nil || err2 != nil {
							fmt.Printf("Parse error: Please provide with first and last indexes of letters to be colored in format: \"number:number\"\n")
							return indexArr
						}
					} else {
						fmt.Printf("Parse error: Please provide with first or/and last indexes of letters to be colored in format: \"number:number\"\n")
						return indexArr
					}
					for i := indexStart; i <= indexEnd; i++ {
						indexArr = append(indexArr, i)
					}
					break
				} else if l == ',' {
					flagFormat = true
					indexes := strings.Split(colorStr, ",")
					for _, v := range indexes {
						i, err := strconv.Atoi(v)
						if err == nil {
							indexArr = append(indexArr, i)
						} else {
							fmt.Printf("Parse error: Please provide with indexes of letters to be colored in format: \"number,number,number...\"\n")
							return indexArr
						}
					}
					break
				}
			}
			if flagFormat == false {
				fmt.Printf("Error format: Please provide letters or indexes as in exapmle: 4:6; 1:; :5; 5,6,4,2; 11, etc.\n")
				return indexArr
			}
		}
	}

	return indexArr
}

func Color(s string) int {
	switch s {
	case "blue":
		return 4
	case "green":
		return 2
	case "red":
		return 1
	case "yellow":
		return 3
	case "purple":
		return 5
	case "magenta":
		return 6
	case "orange":
		return 130
	default:
		num, err := strconv.Atoi(s)
		if err != nil {
			return 7
		}
		return num
	}
}
