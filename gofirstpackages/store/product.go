// package ini untuk memudahkan beberapa folder, untuk dicompile dan dishare ke file2 yang lain
// store untuk file yang bertipe domain
package store

import (
	"codeid.gofirstpackage/master"
	"github.com/google/uuid"
)

// type Product struct {
// 	Name, Category string // export field, dikarenakan variable huruf kapital diawal
// 	Price          float64
// 	qty            int // unexport field, dikarenakan variable huruf kecil diawal
// }

// misalkan mau akses category dari file lain
type Product struct {
	id       uuid.UUID
	Name     string // export field, dikarenakan variable huruf kapital diawal
	Price    float64
	qty      int // unexport field, dikarenakan variable huruf kecil diawal
	Category Category
	// supplier ini beda package, harus memakai import
	Supplier master.Supplier
}

type ProductList []Product

// method, kapital diawal untuk export ke main
func (p *Product) SetQty(newQty int) {
	// qty Product struct
	p.qty = newQty
}
