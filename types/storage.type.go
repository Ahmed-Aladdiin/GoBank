package types

type Storage interface {
	CreateAccount(*Account) (*Account, error)
	DeleteAccount(int) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
	UpdateAccount(*Account) error
	Deposit(int, float64) error
	Withdraw(int, float64) error
	Transfare(int, int, float64) error
}
