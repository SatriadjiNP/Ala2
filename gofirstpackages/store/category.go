package store

import "github.com/google/uuid"

// unexport struct
type Category struct {
	cateId   uuid.UUID
	cateName string
}

// menggunakan pointer bila datanya ingin diupdate langsung
// bila mao merubah unexport menjadi export field, maka buat constructor
func NewCategory(cateName string) *Category {
	return &Category{
		cateId:   uuid.New(),
		cateName: cateName,
	}
}

// method untuk mengakses nama langsung pada struct
func (c *Category) SetName(newName string) {
	c.cateName = newName
}
