/*
Copyright © 2022 Sam Wood <samwooddev@gmail.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// day3Cmd represents the day3 command
var day3Cmd = &cobra.Command{
	Use:   "day3",
	Short: "Advent of Code 2021 day 3 solutions",
}

// day3Part1Cmd represents the part1 command
var day3Part1Cmd = &cobra.Command{
	Use:     "part1",
	Short:   "Advent of Code 2021 day 3 part 1 solution",
	PreRunE: ValidateInput,
	RunE:    day3Part1,
}

// day3Part1 is the part 1 solution code
func day3Part1(cmd *cobra.Command, args []string) error {
	data, err := os.Open(input)
	if err != nil {
		return err
	}
	defer data.Close()

	var lines []string
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	gamma := ""
	epsilon := ""
	n := len(lines[0])
	z, o := 0, 0
	for i := 0; i < n; i++ {
		for _, line := range lines {
			if string(line[i]) == "0" {
				z += 1
			} else {
				o += 1
			}
		}
		if z > o {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		} else {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		}
		z = 0
		o = 0
	}
	gammaDec := binToDec(gamma)
	epsilonDec := binToDec(epsilon)

	fmt.Printf("Answer: %v\n", gammaDec*epsilonDec)

	return nil
}

// day3Part2Cmd represents the part2 command
var day3Part2Cmd = &cobra.Command{
	Use:     "part2",
	Short:   "Advent of Code 2021 day 3 part 2 solution",
	PreRunE: ValidateInput,
	RunE:    day3Part2,
}

// day3Part2 is the part 2 solution code
func day3Part2(cmd *cobra.Command, args []string) error {
	data, err := os.Open(input)
	if err != nil {
		return err
	}
	defer data.Close()

	var lines []string
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	oxygenRating := binToDec(rating(lines, "most"))
	co2Rating := binToDec(rating(lines, "least"))
	lifeSupportRating := oxygenRating * co2Rating
	fmt.Printf("Answer: %v\n", lifeSupportRating)

	return nil
}

// binToDec takes a binary string and return and integer value
func binToDec(s string) int {
	a, _ := strconv.ParseInt(s, 2, 64)
	return int(a)
}

func rating(list []string, s string) string {
	answer := ""
	n := len(list[0])
	z, o := 0, 0
	for i := 0; i < n; i++ {
		mcn := ""
		for _, line := range list {
			if string(line[i]) == "0" {
				z += 1
			} else {
				o += 1
			}
		}
		if s == "most" {
			if z > o {
				mcn = "0"
			} else {
				mcn = "1"
			}
		} else if s == "least" {
			if z > o {
				mcn = "1"
			} else {
				mcn = "0"
			}
		}
		z = 0
		o = 0
		list = filter(list, i, mcn)
		if len(list) == 1 {
			answer = list[0]
		}
	}
	return answer
}

func filter(list []string, pos int, filter string) []string {
	var filtered []string
	for _, item := range list {
		if string(item[pos]) == filter {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func init() {
	rootCmd.AddCommand(day3Cmd)
	day3Cmd.AddCommand(day3Part1Cmd)
	day3Cmd.AddCommand(day3Part2Cmd)

	day3Part1Cmd.Flags().StringVarP(&input, "input", "i", "", "solution input data")
	day3Part1Cmd.MarkFlagRequired("input")

	day3Part2Cmd.Flags().StringVarP(&input, "input", "i", "", "solution input data")
	day3Part2Cmd.MarkFlagRequired("input")
}
