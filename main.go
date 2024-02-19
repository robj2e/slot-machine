package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type GameParams struct {
	reelCount          int
	reelSymbolsLength  int
	reelsWithValues    [][]string
	payTable           map[string]int
	betAmount          int
	spinResult         []string
	spinResultPosition []int
	winner             bool
	winningLetter      string
}

func main() {
	gameParams := NewGameParams()
	for {
		gameParams.Begin()
		gameParams.Spin()
		gameParams.OutputReels()
		gameParams.DetermineOutcome()
		gameParams.Cleanup()
	}
}

func (gp *GameParams) Cleanup() {
	gp.betAmount = 0
	gp.spinResult = []string{}
	gp.spinResultPosition = []int{}
	gp.winner = false
	gp.winningLetter = ""

	fmt.Println("Starting fresh round...")
}

func (gp *GameParams) Begin() {
	for {
		betAmt, err := StringPromptIntReturn("Enter Betting Amount")
		if err != nil {
			fmt.Println(err)
		} else {
			gp.betAmount = betAmt
			break
		}
	}
}

func (gp *GameParams) DetermineOutcome() {
	sortedSlice := make([]string, 4)
	copy(sortedSlice, gp.spinResult)
	sort.Strings(sortedSlice)

	for i := 0; i < len(sortedSlice); i++ {
		if sortedSlice[i] == "X" {
			sortedSlice[i] = sortedSlice[0]
		}
	}

	win, val := arrayAllSameValue(gp, sortedSlice)
	if win {
		gp.winner = true
		gp.winningLetter = val
	} else {
		gp.winner = false
		gp.winningLetter = val
	}

	if gp.winner {
		multiplier := gp.payTable[gp.winningLetter]
		multiplierOutput := fmt.Sprintf("Multiplier: %d", multiplier)
		fmt.Println(multiplierOutput)
		amountWon := fmt.Sprintf("Amount Won: %d\n\n", gp.betAmount*multiplier)
		fmt.Println(amountWon)

	} else {
		fmt.Println("Multiplier: 0")
		fmt.Println("Amount Won: 0")
	}

}

func arrayAllSameValue(gp *GameParams, sortedStrs []string) (b bool, val string) {
	for i := 0; i < len(sortedStrs); i++ {
		if sortedStrs[i] != sortedStrs[0] {
			return false, "No Win"
		}
	}
	return true, sortedStrs[0]
}

func (gp *GameParams) OutputReels() {
	fmt.Println("------------")
	topRow := fmt.Sprintf("%s %s %s %s",
		gp.reelsWithValues[1][WinLineMinusOne(gp.spinResultPosition[0])],
		gp.reelsWithValues[1][WinLineMinusOne(gp.spinResultPosition[1])],
		gp.reelsWithValues[1][WinLineMinusOne(gp.spinResultPosition[2])],
		gp.reelsWithValues[1][WinLineMinusOne(gp.spinResultPosition[3])])

	fmt.Println(topRow)

	winRow := fmt.Sprintf("%s %s %s %s",
		gp.spinResult[0],
		gp.spinResult[1],
		gp.spinResult[2],
		gp.spinResult[3])

	fmt.Println(winRow)

	bottomRow := fmt.Sprintf("%s %s %s %s",
		gp.reelsWithValues[1][WinLinePlusOne(gp.spinResultPosition[0])],
		gp.reelsWithValues[1][WinLinePlusOne(gp.spinResultPosition[1])],
		gp.reelsWithValues[1][WinLinePlusOne(gp.spinResultPosition[2])],
		gp.reelsWithValues[1][WinLinePlusOne(gp.spinResultPosition[3])])

	fmt.Println(bottomRow)
	fmt.Println("------------")
}

func WinLineMinusOne(i int) int {
	if i-1 < 0 {
		return 5
	} else {
		return i - 1
	}
}

func WinLinePlusOne(i int) int {
	if i+1 > 5 {
		return 0
	} else {
		return i + 1
	}
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
		return -1, errors.New("Error parsing input value, retrying...")
	}

	return val, nil
}

func NewGameParams() *GameParams {
	reelCount := flag.Int("reelCount", 4, "Amount of reels in this slot machine")
	reelSymbolLength := flag.Int("reelSymbolLength", 6, "Amount of values in the reel itself")

	flag.Parse()
	reelVals := flag.Args()

	reelSymbols := reelVals[:*reelSymbolLength]

	reelMultipliers := reelVals[*reelSymbolLength:]

	reelsWithValues := make([][]string, *reelCount)

	for i := 0; i < *reelCount; i++ {
		reelsWithValues[i] = reelSymbols
	}

	multiplierMap := make(map[string]int)

	for i := 0; i < len(reelMultipliers); i++ {
		multiplierConverted, err := strconv.Atoi(reelMultipliers[i])
		if err != nil {
			fmt.Println("Startup parameters incorrect, closing program")
			os.Exit(1)
		}

		multiplierMap[reelSymbols[i]] = multiplierConverted
	}

	gp := &GameParams{
		reelCount:         *reelCount,
		reelSymbolsLength: len(reelSymbols),
		reelsWithValues:   reelsWithValues,
		payTable:          multiplierMap,
		winner:            false,
	}

	return gp
}

func (gp *GameParams) Spin() {
	spinResultSlice := make([]string, gp.reelCount)
	spinResultPosition := make([]int, gp.reelCount)

	var wg sync.WaitGroup
	wg.Add(gp.reelCount)

	for i := 0; i < gp.reelCount; i++ {
		go func(spinResSlice []string, spinResPosSlice []int, reel []string, index int) {
			defer wg.Done()
			position := rand.IntN(gp.reelSymbolsLength)
			spinResSlice[index] = reel[position]
			spinResPosSlice[index] = position
		}(spinResultSlice, spinResultPosition, gp.reelsWithValues[i], i)
	}

	wg.Wait()

	gp.spinResult = spinResultSlice
	gp.spinResultPosition = spinResultPosition
}
