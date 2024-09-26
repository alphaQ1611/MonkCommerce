package utils

func DecreaseByPercentage(original float32, percentage float32) float32 {
	discount := original * (percentage / 100)
	return original - discount
}
