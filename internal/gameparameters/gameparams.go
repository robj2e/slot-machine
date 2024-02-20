package gameparams

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"slot-machine/internal/helpers"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// GameParams struct defines game parameters passed in through starting the app + user input post startup
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

// NewGameParams func parses info based on startup flags to populate and return struct
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

// Begin method prompts the user for input for their bet amount
func (gp *GameParams) Begin() {
	for {
		betAmt, err := helpers.StringPromptIntReturn("Enter Betting Amount")
		if err != nil {
			fmt.Println(err)
		} else {
			gp.betAmount = betAmt
			break
		}
	}
}

// Spin method "spins" the reels and places the results within the struct for later analysis
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

// OutputReels method outputs the resulting reel values/positions to the user
func (gp *GameParams) OutputReels() {
	fmt.Println("------------")

	topRow := strings.Join([]string{
		gp.reelsWithValues[1][helpers.WinLineMinusOne(gp.spinResultPosition[0], gp.reelSymbolsLength)],
		gp.reelsWithValues[1][helpers.WinLineMinusOne(gp.spinResultPosition[1], gp.reelSymbolsLength)],
		gp.reelsWithValues[1][helpers.WinLineMinusOne(gp.spinResultPosition[2], gp.reelSymbolsLength)],
		gp.reelsWithValues[1][helpers.WinLineMinusOne(gp.spinResultPosition[3], gp.reelSymbolsLength)],
	}, " ")

	fmt.Println(topRow)

	winRow := strings.Join([]string{
		gp.spinResult[0],
		gp.spinResult[1],
		gp.spinResult[2],
		gp.spinResult[3],
	}, " ")

	fmt.Println(winRow)

	bottomRow := strings.Join([]string{
		gp.reelsWithValues[1][helpers.WinLinePlusOne(gp.spinResultPosition[0], gp.reelSymbolsLength)],
		gp.reelsWithValues[1][helpers.WinLinePlusOne(gp.spinResultPosition[1], gp.reelSymbolsLength)],
		gp.reelsWithValues[1][helpers.WinLinePlusOne(gp.spinResultPosition[2], gp.reelSymbolsLength)],
		gp.reelsWithValues[1][helpers.WinLinePlusOne(gp.spinResultPosition[3], gp.reelSymbolsLength)],
	}, " ")

	fmt.Println(bottomRow)
	fmt.Println("------------")
}

// DetermineOutcome method does analysis on the reel values and determines whether the user has a winning spin and calculates the resulting amounts and outputs hose to the console.
func (gp *GameParams) DetermineOutcome() {
	sortedSlice := make([]string, 4)
	copy(sortedSlice, gp.spinResult)
	sort.Strings(sortedSlice)

	for i := 0; i < len(sortedSlice); i++ {
		if sortedSlice[i] == "X" {
			sortedSlice[i] = sortedSlice[0]
		}
	}

	win, val := helpers.ArrayAllSameValue(sortedSlice)

	if win {
		gp.winner = true
		gp.winningLetter = val
		multiplier := gp.payTable[gp.winningLetter]
		multiplierOutput := fmt.Sprintf("Multiplier: %d", multiplier)
		fmt.Println(multiplierOutput)
		amountWon := fmt.Sprintf("Amount Won: %d\n", gp.betAmount*multiplier)
		fmt.Println(amountWon)
	} else {
		gp.winner = false
		gp.winningLetter = val
		fmt.Println("Multiplier: 0")
		fmt.Println("Amount Won: 0")
	}
}

// Cleanup method resets some of the values within the struct for subsequent spins to have a blank canvas to enter values into.
func (gp *GameParams) Cleanup() {
	gp.betAmount = 0
	gp.spinResult = []string{}
	gp.spinResultPosition = []int{}
	gp.winner = false
	gp.winningLetter = ""

	fmt.Println("Starting fresh round...")
}
