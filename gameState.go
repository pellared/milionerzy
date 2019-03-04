package main

// NewGameState initialized new game
func NewGameState() *GameState {
	return &GameState{
		currentMoney:     0,
		guaranteedMoney:  0,
		canAskAudience:   true,
		canAskFrind:      true,
		canUseHalfByHalf: true,
	}
}

// GameState represent the current state of the game
type GameState struct {
	currentMoney     int
	guaranteedMoney  int
	canAskAudience   bool
	canUseHalfByHalf bool
	canAskFrind      bool
}
