package toko

import "github.com/pushm0v/gounittest/service"

type tokoStorage struct {
	storage service.StorageService
}

type TokoStorage interface {
	UpdateStok(nama string, qty int)
	IsStokAvailable(nama string, qty int) bool
	GetStok(nama string) int
}

func NewTokoStorage() TokoStorage {
	return &tokoStorage{
		storage: service.NewStorage(),
	}
}

func (ts *tokoStorage) UpdateStok(nama string, qty int) {
	ts.storage.Save(nama, qty)
}

func (ts *tokoStorage) IsStokAvailable(nama string, qty int) bool {
	stok, ok := ts.storage.Fetch(nama).(int)

	if ok {
		return stok >= qty
	}

	return false
}

func (ts *tokoStorage) GetStok(nama string) int {
	stok, _ := ts.storage.Fetch(nama).(int)

	return stok
}
