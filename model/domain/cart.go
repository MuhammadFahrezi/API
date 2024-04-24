package domain

type Cart struct {
	Id        int
	UserId    int
	ProductId int
	Product   Product
	Quantity  int
}
