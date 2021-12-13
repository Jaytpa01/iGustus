package service

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Jaytpa01/iGustus/internal/entities"
	"github.com/Jaytpa01/iGustus/pkg/logger"
	"go.uber.org/zap"
)

type mudService struct{}

func NewMudService() entities.MudService {
	return &mudService{}
}

func (m *mudService) Roll(req entities.RollRequest) ([]entities.DiceRollResult, error) {
	r, err := regexp.Compile(`\d+d\d+`)
	if err != nil {
		logger.Log.Error("error compiling regex", zap.Error(err))
		return nil, err
	}

	results := []entities.DiceRollResult{}

	rolls := r.FindAllString(req.Content, -1)
	if len(rolls) == 0 {
		return nil, fmt.Errorf("no valid rolls were provided")
	}

	for _, roll := range rolls {
		split := strings.Split(roll, "d")
		if len(split) != 2 {
			continue
		}

		numOfDice, _ := strconv.Atoi(split[0])
		faces, _ := strconv.Atoi(split[1])
		result := rollDice(numOfDice, faces)
		result.Faces = faces

		results = append(results, result)
	}

	return results, nil
}

func rollDice(numOfDice, numOfFaces int) entities.DiceRollResult {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	results := make([]int, numOfDice)
	for i := 0; i < numOfDice; i++ {
		results[i] = r.Intn(numOfFaces) + 1
	}

	total := 0
	for _, num := range results {
		total += num
	}

	return entities.DiceRollResult{
		Results: results,
		Total:   total,
	}
}
