package usecase

type TransactionUseCase interface {
	Process(data interface{}) (interface{}, error)
}
