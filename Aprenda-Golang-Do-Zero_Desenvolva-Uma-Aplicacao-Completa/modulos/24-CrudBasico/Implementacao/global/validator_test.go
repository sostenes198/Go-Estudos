package global_test

import (
	"github.com/stretchr/testify/require"
	"implementacao/global"
	"sync"
	"testing"
)

func TestValidator(t *testing.T) {

	// arrange
	var waitGroup sync.WaitGroup
	var internalValidator1 *global.ValidatorGlobal
	var internalValidator2 *global.ValidatorGlobal
	var internalValidator3 *global.ValidatorGlobal

	waitGroup.Add(3)

	// act
	go func() {
		internalValidator1 = global.Validator()
		waitGroup.Done()
	}()

	go func() {
		internalValidator2 = global.Validator()
		waitGroup.Done()
	}()

	go func() {
		internalValidator3 = global.Validator()
		waitGroup.Done()
	}()

	waitGroup.Wait()

	// assert

	require.Equal(t, internalValidator1, internalValidator2)
	require.Equal(t, internalValidator1, internalValidator3)
	require.Equal(t, internalValidator2, internalValidator3)
}
