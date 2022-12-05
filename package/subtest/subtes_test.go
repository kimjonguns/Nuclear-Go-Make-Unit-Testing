package subtest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/**
Apabila hanya ingin menjalan 1 sub test nya saja : go test -v -run=TestSubTest/func1
*/

func TestSubTest(t *testing.T) {
	t.Run("func1", func(t *testing.T) {
		result := SubTest("Hafi")
		require.Equal(t, "Hafi", result, "Harus Hafi")
	})
	t.Run("func2", func(t *testing.T) {
		result := SubTest("Hafi")
		require.Equal(t, "Hafi", result, "Harus Hafi")
	})
}
