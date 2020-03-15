package funding

//Fund is a structure..
type Fund struct {
	balance int64
}

//NewFund is for funding the newone
func NewFund(initialBalance int64) *Fund {
	return &Fund{
		balance: initialBalance,
	}
}

//BalanceAmmount Method to start with reciever
func (f *Fund) BalanceAmmount() int64 {
	return f.balance
}

//Withdraw is used for withdrawing money
func (f *Fund) Withdraw(ammount int64) {
	f.balance = f.balance - ammount
}
