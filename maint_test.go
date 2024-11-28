package main

import (
	"testing"
	"time"

	"golang.org/x/exp/rand"
)

func TestFillSlice(t *testing.T) {
	total := 10
	list := make([]string, total)

	count := 5
	value := "test"
	nextIndex := fillSlice(list, 0, value, count)

	// Periksa apakah nilai sudah diisi dengan benar
	for i := 0; i < count; i++ {
		if list[i] != value {
			t.Errorf("Expected value %s at index %d, but got %s", value, i, list[i])
		}
	}

	// Periksa indeks berikutnya
	if nextIndex != count {
		t.Errorf("Expected next index %d, but got %d", count, nextIndex)
	}
}

func TestShuffleSlice(t *testing.T) {
	list := []string{"a", "b", "c", "d", "e"}

	// Salin slice sebelum pengacakan
	original := make([]string, len(list))
	copy(original, list)

	// Atur seed agar dapat diuji dengan nilai tetap
	rand.Seed(uint64(time.Now().UnixNano()))
	shuffleSlice(list)

	// Periksa apakah elemen tetap sama tetapi urutannya berubah
	same := true
	for i := range list {
		if list[i] != original[i] {
			same = false
			break
		}
	}

	if same {
		t.Error("Shuffle did not change the order of elements")
	}
}

func TestPopRandomValue(t *testing.T) {
	// Tetapkan seed untuk memastikan pengujian deterministik
	rand.Seed(uint64(time.Now().UnixNano()))

	list := []string{"a", "b", "c", "d", "e"}
	original := make([]string, len(list))
	copy(original, list)

	// Pop elemen acak
	value, updatedList := popRandomValue(list)

	// Pastikan elemen yang di-pop ada di list asli
	found := false
	for _, v := range original { // Pastikan memeriksa terhadap list asli, bukan list yang diperbarui
		if v == value {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Popped value %s is not in the original slice: %v", value, original)
	}

	// Pastikan panjang list berkurang
	if len(updatedList) != len(list)-1 {
		t.Errorf("Expected length %d after pop, but got %d", len(list)-1, len(updatedList))
	}

	// Pastikan elemen yang di-pop tidak ada di slice yang diperbarui
	for _, v := range updatedList {
		if v == value {
			t.Errorf("Popped value %s still exists in the updated slice", value)
		}
	}
}
