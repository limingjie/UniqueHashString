package main

import (
	"runtime"
	"sync"

	"github.com/limingjie/UniqueHashString/golang/unihash"
)

func main() {
	// Start workers.
	var wg sync.WaitGroup
	tasks := make(chan unihash.Task)
	numCPU := runtime.NumCPU()
	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go unihash.Worker(i, &wg, tasks)
	}

	// Assign task to workers. Starts from 10^19.
	step := uint64(100001)
	for i := uint64(10000000000000000000); i < uint64(10000000000000000000+step*15); i += step {
		tasks <- unihash.Task{Left: i, Right: i + step}
	}

	// Close task channel.
	close(tasks)

	// Waiting for all workers complete.
	wg.Wait()
}
