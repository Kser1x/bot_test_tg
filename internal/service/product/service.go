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

// как то обработаь ошибку
func (s *Service) Get(idx int) (*Product, error) {
	return &allProducts[idx], nil
}
