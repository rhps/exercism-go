package annalyn

// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
	var canAttack = false
	if !knightIsAwake {
		canAttack = true
	}
	return canAttack
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	var canSpying = false
	if knightIsAwake || archerIsAwake || prisonerIsAwake {
		canSpying = true
	}
	return canSpying
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	var canSignal = false
	if prisonerIsAwake && !archerIsAwake {
		canSignal = true
	}
	return canSignal
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	var freePrisoner = false
	if (prisonerIsAwake && (!knightIsAwake && !archerIsAwake)) || (petDogIsPresent && !archerIsAwake) {
		freePrisoner = true
	}
	return freePrisoner
}
