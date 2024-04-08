package utility

import (
	"strings"
	"strconv"
	"math/rand"
	"fmt"
)

type Roll struct {
	dice int
	roll int
}

func ExpressionRoll(exp string) (int, []Roll, error) {

		total := 0
		rolls := make([]Roll, 0)

	initialSplit := strings.Split(exp, "+")
		for _, v := range initialSplit {
			if strings.Contains(v, "d") {
				// split roll find type and number of dice
				roll := strings.Split(v, "d")
				num, err := strconv.Atoi(roll[0])
				if err != nil {
					return 0, nil, fmt.Errorf("error converting number of dice to int: %v", err)
				}
				sides, err := strconv.Atoi(roll[1])
				if err != nil {
					return 0, nil, fmt.Errorf("error converting sides to int: %v", err)
				}
				// roll dice and add value to total
				for range num {
					roll := rand.Intn(sides) + 1
					rolls = append(rolls, Roll{sides, roll})
					total += roll
				}
			} else {
				// add bonuses to roll total
				bonus, err := strconv.Atoi(v)
				if err != nil {fmt.Println("Error converting bonus to int", err)}
				total += bonus
			}
		}
	return total, rolls, nil
}