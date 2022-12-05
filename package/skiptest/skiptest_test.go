package skiptest

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSkipTest(t *testing.T) {
	if runtime.GOOS == "windows" {
		fmt.Println("check ...")
		t.Skip("Windows tidak akan di check")
	}

	hasil := SkipTest("Hafi")
	require.Equal(t, "Hafi", hasil, "Bukan Hafi")
}
