package product

// содежит структуру продукта
var allProducts = []Product{
	{Title: "one"},
	{Title: "two"},
	{Title: "three"},
	{Title: "four"},
	{Title: "five"},
}

type Product struct {
	Title string
}
