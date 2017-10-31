package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
)

var randomBase64 = []byte("Nz746LU-BCcolIygTV9Z0GaeX8puRKO5PEisvWDt3qbnrdFhf1wAMkHxQ_2jYmSJ")

// Some more random base64 strings
//     "jt1vZpX9Fi6qHnRhrTb3-UuakzK0_JEW47wxeCO8Qf5IlgPsoYScDm2yNLdGAMBV";
//     "PGpN7Ws0gaFR6mvJT1UXl3bHBxtnuiyq-d9fj_wYckV2zSKIoA5rLMOeDC4ZhEQ8";
//     "3mqEds5hZkUyjD269ABplRHgI8iYzr-XOLbwF07ctou1SveV4KGQCMPNaxnTWfJ_";
//     "b9Tm5k27HuB-VyLEl13RIMwSKNGUDYQpnPhsJgavc6OiC4ofjFrAxd0_ztZqWX8e";
//     "Na3gFiQx1sS8LKyOuZrYBpjzwGEDPbomdq654RcIX_0e2C7k-WHnUhVAJlMf9Ttv";
//     "Zw84hDk-pN5uKcPy1_LdqIn0tQCJWBAROm26XSijzegxME7FHVbUaTGlorf3sYv9";
//     "fBe04QGcSkXsLud76gbxIFpOyUHajWiZmYMrEnDhtw5KqCRA8v3lPTz_o12VN9-J";
//     "RoFTY_jOZtbkai8651lp-VqzEgd4rLuDJ2WHBUv3xA9C0m7wKnsPhfMSQecGINXy";
//     "r3fi0dH_6kYyOaQ8s2eUBWucGS7PnNq9moFbTEh4C1xwMXJzIv-VZDljtRgLA5pK";
//     "vjOShxu1Cq8-JBsylNTGoiX5Kpt0cAEZr9VP2HMw3mkzFI4YL_bfRUegDWn7Qa6d";

var unRandomBase64 = make([]uint64, 128)

func encode(value uint64) (code [11]byte, size int) {
	var accumulate, remainder, position uint64

	for {
		accumulate += remainder
		remainder = value & 0x3f
		value >>= 6
		position = (accumulate + remainder) & 0x3f
		code[size] = randomBase64[position]
		size++

		if value == 0 {
			break
		}
	}

	return
}

func decode(code []byte) (value uint64) {
	var accumulate, remainder, position uint64

	size := len(code)
	for i := 0; i < size; i++ {
		position = unRandomBase64[code[i]]
		remainder = (position + 64 - accumulate) & 0x3f
		accumulate += remainder
		value |= remainder << uint64(6*i)
	}

	return
}

type task struct {
	left  uint64
	right uint64
}

func worker(id int, wg *sync.WaitGroup, inTask <-chan task) {
	for t := range inTask {
		for i := t.left; i < t.right; i++ {
			code, size := encode(i)
			decode(code[0:size])
			// value := decode(code[0:size])
			// fmt.Println(i, "->", string(code[0:size]), "->", value)

			// if i != value {
			// 	fmt.Println("Decode Error", i, "->", string(code[0:size]), "->", value)
			// }
		}

		fmt.Printf("Worker %d completed calculation of range [%d, %d).\n", id, t.left, t.right)
	}

	wg.Done()
}

func main() {
	// Code for Profiling
	// $ ./UniqueHashString -cpuprofile cpu.prof -memprofile mem.prof
	// $ go tool pprof UniqueHashString cpu.prof
	// (pprof) top
	// Check the top usage
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile `file`")
	var memprofile = flag.String("memprofile", "", "write memory profile to `file`")
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		defer f.Close()
	}
	// Code for Profiling End

	// Reverse random base64 into an array.
	for k, v := range randomBase64 {
		unRandomBase64[v] = uint64(k)
	}

	// Start workers.
	var wg sync.WaitGroup
	chTask := make(chan task)
	numCPU := runtime.NumCPU()
	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go worker(i, &wg, chTask)
	}

	// Assign task to workers. Starts from 10^19.
	step := uint64(6553600)
	for i := uint64(10000000000000000000); i < uint64(10000000000000000000+step*100); i += step {
		chTask <- task{i, i + step}
	}

	// Close task channel.
	close(chTask)

	// Waiting for all workers complete.
	wg.Wait()
}
