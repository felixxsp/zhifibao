package transaction

import "zhifubao/person"

type RealUsecase struct {
	transactionData *Database
	personData      *person.Database
}

type Usecase interface {
	IsActive()
	NewTransaction()
	ViewTransaction()
	ViewAll()
}

func NewUsecase(datalayer *person.Database, self *Database) *RealUsecase {
	return &RealUsecase{
		personData:      datalayer,
		transactionData: self,
	}
}
func (db *RealUsecase) IsActive() {

}

func (db *RealUsecase) NewTransaction() {

}

func (db *RealUsecase) ViewTransaction() {

}

func (db *RealUsecase) ViewAll() {

}
