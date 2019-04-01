package toko

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"gotest.tools/assert"
)

type MockTokoStorage struct {
	mock.Mock
}

func (m *MockTokoStorage) UpdateStok(nama string, qty int) {
	m.Called(nama, qty)
	return
}

func (m *MockTokoStorage) IsStokAvailable(nama string, qty int) bool {
	args := m.Called(nama, qty)
	return args.Bool(0)
}

func (m *MockTokoStorage) GetStok(nama string) int {
	args := m.Called(nama)
	return args.Get(0).(int)
}

func TestWhenTokoBuka(t *testing.T) {
	ts := new(MockTokoStorage)
	toko := NewToko(ts)
	toko.Buka()

	assert.Equal(t, true, toko.IsBuka())
}

func TestWhenTokoTutup(t *testing.T) {
	ts := new(MockTokoStorage)
	toko := NewToko(ts)

	toko.Tutup()

	assert.Equal(t, false, toko.IsBuka())
}

func TestTambahStokMinuman(t *testing.T) {
	ts := new(MockTokoStorage)
	toko := NewToko(ts)

	m := Minuman{
		Nama:  "Sprite",
		Harga: 5000,
	}

	expected := 1
	ts.On("GetStok", m.Nama).Return(0)
	ts.On("UpdateStok", m.Nama, expected)

	toko.TambahStokMinuman(m, expected)
	ts.AssertNumberOfCalls(t, "GetStok", 1)
	ts.AssertNumberOfCalls(t, "UpdateStok", 1)
}

func TestJualMinuman(t *testing.T) {
	ts := new(MockTokoStorage)
	toko := NewToko(ts)

	m := Minuman{
		Nama:  "Sprite",
		Harga: 5000,
	}

	ts.On("UpdateStok", m.Nama, 0)
	ts.On("IsStokAvailable", m.Nama, 1).Return(true)
	ts.On("GetStok", m.Nama).Return(1)

	result := toko.JualMinuman(m, 1)

	assert.Equal(t, true, result)
}

func TestJualMinumanNotAvailable(t *testing.T) {
	ts := new(MockTokoStorage)
	toko := NewToko(ts)

	m := Minuman{
		Nama:  "Sprite",
		Harga: 5000,
	}
	expected := 1
	ts.On("UpdateStok", m.Nama, expected)
	ts.On("IsStokAvailable", m.Nama, expected).Return(false)
	ts.On("GetStok", m.Nama).Return(1)

	result := toko.JualMinuman(m, 1)

	assert.Equal(t, false, result)
}

func TestTotalStokMinuman(t *testing.T) {
	ts := new(MockTokoStorage)
	toko := NewToko(ts)

	m := Minuman{
		Nama:  "Sprite",
		Harga: 5000,
	}
	expected := 1
	ts.On("GetStok", m.Nama).Return(expected)

	result := toko.TotalStokMinuman(m)

	assert.Equal(t, true, expected == result)
}
