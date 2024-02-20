package helpers

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func WinLineMinusOne(i, reelLength int) int {
	if i-1 < 0 {
		return reelLength - 1
	} else {
		return i - 1
	}
}

func WinLinePlusOne(i, reelLength int) int {
	if i+1 > reelLength-1 {
		return 0
	} else {
		return i + 1
	}
}

func ArrayAllSameValue(sortedStrs []string) (b bool, val string) {
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
	if err != nil {
		return -1, errors.New("Error parsing input value, retry...")
	}

	return val, nil
}
