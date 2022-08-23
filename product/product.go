package product

type Product struct {
	Id          int    `json:"id" xml:"id"`
	Mfr         string `json:"mfr" xml:"mfr"`
	Name        string `json:"name" xml:"name"`
	Description string `json:"description" xml:"description"`
}

func ListProducts() []Product {
	p := []Product{
		{
			Id:          1,
			Mfr:         "Megger",
			Name:        "SMRT-36",
			Description: "Relay Test Set",
		},
		{
			Id:          2,
			Mfr:         "Apple",
			Name:        "iPhone14",
			Description: "Smartphone",
		},
		{
			Id:          3,
			Mfr:         "Lenovo",
			Name:        "AMD-A6",
			Description: "Laptop",
		},
	}
	return p
}
