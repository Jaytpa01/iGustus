package mud

import "github.com/Jaytpa01/iGustus/internal/entities"

type DiceRollResult struct {
	Faces   int
	Results []int
	Total   int
}

type MudService interface {
	Roll(entities.RollRequest) ([]DiceRollResult, error)
}
