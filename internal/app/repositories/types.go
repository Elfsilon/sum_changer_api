package repos

type AccountRepository interface {
	Get() (float32, error)
	TopUp(sum float32) error
	Withdraw(sum float32) error
}
