package toko

type Toko interface {
	Buka()
	IsBuka() bool
	Tutup()
	TambahStokMinuman(minuman Minuman, qty int)
	TotalStokMinuman(minuman Minuman) int
	JualMinuman(minuman Minuman, qty int) bool
}

type toko struct {
	isBuka      bool
	stokMinuman map[string]int
}

func NewToko() Toko {
	toko := &toko{}

	return toko
}

func (t *toko) Buka() {
	t.isBuka = true
	t.stokMinuman = make(map[string]int)
}

func (t *toko) IsBuka() bool {
	return t.isBuka
}

func (t *toko) Tutup() {
	t.isBuka = false
}

func (t *toko) JualMinuman(minuman Minuman, qty int) bool {
	if t.minumanIsAvailable(minuman, qty) {
		t.kurangStokMinuman(minuman, qty)
		return true
	}

	return false
}

func (t *toko) TambahStokMinuman(minuman Minuman, qty int) {
	namaMinuman := minuman.Nama

	stok := t.TotalStokMinuman(minuman)
	t.stokMinuman[namaMinuman] = stok + qty
}

func (t *toko) kurangStokMinuman(minuman Minuman, qty int) {
	namaMinuman := minuman.Nama

	stok := t.TotalStokMinuman(minuman)
	if stok > 0 {
		t.stokMinuman[namaMinuman] = stok - qty
	}
}

func (t *toko) minumanIsAvailable(minuman Minuman, qty int) bool {
	return t.TotalStokMinuman(minuman) >= qty
}

func (t *toko) TotalStokMinuman(minuman Minuman) int {
	namaMinuman := minuman.Nama

	if stok, ok := t.stokMinuman[namaMinuman]; ok {
		return stok
	}

	return 0
}
