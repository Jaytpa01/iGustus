package entities

type DiceRollResult struct {
	Faces   int
	Results []int
	Total   int
}

type MudService interface {
	Roll(RollRequest) ([]DiceRollResult, error)
}
