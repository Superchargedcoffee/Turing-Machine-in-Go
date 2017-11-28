package main

import ("fmt"
	"math/rand"
	"time")

var (
	cards = make(map[int][]int)
	tape []int
	)

func makeTape(length int, random bool) {
	if random {
		for i := 0; i < length; i++ {
			tape = append(tape, rand.Int() % 2)
		}
	} else {
		for i := 0; i < length; i++ {
			tape = append(tape, 0)
		}
	}
}

func printTape(headPos, cardNo int) {		 // "-" for 0, "#" for 1
	out, card := "", cards[cardNo]		 // "^" if the machine is increasing a cell
	for i := range tape {			 // "v" if it's decreasing it
		if i == headPos {		 // "=" if it does nothing
			if tape[i] == 0 {
				if card[0] == 1 {
					out += "^"
				} else { out += "=" }
			} else {
				if card[3] == 0 {
					out += "v"
				} else { out += "=" }
			}
		} else {
			if tape[i] == 0 {
				out += "-"
			} else { out += "#" }
		}
	}
	fmt.Println(out)
}

func step(cardNo, Pos int) (nextCard, nextPos int) {
	card := cards[cardNo]
	if tape[Pos] == 0 {
		tape[Pos] = card[0]
		nextPos = Pos + card[1]
		nextCard = card[2]
	} else {
		tape[Pos] = card[3]
		nextPos = Pos + card[4]
		nextCard = card[5]
	}
	return nextCard, nextPos
}

func getBoolIn() bool {
	var input string
	fmt.Scanln(&input)
	if input == "y" || input == "Y" {
		return true
	} else if input == "n" || input == "N" {
		return false
	}
	fmt.Println("Invalid input, please enter one of Y/y or N/n")
	return getBoolIn()
}

func randCard(n, seed int) map[int][]int {		// Returns a map of n random instruction cards
	rand.Seed(int64(seed))
	var cards = make(map[int][]int)
	for i := 1; i <= n; i++ {
		var card_i []int
		for j := 0; j < 2; j++ {
			card_i = append(card_i, rand.Int() % 2)
			card_i = append(card_i, (rand.Int() % 3) - 1)
			card_i = append(card_i, rand.Int() % (n + 1))
		}
		cards[i] = card_i
	}
	return cards
}

func getIntIn() int {
        var input string
        fmt.Scanln(&input)
        if len(input) > 5 {
                fmt.Println("Input too large, please enter a number between 0 and 99999")
                return getIntIn()
        }
        for i := len(input); i < 5; i++ {
                input = "0" + input
        }
        input = "id:" + input
        var num int
        if _, err := fmt.Sscanf(input, "id:%5d", &num); err == nil {
                return num
        }
        fmt.Println("Error: could not interpret input")
        return getIntIn()
}

func listSum(array []int) int {
	sum := 0
	for i := range array {
		sum += array[i]
	}
	return sum
}

func main() {
	fmt.Println("Use looping tape? [y/n]")
	looping := getBoolIn()
	fmt.Println("Use random instruction set? [y/n]")
	if getBoolIn() {
		fmt.Println("How many instruction cards do you want?")
		cardnum := getIntIn()
		fmt.Println("Input a seed to use")
		seed := getIntIn()
		cards = randCard(cardnum, seed)
	} else {
		card1, card2, card3, card4, card5 := []int{1, 1, 2, 1, -1, 3}, []int{1, 1, 3, 1, 1, 2}, []int{1, 1, 4, 0, -1, 5}, []int{1, -1, 1, 1, -1, 4}, []int{1, 1, 0, 0, -1, 1}
		/*
		First three entries indicate action to take on 0: What to print to the tape, how to move, and what card to go to
		Second three indicate the same for 1
		*/
		cards[1], cards[2], cards[3], cards[4], cards[5] = card1, card2, card3, card4, card5
	}
	fmt.Println("Use random tape? [y/n] (Default is blank)")
	makeTape(104, getBoolIn())
	fmt.Println("Enter a failsafe (will halt after this many steps)")
	pos, card, steps, safety := len(tape)/2, 1, 0, getIntIn()
	if !looping {
		for card != 0 && pos != -1 && pos != len(tape) && steps < safety {
			printTape(pos, card)
			card, pos = step(card, pos)
			steps += 1
			time.Sleep(time.Second/100)
		}
	} else {
		for card != 0 && steps < safety {
			printTape(pos, card)
			card, pos = step(card, pos)
			if pos < 0 {
				pos = len(tape) - 1
			} else { pos = pos % len(tape) }
			steps += 1
			time.Sleep(time.Second/100)
		}
	}
	if card == 0 {
		fmt.Println("Halted")
		fmt.Println("Printed", listSum(tape), "ones in", steps, "steps")
	} else if steps == safety {
		fmt.Println("Safety limit reached, halted")
	} else {
		fmt.Println("Out of tape")
		fmt.Println("Took", steps, "steps")
	}
}
