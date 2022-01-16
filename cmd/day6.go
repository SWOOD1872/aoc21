/*
Copyright © 2022 Sam Wood <samwooddev@gmail.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// day6Cmd represents the day6 command
var day6Cmd = &cobra.Command{
	Use:   "day6",
	Short: "Advent of Code 2021 day 6 solutions",
}

// day6Part1Cmd represents the part1 command
var day6Part1Cmd = &cobra.Command{
	Use:     "part1",
	Short:   "Advent of Code 2021 day 6 part 1 solution",
	PreRunE: ValidateInput,
	RunE:    day6Part1,
}

// day6Part1 is the part 1 solution code
func day6Part1(cmd *cobra.Command, args []string) error {
	data, err := os.Open(input)
	if err != nil {
		return err
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	var nums []int
	var num int
	for scanner.Scan() {
		line := scanner.Text()
		numsStr := strings.Split(line, ",")
		for _, numStr := range numsStr {
			num, err = strconv.Atoi(numStr)
			if err != nil {
				return err
			}
			nums = append(nums, num)
		}
	}

	var bucket = map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
	}

	for _, n := range nums {
		bucket[n] += 1
	}

	days := 80
	cycle := 7
	newCycle := 9
	for d := 0; d < days; d++ {
		birthingFish := bucket[0]
		for i := 0; i < newCycle-1; i++ {
			bucket[i] = bucket[i+1]
		}
		bucket[cycle-1] += birthingFish
		bucket[newCycle-1] = birthingFish
	}

	var sums int
	for _, n := range bucket {
		sums += n
	}
	fmt.Printf("Answer: %v\n", sums)

	return nil
}

// day6Part2Cmd represents the part2 command
var day6Part2Cmd = &cobra.Command{
	Use:     "part2",
	Short:   "Advent of Code 2021 day 6 part 2 solution",
	PreRunE: ValidateInput,
	RunE:    day6Part2,
}

// day6Part2 is the part 2 solution code
func day6Part2(cmd *cobra.Command, args []string) error {
	data, err := os.Open(input)
	if err != nil {
		return err
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	var nums []int
	var num int
	for scanner.Scan() {
		line := scanner.Text()
		numsStr := strings.Split(line, ",")
		for _, numStr := range numsStr {
			num, err = strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}
	}

	var bucket = map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
	}

	for _, n := range nums {
		bucket[n] += 1
	}

	days := 256
	cycle := 7
	newCycle := 9
	for d := 0; d < days; d++ {
		birthingFish := bucket[0]
		for i := 0; i < newCycle-1; i++ {
			bucket[i] = bucket[i+1]
		}
		bucket[cycle-1] += birthingFish
		bucket[newCycle-1] = birthingFish
	}

	var sums int
	for _, n := range bucket {
		sums += n
	}
	fmt.Printf("Answer: %v\n", sums)

	return nil
}

func init() {
	rootCmd.AddCommand(day6Cmd)
	day6Cmd.AddCommand(day6Part1Cmd)
	day6Cmd.AddCommand(day6Part2Cmd)

	day6Part1Cmd.Flags().StringVarP(&input, "input", "i", "", "solution input data")
	day6Part1Cmd.MarkFlagRequired("input")

	day6Part2Cmd.Flags().StringVarP(&input, "input", "i", "", "solution input data")
	day6Part2Cmd.MarkFlagRequired("input")
}
