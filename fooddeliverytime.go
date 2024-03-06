package piscine

import "fmt"

type food struct {
	preptime int
}

var menu = map[string]food{
	"burger":  {preptime: 15},
	"chips":   {preptime: 10},
	"nuggets": {preptime: 12},
}

func FoodDeliveryTime(order string) (int, error) {
	item, exists := menu[order]
	if !exists {
		return 0, fmt.Errorf("404: Item not found in the menu")
	}

	return item.preptime, nil
}
