/*
Copyright © 2022 Sam Wood <samwooddev@gmail.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// day7Cmd represents the day7 command
var day7Cmd = &cobra.Command{
	Use:   "day7",
	Short: "Advent of Code 2021 day 7 solutions",
}

// day7Part1Cmd represents the part1 command
var day7Part1Cmd = &cobra.Command{
	Use:     "part1",
	Short:   "Advent of Code 2021 day 7 part 1 solution",
	PreRunE: ValidateInput,
	RunE:    day7Part1,
}

// day7Part1 is the part 1 solution code
func day7Part1(cmd *cobra.Command, args []string) error {
	data, err := os.Open(input)
	if err != nil {
		return err
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	var horizontalPositions []int
	var num int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		for _, numStr := range line {
			num, err = strconv.Atoi(numStr)
			if err != nil {
				return err
			}
			horizontalPositions = append(horizontalPositions, num)
		}
	}
	sort.Ints(horizontalPositions)

	minPos := horizontalPositions[0]
	maxPos := horizontalPositions[len(horizontalPositions)-1]
	var costs []int
	for i := minPos; i < maxPos; i++ {
		cost := 0
		for _, p := range horizontalPositions {
			cost += int(math.Abs(float64(p - i)))
		}
		costs = append(costs, cost)
	}
	sort.Ints(costs)

	fmt.Printf("Answer: %d\n", costs[0])

	return nil
}

// day7Part2Cmd represents the part2 command
var day7Part2Cmd = &cobra.Command{
	Use:     "part2",
	Short:   "Advent of Code 2021 day 7 part 2 solution",
	PreRunE: ValidateInput,
	RunE:    day7Part2,
}

// day7Part2 is the part 2 solution code
func day7Part2(cmd *cobra.Command, args []string) error {
	data, err := os.Open(input)
	if err != nil {
		return err
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	var horizontalPositions []int
	var num int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		for _, numStr := range line {
			num, err = strconv.Atoi(numStr)
			if err != nil {
				return err
			}
			horizontalPositions = append(horizontalPositions, num)
		}
	}
	sort.Ints(horizontalPositions)

	minPos := horizontalPositions[0]
	maxPos := horizontalPositions[len(horizontalPositions)-1]
	var costs []int
	for i := minPos; i < maxPos; i++ {
		cost := 0
		for _, p := range horizontalPositions {
			moves := math.Abs(float64(p - i))
			cost += costOfMoves(moves)
		}
		costs = append(costs, cost)
	}
	sort.Ints(costs)

	fmt.Printf("Answer: %d\n", costs[0])

	return nil
}

func costOfMoves(moves float64) int {
	c := ((moves * moves) + moves) / 2
	return int(c)
}

func init() {
	rootCmd.AddCommand(day7Cmd)
	day7Cmd.AddCommand(day7Part1Cmd)
	day7Cmd.AddCommand(day7Part2Cmd)

	day7Part1Cmd.Flags().StringVarP(&input, "input", "i", "", "solution input data")
	day7Part1Cmd.MarkFlagRequired("input")

	day7Part2Cmd.Flags().StringVarP(&input, "input", "i", "", "solution input data")
	day7Part2Cmd.MarkFlagRequired("input")
}
