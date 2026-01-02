package main

import (
	"bufio"
    "fmt"
    "os"
	"strconv"
	"math"
)

func day1_part1() (int, error) {
    file, err := os.Open("inputs/day1.txt")
    if err != nil {
		return 0, fmt.Errorf("error opening file")
	}
	defer file.Close()

	position := 50
	counter := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value := scanner.Text()
		if string(value[0]) == "L" {
			num, err := strconv.Atoi(value[1:])
			if err != nil {
				return 0, fmt.Errorf("error reading in direction")
			}
			position -= num
		} else {
			num, err := strconv.Atoi(value[1:])
			if err != nil {
				return 0, fmt.Errorf("error reading in direction")
			}
			position += num
		}
		if position%100 == 0 {
			counter += 1
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error scanning in file")
	}

	return counter, nil
}

func day1_part2() (int, error) {
    file, err := os.Open("inputs/day1.txt")
    if err != nil {
		return 0, fmt.Errorf("error opening file")
	}
	defer file.Close()

	prevPosition := 50
	position := 50
	counter := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value := scanner.Text()
		num, err := strconv.Atoi(value[1:])
		if err != nil {
			return 0, fmt.Errorf("error reading in direction")
		}
		if num > 99 {
			counter += int(math.Floor(float64(num/100)))
			num = num%100
		}
		prevPosition = position
		if string(value[0]) == "L" {
			position -= num
			if position == 0 {
				counter += 1
			} else if position < 0 {
				if prevPosition != 0 {
					counter += 1
				}
				position += 100
			}
		} else {
			position += num
			if position > 99 {
				if prevPosition != 0 {
					counter += 1
				}
				position -= 100
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error scanning in file")
	}

	return counter, nil

}

func main() {
	fmt.Println(day1_part2())
}