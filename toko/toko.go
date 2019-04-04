package toko

import "github.com/pushm0v/gounittest/service"

type Toko interface {
	Buka()
	IsBuka() bool
	Tutup()
	TambahStokMinuman(minuman Minuman, qty int)
	TotalStokMinuman(minuman Minuman) int
	JualMinuman(minuman Minuman, qty int) bool
}

type toko struct {
	isBuka       bool
	kulkas       TokoStorage
	dummyService service.DummyService
}

func NewToko(ts TokoStorage) Toko {
	ds := service.NewDummyService()
	toko := &toko{
		kulkas:       ts,
		dummyService: ds,
	}

	return toko
}

func (t *toko) Buka() {
	t.isBuka = true
}

func (t *toko) IsBuka() bool {
	return t.isBuka
}

func (t *toko) Tutup() {
	t.isBuka = false
}

func (t *toko) JualMinuman(minuman Minuman, qty int) bool {
	if t.kulkas.IsStokAvailable(minuman.Nama, qty) {
		t.kurangStokMinuman(minuman, qty)
		return true
	}

	return false
}

func (t *toko) TambahStokMinuman(minuman Minuman, qty int) {
	namaMinuman := minuman.Nama

	stok := t.kulkas.GetStok(namaMinuman)
	t.kulkas.UpdateStok(namaMinuman, stok+qty)
}

func (t *toko) kurangStokMinuman(minuman Minuman, qty int) {
	namaMinuman := minuman.Nama

	stok := t.kulkas.GetStok(namaMinuman)
	t.kulkas.UpdateStok(namaMinuman, stok-qty)
}

func (t *toko) TotalStokMinuman(minuman Minuman) int {
	namaMinuman := minuman.Nama
	return t.kulkas.GetStok(namaMinuman)
}
