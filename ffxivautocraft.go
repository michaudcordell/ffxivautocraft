package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

/*
Display custom usage information for the command-line interface.
*/
func usage() {
	usageCommandString := "Usage: %s [-h] [-click_delay CLICK_DELAY] " +
		"[-craft_delay CRAFT_DELAY] [-initial_delay INITIAL_DELAY] " +
		"craft_count craft_length macro_key"
	fmt.Printf(usageCommandString+"\n\n", os.Args[0])

	programDescription := "Automate basic crafting task with macro. " +
		"For crafts spanning multiple macros,\n" +
		"I recommend using the /nextmacro command provided by the " +
		"Macro Chain Dalamud plugin."
	fmt.Println(programDescription + "\n")

	fmt.Println("positional arguments:")
	fmt.Println("  craft_count int\n\tthe number of items to craft")
	fmt.Println("  craft_length float\n\tthe duration of the crafting macro")
	fmt.Println("  macro_key string\n\tthe key corresponding to the crafting macro")

	fmt.Println("\noptional arguments:")
	flag.PrintDefaults()
}

func main() {
	// initialize constants
	const (
		PREFIX            = "Progress: "
		DECIMAL_PRECISION = 0
		BAR_LENGTH        = 30
		BAR_CHAR          = "█"
	)

	//
	// setup command-line interface
	//

	// define positional arguments
	var craftCount int
	var craftLength float64
	var macroKey string

	// define optional arguments
	var clickDelay float64
	flag.Float64Var(&clickDelay, "click_delay", 1.0,
		"the time (in seconds) to wait after automated clicking")
	var craftDelay float64
	flag.Float64Var(&craftDelay, "craft_delay", 4.0,
		"the time (in seconds) to ait after finishing an automated crafting macro")
	var initialDelay float64
	flag.Float64Var(&initialDelay, "initial_delay", 10.0,
		"the time (in seconds) before starting the craft automation")

	// enable custom usage printout
	flag.Usage = usage

	// parse args and store positional arguments in their corresponding variables
	flag.Parse()

	craftCount, _ = strconv.Atoi(flag.Arg(0))
	craftLength, _ = strconv.ParseFloat(flag.Arg(1), 64)
	macroKey = flag.Arg(2)

	// setup random number generation
	rand.Seed(time.Now().UnixNano())

	// wait for initial delay
	robotgo.MilliSleep(int(initialDelay * 1000.0))

	// start automated crafting loop
	itemsCrafted := 0
	for itemsCrafted < craftCount {
		// estimate and display remaining craft time
		estimatedTimeRemaining := estimateTime(craftCount, itemsCrafted,
			clickDelay, craftDelay, craftLength)
		fmt.Printf("Estimated time left for craft completion: %dm%ds.",
			estimatedTimeRemaining[0], estimatedTimeRemaining[1])

		// click the button to start crafting
		robotgo.Click()
		robotgo.MilliSleep(int(clickDelay*1000.0 + rand.Float64()*400))

		// start the crafting macro
		robotgo.KeyTap(macroKey)
		robotgo.MilliSleep(int(500.0 + rand.Float64()*200.0))

		// wait until the crafting macro is finished
		robotgo.MilliSleep(int(craftLength*1000.0 + rand.Float64()*2000.0))

		// increment crafted item count and display progress bar
		itemsCrafted++
		SUFFIX := fmt.Sprintf(" | %d/%d Items Crafted", itemsCrafted, craftCount)
		progressBar := createProgressBar(itemsCrafted, craftCount, PREFIX,
			SUFFIX, DECIMAL_PRECISION, BAR_LENGTH, BAR_CHAR)
		fmt.Println(progressBar)

		// wait before starting crafting for next item
		robotgo.MilliSleep(int(craftDelay*1000.0 + rand.Float64()*400))
	}

	fmt.Println("Crafting complete!")

	os.Exit(0)

}

/*
Create a string representing a progress bar given the current progress, the
total progress goal, and other properties of the bar.

A representation of a created bar is as follows:
Prefix |****      | 40.0% Suffix
*/
func createProgressBar(progress int, total int, prefix string, suffix string,
	decimalPrecision int, barLength int, barChar string) string {

	percent := float64(progress) / float64(total)
	barCount := int(percent * float64(barLength))
	barsStr := strings.Repeat(barChar, barCount)
	spacesStr := strings.Repeat(" ", barLength-barCount)
	formattedPercent := strconv.FormatFloat(percent * 100.0, 'f', decimalPrecision, 64)

	progressBarStr := fmt.Sprintf("%s|%s%s| %s%%%s", prefix, barsStr, spacesStr, formattedPercent, suffix)

	return progressBarStr
}

/*
Estimate the time remaining to complete all crafts in minutes and seconds.
*/
func estimateTime(craftCount int, craftProgress int, clickDelay float64,
	craftDelay float64, craftLength float64) [2]int {

	timePerCraft := (0.001 + 0.05) + (0.001 + 1.2*clickDelay) +
		(0.001 + 0.1) + (craftLength + 1.0) + (1.4 * craftDelay)

	craftsRemaining := craftCount - craftProgress

	timeRemaining := float64(craftsRemaining) * timePerCraft

	minutes := int(timeRemaining / 60.0)
	seconds := int(timeRemaining - float64(minutes)*60.0)

	return [2]int{minutes, seconds}
}
