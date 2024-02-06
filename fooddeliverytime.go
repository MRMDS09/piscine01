package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	if order == "burger" {
		x := food{15}
		return x.preptime
	}
	if order == "chips" {
		x := food{10}
		return x.preptime
	}
	if order == "nuggets" {
		x := food{12}
		return x.preptime
	}
	return 404
}
