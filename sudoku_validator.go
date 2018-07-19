package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func isValidSudoku(grid string) (bool, error) {
	// Base case. Check if input grid is empty
	if grid == "" || len(grid) == 0 {
		return false, errors.New("Invalid grid. Input is empty")
	}

	// Maintain a hashset to detect duplicate number within a row/column/block
	// In Go, there is no 'hashset' type and hence we use hashmap[key, value] with value = true
	// to simulate a hashset
	set := make(map[string]bool)

	// split the input string by ','. s is an array of strings
	s := strings.Split(grid, ",")

	// check if we have more than 9 rows in input. If yes, return false
	if len(s) != 9 {
		return false, errors.New("Invalid input grid. Number of rows is not equal to 9")
	}
	for i := 0; i < 9; i++ {
		// check if we have more than 9 columns in each row. If yes, return false
		if len(s[i]) != 9 {
			return false, errors.New("Invalid input grid. Number of columns is not equal to 9")
		}
		for j, c := range s[i] {
			// We can't use '+' to concatenate strings in Go.
			// Hence using strings.Builder to append strings
			var setKey strings.Builder

			// One number can be present only once in a row. We are constructing a string
			// of this format "3 in row 1" as key and inserting into our set.
			setKey.WriteString(string(c))
			setKey.WriteString(" in row ")
			setKey.WriteString(strconv.Itoa(i))
			// Convert StringBuilder to String type
			k := setKey.String()
			// Get from set before inserting
			_, ok := set[k]
			// If get is successful, we already encountered the same number and it's a duplicate
			if ok {
				return false, errors.New("Invalid input grid. Duplicate number in row")
			}
			set[k] = true
			setKey.Reset()

			// One number can be present only once in a column. We are constructing a string
			// of this format "6 in column 1" as key and inserting into our set.
			setKey.WriteString(string(c))
			setKey.WriteString(" in column ")
			setKey.WriteString(strconv.Itoa(j))
			k = setKey.String()
			_, ok = set[k]
			if ok {
				return false, errors.New("Invalid input grid. Duplicate number in column")
			}
			set[k] = true
			setKey.Reset()

			// One number can be present only once in a block. We are constructing a string
			// of this format "6 in block 1 - 1" as key and inserting into our set
			setKey.WriteString(string(c))
			setKey.WriteString(" in block ")
			num1 := i / 3
			num2 := j / 3
			setKey.WriteString(strconv.Itoa(num1))
			setKey.WriteString(" - ")
			setKey.WriteString(strconv.Itoa(num2))
			k = setKey.String()
			_, ok = set[k]
			if ok {
				return false, errors.New("Invalid input grid. Duplicate number in block")
			}
		}
	}
	return true, nil
}

// Tests for isValidSudoku function
func main() {
	// Assumption: Input is comma seperated rows

	// Empty input
	input1 := ""
	fmt.Printf("input1 is '%s'\n", input1)
	result1, err := isValidSudoku(input1)
	fmt.Print(result1)
	fmt.Print(". ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	// Invalid number of rows test. Input has 10 rows.
	input2 := "759462813,463518297,128973645,697184532,584236971,231795486,345821769,912647358,876359124,123456789"
	fmt.Printf("input2 is '%s'\n", input2)
	result2, err := isValidSudoku(input2)
	fmt.Print(result2)
	fmt.Print(". ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	// Invalid number of columns test. First row has 10 column values
	input3 := "7594628135,463518297,128973645,697184532,584236971,231795486,345821769,912647358,876359124"
	fmt.Printf("input3 is '%s'\n", input3)
	result3, err := isValidSudoku(input3)
	fmt.Print(result3)
	fmt.Print(". ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	// Duplicate number in row test. Second row has 2 entries of 1
	input4 := "835416927,196857431,417293658,569134782,123678549,748529163,652781394,981345276,374962815"
	fmt.Printf("input4 is '%s'\n", input4)
	result4, err := isValidSudoku(input4)
	fmt.Print(result4)
	fmt.Print(". ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	// Duplicate number in column test. First and third row has same first column value
	input5 := "835416927,296857431,817293658,569134782,123678549,748529163,652781394,981345276,374962815"
	fmt.Printf("input5 is '%s'\n", input5)
	result5, err := isValidSudoku(input5)
	fmt.Print(result5)
	fmt.Print(". ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	// Duplicate number in block test. First block has 9 in both second and third row
	input6 := "835416927,296857431,419293658,569134782,123678549,748529163,652781394,981345276,374962815"
	fmt.Printf("input6 is '%s'\n", input6)
	result6, err := isValidSudoku(input6)
	fmt.Print(result6)
	fmt.Print(". ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	// valid sudoku
	input7 := "835416927,296857431,417293658,569134782,123678549,748529163,652781394,981345276,374962815"
	fmt.Printf("input7 is '%s'\n", input7)
	result7, err := isValidSudoku(input7)
	fmt.Print(result7)
	fmt.Print(". ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
}
