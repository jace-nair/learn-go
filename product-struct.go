package main

// Product struct
type Product struct {
	Product_ID   uint64
	Product_Name string
	Price        uint64
	Rating       uint8
	Image        string
	Specs
}

// Create new product
func createProduct(productID uint64, productName string, price uint64, rating uint8, image string, specs Specs) Product {
	//create an instance of struct using name, empty map and zero tip and returns the newly created myStruct
	p := Product{
		Product_ID:   productID,
		Product_Name: productName,
		Price:        price,
		Rating:       rating,
		Image:        image,
		Specs:        specs,
	}
	//return the new user
	return p
}

// Specs struct used as part of Product struct
type Specs struct {
	Length   uint64
	Width    uint64
	Height   uint64
	Colour   string
	Material string
}

// Create Specs for new Product
func createSpecs(length uint64, width uint64, height uint64, colour string, material string) Specs {
	s := Specs{
		Length:   length,
		Width:    width,
		Height:   height,
		Colour:   colour,
		Material: material,
	}
	//return the specs
	return s
}
