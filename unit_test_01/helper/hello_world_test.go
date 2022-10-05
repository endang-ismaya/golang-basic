package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWord(t *testing.T) {
	result := HelloWorld("Endang")

	if result != "Hello Endang" {
		// error
		t.Fail()
	}

	fmt.Println("TestHelloWorld Endang")
}

func TestHelloWord2(t *testing.T) {
	result := HelloWorld("Endang 2")

	if result != "Hello Endang 2" {
		// error
		t.FailNow()
	}

	fmt.Println("TestHelloWorld Endang 2")
}

func TestHelloWordError(t *testing.T) {
	result := HelloWorld("Endang")

	if result != "Hello Endang Error" {
		// error
		t.Error("Harusnya: Hello Endang Error")
	}

	fmt.Println("TestHelloWorld Endang Error")
}

func TestHelloWordFatal(t *testing.T) {
	result := HelloWorld("Endang")

	if result != "Hello Endang Fatal" {
		// error
		t.Fatal("Harusnya: Hello Endang Fatal")
	}

	fmt.Println("TestHelloWorld Endang Fatal")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Endang")
	assert.Equal(t, "Hi Endang", result, "should be 'Hello Endang'")
	fmt.Println("TestHelloWorldAssert Completed")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Endang")
	require.Equal(t, "Hi Endang", result, "should be 'Hello Endang'")
	fmt.Println("TestHelloWorldAssert Require Completed")
}

func TestSkip(t *testing.T) {
	operationSystem := runtime.GOOS
	if operationSystem == "linux" {
		t.Skip("Can't run on Linux")
	}

	result := HelloWorld("Endang")
	require.Equal(t, "Hello Endang", result, "Result must be 'Hello Endang'")
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("Connection to database...")

	m.Run()

	// after
	fmt.Println("Disconect database...")
}

func TestSubTest(t *testing.T) {
	t.Run("Endang", func(t *testing.T) {
		result := HelloWorld("Endang")
		require.Equal(t, "Hello Endang", result, "Result must be 'Hello Endang'")
	})

	t.Run("Wijaya", func(t *testing.T) {
		result := HelloWorld("Wijaya")
		require.Equal(t, "Hello Wijaya", result, "Result must be 'Hello Wijaya'")
	})
}
