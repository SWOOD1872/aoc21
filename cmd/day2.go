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

type part1Position struct {
	horizontal int
	depth      int
}

func (p *part1Position) move(n int) {
	p.horizontal += n
}

func (p *part1Position) climb(n int) {
	p.depth -= n
}

func (p *part1Position) dive(n int) {
	p.depth += n
}

type part2Position struct {
	horizontal int
	depth      int
	aim        int
}

func (p *part2Position) move(n int) {
	p.horizontal += n
	p.depth += p.aim * n
}

func (p *part2Position) climb(n int) {
	p.aim -= n
}

func (p *part2Position) dive(n int) {
	p.aim += n
}

// day2Cmd represents the day2 command
var day2Cmd = &cobra.Command{
	Use:   "day2",
	Short: "Advent of Code 2021 day 2 solutions",
}

// day2Part1Cmd represents the part1 command
var day2Part1Cmd = &cobra.Command{
	Use:     "part1",
	Short:   "Advent of Code 2021 day 2 part 1 solution",
	PreRunE: ValidateInput,
	RunE:    day2Part1,
}

// day2Part1 is the part 1 solution code
func day2Part1(cmd *cobra.Command, args []string) error {
	data, err := os.Open(input)
	if err != nil {
		return err
	}
	defer data.Close()

	p := part1Position{0, 0}

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		i := line[0]
		v, _ := strconv.Atoi(line[1])

		if strings.ToLower(i) == "forward" {
			p.move(v)
		}
		if strings.ToLower(i) == "down" {
			p.dive(v)
		}
		if strings.ToLower(i) == "up" {
			p.climb(v)
		}
	}

	answer := p.horizontal * p.depth
	fmt.Printf("Answer: %v\n", answer)

	return nil
}

// day2Part2Cmd represents the part2 command
var day2Part2Cmd = &cobra.Command{
	Use:     "part2",
	Short:   "Advent of Code 2021 day 2 part 2 solution",
	PreRunE: ValidateInput,
	RunE:    day2Part2,
}

// day2Part2 is the part 2 solution code
func day2Part2(cmd *cobra.Command, args []string) error {
	data, err := os.Open(input)
	if err != nil {
		return err
	}
	defer data.Close()

	p := part2Position{0, 0, 0}

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		i := line[0]
		v, _ := strconv.Atoi(line[1])

		if strings.ToLower(i) == "forward" {
			p.move(v)
		}
		if strings.ToLower(i) == "down" {
			p.dive(v)
		}
		if strings.ToLower(i) == "up" {
			p.climb(v)
		}
	}

	answer := p.horizontal * p.depth
	fmt.Printf("Answer: %v\n", answer)

	return nil
}

func init() {
	rootCmd.AddCommand(day2Cmd)
	day2Cmd.AddCommand(day2Part1Cmd)
	day2Cmd.AddCommand(day2Part2Cmd)

	day2Part1Cmd.Flags().StringVarP(&input, "input", "i", "", "solution input data")
	day2Part1Cmd.MarkFlagRequired("input")

	day2Part2Cmd.Flags().StringVarP(&input, "input", "i", "", "solution input data")
	day2Part2Cmd.MarkFlagRequired("input")
}
