package helpers

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function WinLineMinusOne is used to determine slice position for the row above the winning row.
func WinLineMinusOne(i, reelLength int) int {
	if i-1 < 0 {
		return reelLength - 1
	} else {
		return i - 1
	}
}

// Function WinLineplusOne is used to determine slice position for the row below the winning row.
func WinLinePlusOne(i, reelLength int) int {
	if i+1 > reelLength-1 {
		return 0
	} else {
		return i + 1
	}
}

// Function ArrayAllSameValue checks for winning reels position
func ArrayAllSameValue(sortedStrs []string) (bool, string) {
	for i := 0; i < len(sortedStrs); i++ {
		if sortedStrs[i] != sortedStrs[0] {
			return false, "No Win"
		}
	}

	if sortedStrs[0] == "X" {
		return false, "No Win"
	}

	return true, sortedStrs[0]
}

// Function StringPromptIntReturn prompts the user for their bet amount, takes it as a string and converts it to an int. Performs some basic validation too
func StringPromptIntReturn(str string) (int, error) {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, str+" \n")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}

	val, err := strconv.Atoi(strings.TrimSpace(s))
	if val <= 0 || err != nil {
		err := errors.New("error parsing input value, please retry")
		return -1, err
	}

	return val, nil
}
