package utility

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
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
		if strings.Contains(v, "d100") {
			// handle percentile rolls
			roll := strings.Split(v, "d")
			num := 0
			if roll[0] == "" {
				num = 1
			} else {
				numb, err := strconv.Atoi(roll[0])
				if err != nil {
					return 0, nil, fmt.Errorf("error converting number of dice to int: %v", err)
				}
				num = numb
			}
			// roll d100
			for range num {
				roll := (rand.Intn(10) + 1) * 10
				rolls = append(rolls, Roll{100, roll})
				total += roll
			}
			// roll d10
			if total < 100 {
				for range num {
					roll := (rand.Intn(10) + 1)
					rolls = append(rolls, Roll{10, roll})
					total += roll
				}
			}

		} else if strings.Contains(v, "d") {
			// split roll find type and number of dice
			roll := strings.Split(v, "d")
			num := 0
			if roll[0] == "" {
				num = 1
			} else {
				numb, err := strconv.Atoi(roll[0])
				if err != nil {
					return 0, nil, fmt.Errorf("error converting number of dice to int: %v", err)
				}
				num = numb
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
			trim := strings.TrimSpace(v)
			bonus, err := strconv.Atoi(trim)
			if err != nil {
				fmt.Println("Error converting bonus to int", err)
			}
			total += bonus
		}
	}
	return total, rolls, nil
}
