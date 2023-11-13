package usecase

import "github.com/dmitrykondrakhin/word-of-wisdom/internal/repository"

type Usecases struct {
	WordOfWidsomUseCase *WordOfWidsomUseCase
}

func CreateUsecases(repos repository.Repositories) Usecases {
	return Usecases{
		WordOfWidsomUseCase: NewWordOfWidsomUseCase(repos),
	}
}
