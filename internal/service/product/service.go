package product

// выводит все продукты
type Service struct {
}

func NewService() *Service {

	return &Service{}
}
func (*Service) List() []Product {
	return allProducts
}
