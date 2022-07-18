package dataloader

import (
	"fmt"
	"math/rand"
	"time"

	"gitlab.com/deepcypher/darklantern/dataset"
)

// http://www.golangpatterns.info/concurrency/generators
//  func generateRandomNumbers (n int) {
//     ch := make (chan float)
//     sem := make (semaphore, n)
//     for i := 0; i < n; i++ {
//         go func () {
//             ch <- rand.Float()
//             sem.Signal()
//         } ()
//     }
//     // launch extra goroutine to eventually close ch
//     go func () {
//         sem.Wait(n)
//         close(ch)
//     }
//     return ch
// }

// define semaphore types
type empty interface{}
type semaphore chan empty

// SimpleDataloader load data from dataset using goroutines concurrently based on number of workes, into batches, and whether or not to shuffle the dataset.
func SimpleDataloader[D dataset.DatasetCompat[L], L dataset.LattigoCompat](ds dataset.Dataset[D, L], workers int, batchSize int, shuffle bool) (dataset.Dataset[D, L], error) {
	// create a semaphore so we know when we are finished and to lock all resources
	length, _ := ds.Length()
	// constructing default mapping (before shuffling) to indicate which examples fall in which order
	mapping := make([]int, length)
	for i := 0; i < len(mapping); i++ {
		mapping[i] = i
	}
	// now shuffle the mapping if shuffle is true
	if shuffle == true {
		// seeding randomness
		// https://stackoverflow.com/a/46185753/11164973
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(length, func(i, j int) {
			// shuffling/ swapping by returned indexes
			mapping[i], mapping[j] = mapping[j], mapping[i]
		})
	}
	// semaphore to dictate when master channel should close
	sem := make(semaphore, length)
	// single aggregator channel of all workers for for loops
	ch := make(chan D)

	// go func() {
	// 	sem.Wait(length)
	// 	close(ch)
	// }
	// now we create the channels for each worker
	fmt.Println("mapping", mapping)
	fmt.Println("semaphore", sem)
	fmt.Println("channel", ch)
	fmt.Println("dataset", ds)
	fmt.Println("workers", workers)
	fmt.Println("batchSize", workers)
	return ds, nil
}
