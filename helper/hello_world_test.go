package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

// Benchmark Test
/*
- run semua unit test dan benchmark pada module: go test -v -bench=.
- run hanya benchmark tanpa unit test: go test -v -run=NotMatchUnitTest -bench=.
- run benchmark tertentu tanpa unit test: go test -v -run=NotMatchUnitTest -bench=nama_benchmark_test
- run dari root folder: go test -bench=. ./...
*/
func BenchmarkHelloWorldRafli(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Rafli")
	}
}

func BenchmarkHelloWorldBima(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Bima")
	}
}

// Benchmark Sub
// how to run: go test -v -run=TestTidakAda -bench=BenchmarkSub
// jika ingin run ke sub-nya: go test -v -run=TestTidakAda -bench=BenchmarkSub/nama_sub
func BenchmarkSub(b *testing.B) {
	b.Run("Rafli", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Rafli")
		}
	})

	b.Run("Bima", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Bima")
		}
	})
}

// Benchmark Table
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Rafli",
			request: "Rafli",
		},
		{
			name:    "Bima",
			request: "Bima",
		},
		{
			name:    "Pratandra",
			request: "Pratandra",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.name)
			}
		})
	}
}

// how to run unit test: go test -v
// run 1 func: go test -v -run=nama_func
// menjalankan semua test: go test ./...

// Table Test
func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Rafli",
			request:  "Rafli",
			expected: "Hello Rafli",
		},
		{
			name:     "Bima",
			request:  "Bima",
			expected: "Hello Bima",
		},
		{
			name:     "Pratandra",
			request:  "Pratandra",
			expected: "Hello Pratandra",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "Result must be "+test.expected)
		})
	}
}

// Sub Test
// jika ingin running ke sub-nya: go test -v -run=TestSubTest/nama_sub
func TestSubTest(t *testing.T) {
	t.Run("Rafli", func(t *testing.T) {
		result := HelloWorld("rafli")
		assert.Equal(t, "Hello rafli", result, "Result must be 'Hello rafli'")
	})
	t.Run("Bima", func(t *testing.T) {
		result := HelloWorld("bima")
		assert.Equal(t, "Hello bima", result, "Result must be 'Hello bima'")
	})
}

// before dan after test | hanya berjalan pada 1 package, jika ingin berjalan pada package lain maka harus membuat TestMain lagi
func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

/*
library Testify ada 2 package, assert dan require.
- assert memanggil Fail(), eksekusi unit test akan tetap dilanjutkan
- require memanggil FailNow(), eksekusi unit test tidak akan dilanjutkan
*/
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("rafli")
	assert.Equal(t, "Hello rafli", result, "Result must be 'Hello rafli'")
	fmt.Println("TestHelloWorld with assert done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("rafli")

	require.Equal(t, "Hello rafli", result, "Result must be 'Hello rafli'")
	fmt.Println("TestHelloWorld with required done") // tidak akan terpanggil karena require memanggil FailNow()
}

// Skip Test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("can't run on windows")
	}

	result := HelloWorld("rafli")
	assert.Equal(t, "Hello rafli", result, "Result must be 'Hello rafli'")
}

func TestHelloWorldRafli(t *testing.T) {
	result := HelloWorld("rafli")

	if result != "Hello rafli" {
		//t.Fail() // tetap menggagalkan test tetapi tetap menjalankan program hingga akhir
		t.Error("Result must be 'Hello rafli'") // akan output log print dengan args di parameter, lalu memanggil Fail()
	}

	fmt.Println("TestHelloWorldRafli done")
}

func TestHelloWorldBima(t *testing.T) {
	result := HelloWorld("Bimbim")

	if result != "Hello Bimbim" {
		//t.FailNow() // menggagalkan test dengan tidak melanjutkan program
		t.Fatal("Result must be 'Hello Bimbim'") // mirip seperti t.Error(), beda-nya adalah memanggil FailNow()
	}

	fmt.Println("TestHelloWorldBima done")
}
