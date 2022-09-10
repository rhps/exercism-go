package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	tot := 0
	for i := 0; i < len(birdsPerDay); i++ {
		tot = tot + birdsPerDay[i]
	}

	return tot
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	tot := 0
	startdate := 0
	if week > 1 {
		startdate = week*7 - 7
	}
	enddate := startdate + 7
	for i := startdate; i < enddate; i++ {
		tot = tot + birdsPerDay[i]
	}

	return tot
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i] = birdsPerDay[i] + 1
		}
	}
	return birdsPerDay
}
