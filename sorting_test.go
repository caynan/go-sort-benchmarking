package sortingBenchmark

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func load(path string, inputSize int) []int {
	values := make([]int, inputSize)
	lines, err := readLines(path)
	for i, line := range lines {
		if i <= len(values) {
			values[i], err = strconv.Atoi(line)
		}
		if err != nil {
			fmt.Println(err)
		}
	}

	return values
}

func benchmarkMergeSort(path string, inputSize int, b *testing.B) {
	values := load(path, inputSize)
	// Reset timer after loading file
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		values = mergeSort(values)
	}
}

func benchmarkParallelMergeSort(path string, inputSize int, b *testing.B) {
	values := load(path, inputSize)
	c := make(chan []int)
	// Reset timer after loading file and creating channel
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		go parallelMergeSort(values, c)
		// we save result to file to avoid compiler optimizations
		values = <-c
	}
}

// Linear MergeSort
func BenchmarkMergesort10(b *testing.B)   { benchmarkMergeSort("10.txt", 10, b) }
func BenchmarkMergesort1k(b *testing.B)   { benchmarkMergeSort("1k.txt", 1000, b) }
func BenchmarkMergesort1kk(b *testing.B)  { benchmarkMergeSort("1kk.txt", 1000000, b) }
func BenchmarkMergesort10kk(b *testing.B) { benchmarkMergeSort("10kk.txt", 10000000, b) }

// Parallel MergeSort
func BenchmarkParallelMergesort10(b *testing.B)   { benchmarkParallelMergeSort("10.txt", 10, b) }
func BenchmarkParallelMergesort1k(b *testing.B)   { benchmarkParallelMergeSort("1k.txt", 1000, b) }
func BenchmarkParallelMergesort1kk(b *testing.B)  { benchmarkParallelMergeSort("1kk.txt", 1000000, b) }
func BenchmarkParallelMergesort10kk(b *testing.B) { benchmarkParallelMergeSort("10kk.txt", 10000000, b) }
