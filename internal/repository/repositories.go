package repository

type Repositories struct {
	WordOfWidsomRepo *WordOfWidsomeRepo
}

func CreateRepositories() Repositories {
	return Repositories{
		WordOfWidsomRepo: NewWordOfWidsomeRepo(),
	}
}
