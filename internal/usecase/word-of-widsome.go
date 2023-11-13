package usecase

import (
	"context"

	"github.com/dmitrykondrakhin/word-of-wisdom/internal/repository"
)

type WordOfWidsomRepo interface {
	GetQuote(ctx context.Context) (string, error)
}

type WordOfWidsomUseCase struct {
	wordOfWidsomeRepo WordOfWidsomRepo
}

func NewWordOfWidsomUseCase(repos repository.Repositories) *WordOfWidsomUseCase {
	return &WordOfWidsomUseCase{
		wordOfWidsomeRepo: repos.WordOfWidsomRepo,
	}
}

func (w WordOfWidsomUseCase) GetQuote(ctx context.Context) (string, error) {
	return w.wordOfWidsomeRepo.GetQuote(ctx)
}
