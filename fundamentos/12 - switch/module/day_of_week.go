package module

func DayOfWeek(day string) int {
	// fmt.Printf("%T - %q\n\n", day, day)
	switch string(day) {
	case "Sunday":
		return 1 
	case "Monday": 
		return 2
	case "Tuesday":
		return 3
	case "Wednesday": 
		return 4
	case "Thursday": 
		return 5
	case "Friday": 
		return 6
	case "Saturday":
		return 7
	default:
		return -1
	}
}