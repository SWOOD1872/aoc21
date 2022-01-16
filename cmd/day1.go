/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	input string // path to input data

	ErrEmptyInput error = errors.New("input should not be empty")
)

// day1Cmd represents the day1 command
var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "Advent of Code 2021 day 1 solutions",
}

// part1Cmd represents the part1 command
var part1Cmd = &cobra.Command{
	Use:     "part1",
	Short:   "Advent of Code 2021 day 1 part 1 solution",
	PreRunE: validateInput,
	RunE:    part1,
}

// part1 is the day 1 part 1 solution code
func part1(cmd *cobra.Command, args []string) error {
	data, err := os.Open(input)
	if err != nil {
		return err
	}
	defer data.Close()

	curNo := 0
	preNo := 0
	countGreater := 0
	i := 0

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		if i == 0 {
			i += 1
			preNo = curNo
			continue
		}

		curNo, err = strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}

		if curNo > preNo {
			countGreater += 1
		}

		preNo = curNo
		i += 1
	}

	fmt.Printf("Answer = %v\n", countGreater)

	return nil
}

// part2Cmd represents the part2 command
var part2Cmd = &cobra.Command{
	Use:     "part2",
	Short:   "Advent of Code 2021 day 1 part 2 solution",
	PreRunE: validateInput,
	RunE:    part2,
}

// part2 is the day 1 part 2 solution code
func part2(cmd *cobra.Command, args []string) error {
	data, err := os.Open(input)
	if err != nil {
		return err
	}
	defer data.Close()

	var n1, n2, n3 int = 0, 0, 0
	countGreater := 0

	scanner := bufio.NewScanner(data)

	// Need the first 3 numbers initially so we can start comparing and shifting them in Scan()
	scanner.Scan()
	n3, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	n2, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	n1, _ = strconv.Atoi(scanner.Text())

	for scanner.Scan() {
		// Get the next number from the input
		n, _ := strconv.Atoi(scanner.Text())

		// If the next/latest number is greater than oldest (n+3), increment the counter
		if n > n3 {
			countGreater++
		}

		// Ensure we shift each number for the next loop
		n3 = n2
		n2 = n1
		n1 = n
	}

	fmt.Printf("Answer: %v\n", countGreater)

	return nil
}

// validateInput does validation on the input value and converts it to an absolute path
func validateInput(cmd *cobra.Command, args []string) error {
	if input == "" {
		return ErrEmptyInput
	}

	var err error
	input, err = filepath.Abs(input)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(day1Cmd)
	day1Cmd.AddCommand(part1Cmd)
	day1Cmd.AddCommand(part2Cmd)

	part1Cmd.Flags().StringVarP(&input, "input", "i", "", "solution input data")
	part1Cmd.MarkFlagRequired("input")

	part2Cmd.Flags().StringVarP(&input, "input", "i", "", "solution input data")
	part2Cmd.MarkFlagRequired("input")
}
