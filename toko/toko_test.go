package toko

import (
	"testing"

	"gotest.tools/assert"
)

func TestWhenTokoBuka(t *testing.T) {
	toko := NewToko()
	toko.Buka()

	assert.Equal(t, true, toko.IsBuka())
}

func TestWhenTokoTutup(t *testing.T) {
	toko := NewToko()
	toko.Tutup()

	assert.Equal(t, false, toko.IsBuka())
}

func TestTambahStokMinuman(t *testing.T) {
	toko := NewToko()
	toko.Buka()
	m := Minuman{
		Nama:  "Sprite",
		Harga: 5000,
	}
	toko.TambahStokMinuman(m, 10)
	result := toko.JualMinuman(m, 10)

	assert.Equal(t, true, result)
}

func TestJualMinumanWhenNotAvailable(t *testing.T) {
	toko := NewToko()
	toko.Buka()
	m := Minuman{
		Nama:  "Sprite",
		Harga: 5000,
	}
	toko.TambahStokMinuman(m, 1)
	result := toko.JualMinuman(m, 10)

	assert.Equal(t, false, result)
}
