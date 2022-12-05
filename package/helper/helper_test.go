package helper

import (
	"fmt"
	"testing"
)

// *testing.T testing ini untuk unit test

// menjalankan testing dengan cara : go test
// menjalankan testing dengan melihat detail function test apa saja yang sudah running : go test -v
// menjalankan testing dengan spesifik function : go test -v -run=TestHaloDunia

/*
*
cara menggagalkan unit test jangan memakai panic, tetapi gunakan:
Fail() -> apabila gagal, tetap dieksekusi
FailNow() -> apabila gagal, tanpa melanjutkan eksekusi
Error() -> apabila gagal, akan melakukan log (print) Error, setelah itu akan secara otomatis memanggil

	function Fail()

Fatal() -> apabila gagal, akan melakukan log (print) Error, setelah itu akan secara otomatis memanggil

	function FailNow()
*/
func TestHaloDunia(t *testing.T) {
	result := HaloDunia("Haf")
	if result != "Hafi" {
		t.FailNow()
	}

	fmt.Println("Check selesai")
}

func TestHaloDunia2(t *testing.T) {
	result := HaloDunia("Haf")
	if result != "Hafi" {
		t.Fail()
	}

	fmt.Println("Check selesai")
}

func TestHaloDunia3(t *testing.T) {
	result := HaloDunia("Haf")
	if result != "Hafi" {
		t.Error()
	}

	fmt.Println("Check selesai")
}

func TestHaloDunia4(t *testing.T) {
	result := HaloDunia("Haf")
	if result != "Hafi" {
		t.Fatal()
	}

	fmt.Println("Check selesai")
}
