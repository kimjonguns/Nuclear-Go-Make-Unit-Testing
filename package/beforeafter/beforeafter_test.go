package beforeafter

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Jalan")
	m.Run()
	fmt.Println("Berhenti")
}

func TestBeforeAfter(t *testing.T) {
	val := BeforeAfter("Hafi")

	if val != "Hafi" {
		t.Fatal()
	}

	fmt.Println("Berhasil")
}
