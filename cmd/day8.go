/*
Copyright © 2022 Sam Wood <samwooddev@gmail.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// day8Cmd represents the day8 command
var day8Cmd = &cobra.Command{
	Use:   "day8",
	Short: "Advent of Code 2021 day 8 solutions",
}

// day8Part1Cmd represents the part1 command
var day8Part1Cmd = &cobra.Command{
	Use:     "part1",
	Short:   "Advent of Code 2021 day 8 part 1 solution",
	PreRunE: ValidateInput,
	RunE:    day8Part1,
}

// day8Part1 is the part 1 solution code
func day8Part1(cmd *cobra.Command, args []string) error {
	data, err := os.Open(input)
	if err != nil {
		return err
	}
	defer data.Close()

	var rawData []string
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		rawData = append(rawData, line)
	}

	var combinations []string
	var encodedOutput []string
	for _, line := range rawData {
		var c1, c2, c3, c4, c5, c6, c7, c8, c9, c10 string
		var o1, o2, o3, o4 string
		_, err := fmt.Sscanf(line,
			"%s %s %s %s %s %s %s %s %s %s | %s %s %s %s",
			&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10, &o1, &o2, &o3, &o4)
		if err != nil {
			return err
		}
		combinations = append(combinations, c1, c2, c3, c4, c5, c6, c7, c8, c9, c10)
		encodedOutput = append(encodedOutput, o1, o2, o3, o4)
	}

	count := 0
	for _, v := range encodedOutput {
		if len(v) == 2 || len(v) == 3 || len(v) == 4 || len(v) == 7 {
			count += 1
		}
	}

	fmt.Printf("Answer: %d\n", count)

	return nil
}

// day8Part2Cmd represents the part2 command
var day8Part2Cmd = &cobra.Command{
	Use:     "part2",
	Short:   "Advent of Code 2021 day 8 part 2 solution",
	PreRunE: ValidateInput,
	RunE:    day8Part2,
}

// day8Part2 is the part 2 solution code
func day8Part2(cmd *cobra.Command, args []string) error {
	data, err := os.Open(input)
	if err != nil {
		return err
	}
	defer data.Close()

	var rawData []string
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		rawData = append(rawData, line)
	}

	var solData []string
	for _, line := range rawData {
		var c1, c2, c3, c4, c5, c6, c7, c8, c9, c10 string
		var o1, o2, o3, o4 string
		_, err := fmt.Sscanf(line,
			"%s %s %s %s %s %s %s %s %s %s | %s %s %s %s",
			&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10, &o1, &o2, &o3, &o4)
		if err != nil {
			return err
		}
		solData = append(solData, c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, o1, o2, o3, o4)
	}

	allLines := make([][]string, len(solData)/14)
	c := 0
	for i := 0; i < len(solData); i += 14 {
		var line []string
		if i > len(solData)-14 {
			line = solData[i:]
		} else {
			line = solData[i : i+14]
		}
		allLines[c] = line
		c += 1
	}

	var decodedNums []int
	for _, rows := range allLines {
		mapping := make(map[int]string)
		mappingAlt := make(map[string]int)

		// Find the numbers with a unique length i.e. the "known numbers"
		for ic, col := range rows[0:10] {
			sCol := sortStringByCharacter(col)
			if ic > 9 {
				break
			}
			if len(sCol) == 2 {
				mapping[1] = sCol
				mappingAlt[sCol] = 1
			}
			if len(sCol) == 3 {
				mapping[7] = sCol
				mappingAlt[sCol] = 7
			}
			if len(sCol) == 4 {
				mapping[4] = sCol
				mappingAlt[sCol] = 4
			}
			if len(sCol) == 7 {
				mapping[8] = sCol
				mappingAlt[sCol] = 8
			}
		}

		// Now find the numbers we don't know and add them to the map
		for ic, col := range rows[0:10] {
			sCol := sortStringByCharacter(col)
			if ic > 9 {
				break
			}
			if len(sCol) == 5 {
				if containsAll(sCol, mapping[7]) {
					mapping[3] = sCol
					mappingAlt[sCol] = 3
					continue
				} else if containsN(mapping[4], sCol) == 2 {
					mapping[2] = sCol
					mappingAlt[sCol] = 2
					continue
				} else if containsN(mapping[4], sCol) == 3 {
					mapping[5] = sCol
					mappingAlt[sCol] = 5
					continue
				}
			}
			if len(sCol) == 6 {
				if containsAll(sCol, mapping[1]) && containsAll(sCol, mapping[3]) && containsAll(sCol, mapping[4]) && containsAll(sCol, mapping[5]) {
					mapping[9] = sCol
					mappingAlt[sCol] = 9
					continue
				} else if containsN(mapping[1], sCol) == 1 && containsAll(sCol, mapping[5]) {
					mapping[6] = sCol
					mappingAlt[sCol] = 6
					continue
				} else {
					mapping[0] = sCol
					mappingAlt[sCol] = 0
					continue
				}
			}
		}

		// Build the string of decoded numbers i.e. the last 4 digits of the input
		builder := strings.Builder{}
		for _, col := range rows[10:] {
			sCol := sortStringByCharacter(col)
			builder.WriteString(strconv.Itoa(mappingAlt[sCol]))
		}
		builderStringInt, _ := strconv.Atoi(builder.String())
		decodedNums = append(decodedNums, builderStringInt)
		builder.Reset()
	}

	// Sum all the decoded numbers to get the answer
	var answer int
	for _, num := range decodedNums {
		answer += num
	}

	fmt.Printf("Answer: %d\n", answer)

	return nil
}

// containsAll returns true if a given string contains all of a given set of characters
func containsAll(s, c string) bool {
	chars := strings.Split(c, "")
	for _, char := range chars {
		if !strings.Contains(s, char) {
			return false
		}
	}
	return true
}

// containsN returns the number of characters from a given string, that are in another given string
func containsN(from, in string) int {
	fromChars := strings.Split(from, "")
	c := 0
	for _, v := range fromChars {
		if strings.Contains(in, v) {
			c += 1
		}
	}
	return c
}

// stringToRuneSlice converts a string to a slice of runes
func stringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

// sortStringByCharacter sorts a string by character
func sortStringByCharacter(s string) string {
	r := stringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func init() {
	rootCmd.AddCommand(day8Cmd)
	day8Cmd.AddCommand(day8Part1Cmd)
	day8Cmd.AddCommand(day8Part2Cmd)

	day8Part1Cmd.Flags().StringVarP(&input, "input", "i", "", "solution input data")
	day8Part1Cmd.MarkFlagRequired("input")

	day8Part2Cmd.Flags().StringVarP(&input, "input", "i", "", "solution input data")
	day8Part2Cmd.MarkFlagRequired("input")
}
