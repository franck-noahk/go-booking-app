package booking

type Buyer struct {
	firstName      string
	lastName       string
	ticketsPurched int
}

func (b Buyer) GetUserName() string {
	return b.firstName + " " + b.lastName
}
