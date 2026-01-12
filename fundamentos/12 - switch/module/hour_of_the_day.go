package module

func HourOfTheDay(hour int) string {
	switch {
	case hour >= 5 && hour < 12:
		return "morning"
	case hour >= 12 && hour < 18:
		return "afternoon"
	case hour >= 18 && hour < 24:
		return "night"
	case hour >= 0 && hour < 5:
		return "early morning"
	default:
		return "value invalid"
	}
}