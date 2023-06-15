// root project, ada diluar folder lain2nya
package main

import (
	"fmt"

	"codeid.gofirstpackage/store"
	// bila ingin memakai alias
	// st "codeid.gofirstpackage/store"
)

// kalo mao akses atribut/variablenya, di struct file store harus dimulai dengan huruf Kapital
func main() {
	// contoh alias
	// category := st.NewCategory(1, "Gadget")

	category := store.NewCategory("Gadget")
	category.SetName("Laptop")

	// product := store.Product{Name: "Samsung", Category: "Gadget", Price: 875}

	// memanggil dengan struct category didalam struct product
	product := store.Product{Name: "Samsung", Category: *category, Price: 875}

	product.SetQty(100)

	fmt.Println(category)
	fmt.Println(product)
}
