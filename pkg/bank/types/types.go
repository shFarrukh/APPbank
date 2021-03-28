package types 
//money is kept in minimum currencies (cent kopeck diram) 
type Money int64
//payment categories (foods , clothes...)
type Category string
// Paymant information
type Payment struct{
	ID int
	Amount Money
	Category Category
}