package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"
)

func get_deck() map[string]bool {
	var card_names = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	var card_suits = []string{"S", "C", "D", "H"}
	var deck = make(map[string]bool)

	for _, suit := range card_suits {
		for _, name := range card_names {
			deck[name + suit] = true
		}
	}
	return deck
}

func parse_input(input []string) [][]string {
	var output [][]string

	for _, line := range input {
		output = append(output, strings.Split(line, " "))
	}
	return output
}

func is_set(cards []string) bool {
	return cards[0][0] == cards[1][0]
}

func update_deck(deck map[string]bool, cards []string) {
	var card_names = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	var card_suits = []string{"S", "C", "D", "H"}

	if is_set(cards) {
		for _, card_name := range card_names {
			if string(cards[0][0]) == card_name { break }
			for _, card_suit := range card_suits {
				deck[card_name + card_suit] = false
			}
		}
	} else {
	}
}

func how_to_win(participant_cards [][]string) map[string]bool {
	var deck map[string]bool = get_deck()

	for _, cards := range participant_cards {
		for _, card := range cards {
			deck[card] = false
		}
		update_deck(deck, cards)
	}

	return deck
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n_cases, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n_cases; i++ {
		scanner.Scan()
		n_lines, _ := strconv.Atoi(scanner.Text())
		lines := make([]string, n_lines)
		for j := 0; j < n_lines; j++ {
			scanner.Scan()
			lines[j] = scanner.Text()
		}
		participant_cards := parse_input(lines)
		fmt.Println(how_to_win(participant_cards))
	}
}
