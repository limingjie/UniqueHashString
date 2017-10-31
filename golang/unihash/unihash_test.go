package unihash_test

import (
	"runtime"
	"sync"
	"testing"

	"../unihash"
)

// TestEncode - Unit Test
func TestEncode(t *testing.T) {
	inputs := []uint64{
		0, 1, 63, 64, 4095, 4096, 65535, 65536,
		12345678912345678912, 12345678912345678913, 12345678912345678914,
		12345678912345678915, 12345678912345678916, 12345678912345678917,
		18446744073709551614, 18446744073709551615,
	}
	outputs := []string{
		"N", "z", "J", "Nz", "JS", "NNz", "JSI", "NNT",
		"NVsYmjZmVSB", "z9vmSY0S9JC", "7ZWSJmGJZNc",
		"40DJNSaN0zo", "6GtNzJezG7l", "La3z7NX7a4I",
		"SmYj2_QxHk6", "JSmYj2_QxHL",
	}

	for index, input := range inputs {
		code, size := unihash.Encode(input)
		if string(code[:size]) != outputs[index] {
			t.Errorf("Encode(%d) failed, expected \"%s\", got \"%s\".", input, outputs[index], string(code[:size]))
		}
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unihash.Encode(uint64(i))
	}
}

func TestDecode(t *testing.T) {
	inputs := []string{
		"N", "z", "J", "Nz", "JS", "NNz", "JSI", "NNT",
		"NVsYmjZmVSB", "z9vmSY0S9JC", "7ZWSJmGJZNc",
		"40DJNSaN0zo", "6GtNzJezG7l", "La3z7NX7a4I",
		"SmYj2_QxHk6", "JSmYj2_QxHL",
	}
	outputs := []uint64{
		0, 1, 63, 64, 4095, 4096, 65535, 65536,
		12345678912345678912, 12345678912345678913, 12345678912345678914,
		12345678912345678915, 12345678912345678916, 12345678912345678917,
		18446744073709551614, 18446744073709551615,
	}

	for index, input := range inputs {
		value := unihash.Decode([]byte(input))
		if value != outputs[index] {
			t.Errorf("Decode(\"%s\") failed, expected %d, got %d.", input, outputs[index], value)
		}
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unihash.Decode([]byte("NVsYmjZmVSB"))
	}
}

func TestWorker(t *testing.T) {
	var wg sync.WaitGroup
	tasks := make(chan unihash.Task)
	numCPU := runtime.NumCPU()
	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go unihash.Worker(i, &wg, tasks)
	}

	// Assign task to workers. Starts from 10^19.
	step := uint64(100)
	for i := uint64(10000000000000000000); i < uint64(10000000000000000000+step*uint64(numCPU)); i += step {
		tasks <- unihash.Task{i, i + step}
	}

	// Close task channel.
	close(tasks)

	// Waiting for all workers complete.
	wg.Wait()
}

func BenchmarkWorker(b *testing.B) {
	var wg sync.WaitGroup
	tasks := make(chan unihash.Task)
	numCPU := runtime.NumCPU()
	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go unihash.Worker(i, &wg, tasks)
	}

	// Assign task to workers. Starts from 10^19.
	step := uint64(1001001)
	for i := uint64(10000000000000000000); i < uint64(10000000000000000000+step*uint64(b.N)); i += step {
		tasks <- unihash.Task{i, i + step}
	}

	// Close task channel.
	close(tasks)

	// Waiting for all workers complete.
	wg.Wait()
}
