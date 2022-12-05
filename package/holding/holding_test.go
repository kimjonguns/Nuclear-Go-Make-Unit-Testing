package holding

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/**
pengganti dari if else adalah assertion dengan menggunakan library Testify
*/

func TestSumData(t *testing.T) {
	hasil := SumData(2, 1)
	assert.Greater(t, hasil, 0, "Hasil seharusnya lebih dari 0")
	fmt.Println("Berhasil")
}

func TestSumData2(t *testing.T) {
	hasil := SumData(-2, 1)
	require.Greater(t, hasil, 0, "Hasil seharusnya lebih dari 0")
	fmt.Println("Berhasil")
}
