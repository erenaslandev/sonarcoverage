package sonarcoverage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	t.Run("without cache fibonacci 0", func(t *testing.T) {
		actual := FibonacciWithoutCache(0)
		assert.Equal(t, 0, actual)
	})

	t.Run("without cache fibonacci 1", func(t *testing.T) {
		actual := FibonacciWithoutCache(1)
		assert.Equal(t, 1, actual)
	})

	t.Run("without cache fibonacci 10", func(t *testing.T) {
		actual := FibonacciWithoutCache(10)
		assert.Equal(t, 55, actual)
	})

	t.Run("without cache fibonacci 50", func(t *testing.T) {
		actual := FibonacciWithoutCache(50)
		assert.Equal(t, 12586269025, actual)
	})

	t.Run("with cache fibonacci 0", func(t *testing.T) {
		actual := FibonacciWithCache(0)
		assert.Equal(t, 0, actual)
	})

	t.Run("with cache fibonacci 1", func(t *testing.T) {
		actual := FibonacciWithCache(1)
		assert.Equal(t, 1, actual)
	})

	t.Run("with cache fibonacci 10", func(t *testing.T) {
		actual := FibonacciWithCache(10)
		assert.Equal(t, 55, actual)
	})

	t.Run("with cache fibonacci 50", func(t *testing.T) {
		actual := FibonacciWithCache(50)
		assert.Equal(t, 12586269025, actual)
	})
}

func BenchmarkFibonacci(b *testing.B) {
	b.Run("without cache fibonacci 50", func(b *testing.B) {
		actual := FibonacciWithoutCache(50)
		assert.Equal(b, 12586269025, actual)
	})

	b.Run("with cache fibonacci 50", func(b *testing.B) {
		actual := FibonacciWithCache(50)
		assert.Equal(b, 12586269025, actual)
	})
}
