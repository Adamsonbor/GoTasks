package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"
)

func get_deck(my_cards []string, participant_cards [][]string) map[string]bool {
	var card_names = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	var card_suits = []string{"S", "C", "D", "H"}
	var deck = make(map[string]bool)

	for _, suit := range card_suits {
		for _, name := range card_names {
			deck[name + suit] = true
		}
	}
	for _, card := range my_cards {
		deck[card] = false
	}
	for _, participant := range participant_cards {
		for _, card := range participant {
			deck[card] = false
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

func max_card(cards []string) string {
	var card_weights = map[string]int{
		"2": 0, "3": 1, "4": 2, "5": 3,
		"6": 4, "7": 5, "8": 6, "9": 7,
		"T": 8, "J": 9, "Q": 10, "K": 11, "A": 12}
	var max_card string = cards[0]

	for _, card := range cards {
		if card_weights[card] > card_weights[max_card] {
			max_card = card
		}
	}
	return max_card
}

func is_set(cards []string) bool {
	return cards[0][0] == cards[1][0]
}

func part_1(my_cards []string, participant_cards []string, deck map[string]bool) {
	var card_suits = []string{"S", "C", "D", "H"}

	for _, suit := range card_suits {
		deck[string(participant_cards[0][0]) + suit] = false
	}
}

func part_2(my_cards []string, participant_cards []string, deck map[string]bool) {
}

func part_3(my_cards []string, participant_cards []string, deck map[string]bool) {
}

func part_4(my_cards []string, participant_cards []string, deck map[string]bool) {
}

func part_5(my_cards []string, participant_cards []string, deck map[string]bool) {
}

func part_6(my_cards []string, participant_cards []string, deck map[string]bool) {
}

func part_7(my_cards []string, participant_cards []string, deck map[string]bool) {
}

func part_8(my_cards []string, participant_cards []string, deck map[string]bool) {
}

func part_9(my_cards []string, participant_cards []string, deck map[string]bool) {
}

func part_10(my_cards []string, participant_cards []string, deck map[string]bool) {
}

func part_11(my_cards []string, participant_cards []string, deck map[string]bool) {
}

func how_to_win(my_cards []string, participant_cards [][]string) map[string]bool {
	var card_weights = map[string]int{
		"2": 0, "3": 1, "4": 2, "5": 3,
		"6": 4, "7": 5, "8": 6, "9": 7,
		"T": 8, "J": 9, "Q": 10, "K": 11, "A": 12}
	var deck map[string]bool = get_deck(my_cards, participant_cards)

	if is_set(my_cards) {
		if is_set(participant_cards[0]) {
			if card_weights[my_cards[0]] > card_weights[participant_cards[0][0]] {
				part_1(my_cards, participant_cards[0], deck)
			} else if card_weights[my_cards[0]] < card_weights[participant_cards[0][0]] {
				part_2(my_cards, participant_cards[0], deck)
			} else {
				part_3(my_cards, participant_cards[0], deck)
			}
		} else {
			if card_weights[my_cards[0]] > card_weights[max_card(participant_cards[0])] {
			} else if card_weights[my_cards[0]] < card_weights[max_card(participant_cards[0])] {
				part_4(my_cards, participant_cards[0], deck)
			} else {
				part_5(my_cards, participant_cards[0], deck)
			}
		}
	} else {
		if is_set(participant_cards[0]) {
			if card_weights[max_card(my_cards)] > card_weights[participant_cards[0][0]] {
				part_6(my_cards, participant_cards[0], deck)
			} else if card_weights[max_card(my_cards)] < card_weights[participant_cards[0][0]] {
				part_7(my_cards, participant_cards[0], deck)
			} else {
				part_8(my_cards, participant_cards[0], deck)
			}
		} else {
			if card_weights[max_card(my_cards)] > card_weights[max_card(participant_cards[0])] {
				part_9(my_cards, participant_cards[0], deck)
			} else if card_weights[max_card(my_cards)] < card_weights[max_card(participant_cards[0])] {
				part_10(my_cards, participant_cards[0], deck)
			} else {
				part_11(my_cards, participant_cards[0], deck)
			}
		}
	}
	return deck
}

func main() {
	var (
		scanner *bufio.Scanner
		n_cases int
		n_lines int
		participant_cards [][]string
		my_cards []string
	)

	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n_cases, _ = strconv.Atoi(scanner.Text())
	for i := 0; i < n_cases; i++ {
		scanner.Scan()
		n_lines, _ = strconv.Atoi(scanner.Text())
		lines := make([]string, n_lines)
		for j := 0; j < n_lines; j++ {
			scanner.Scan()
			lines[j] = scanner.Text()
		}
		participant_cards = parse_input(lines)
		my_cards = participant_cards[0]
		participant_cards = participant_cards[1:]
		fmt.Println(how_to_win(my_cards, participant_cards))
	}
}
