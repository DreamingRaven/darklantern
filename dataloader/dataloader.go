package dataloader

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
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

// SimpleDataloader load data from dataset using goroutines concurrently based on number of workes, into batches, and whether or not to shuffle the dataset.
func SimpleDataloaderOri[D dataset.DatasetCompat[L], L dataset.LattigoCompat](ds dataset.Dataset[D, L], workers int, batchSize int, shuffle bool, allowSmallBatch bool) (chan D, error) {
	// create a semaphore so we know when we are finished and to lock all resources
	length, _ := ds.Length()
	// constructing default mapping (dsidx) (before shuffling) to indicate which examples fall in which order
	dsidx := make([]int, length)
	for i := 0; i < len(dsidx); i++ {
		dsidx[i] = i
	}
	// now shuffle the mapping if shuffle is true
	if shuffle == true {
		// seeding randomness
		// https://stackoverflow.com/a/46185753/11164973
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(length, func(i, j int) {
			// shuffling/ swapping by returned indexes
			dsidx[i], dsidx[j] = dsidx[j], dsidx[i]
		})
	}

	sample := 0
	proportion := float64(length) / float64(batchSize)
	num_batches := int(0)
	if allowSmallBatch == true {
		num_batches = int(math.Ceil(proportion))
	} else {
		num_batches = int(math.Floor(proportion))
	}

	// semaphore to dictate when master channel should close
	// sem := make(semaphore, length)
	var wg sync.WaitGroup
	// single aggregator channel of all workers for for loops
	ch := make(chan D)
	// for each batch (zero indexed)
	for i := 0; i < num_batches; i++ {
		wg.Add(1)
		go func(i int) {
			batch := make([]*D, batchSize)
			// for each slot in batch
			for j := 0; j < batchSize; j++ {
				example, _ := ds.Get(dsidx[sample])
				batch[j] = example
				// sem <- empty
				sample++
				wg.Done()
			}
		}(i)
		if i%workers == 0 {
			wg.Wait()
		}
		// fmt.Println(fmt.Sprintf("sample:%v,bid:%v,batch:%v", sample, i, batch))
	}
	// bid := 0
	// batch := make([]D, batchSize)
	// for i := 0; i < len(dsidx); i++ {
	// 	if i
	// 	data, _ := ds.Get(dsidx[i])
	// 	fmt.Println("i", i, "dsidx", dsidx[i], "=", data)
	// }

	go func() {
		// sem.Wait(length)
		close(ch)
	}()
	// now we create the channels for each worker
	fmt.Println("dsidx", dsidx)
	// fmt.Println("semaphore", sem)
	fmt.Println("channel", ch)
	fmt.Println("dataset", ds)
	fmt.Println("workers", workers)
	fmt.Println("batchSize", workers)
	return ch, nil
}

// SimpleDataloader loads arbitrary dataset into batches and divides work between workers
func SimpleDataloader[D dataset.DatasetCompat[L], L dataset.LattigoCompat](ds dataset.Dataset[D, L], workers int, batchSize int, shuffle bool, allowSmallBatch bool) (chan []*D, error) {
	// create a semaphore so we know when we are finished and to lock all resources
	length, _ := ds.Length()
	// constructing default mapping (dsidx) (before shuffling) to indicate which examples fall in which order
	dsidx := make([]int, length)
	for i := 0; i < len(dsidx); i++ {
		dsidx[i] = i
	}
	// now shuffle the mapping if shuffle is true
	if shuffle == true {
		// seeding randomness
		// https://stackoverflow.com/a/46185753/11164973
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(length, func(i, j int) {
			// shuffling/ swapping by returned indexes
			dsidx[i], dsidx[j] = dsidx[j], dsidx[i]
		})
	}

	// calculate specific number of batches in data to be returned
	proportion := float64(length) / float64(batchSize)
	num_batches := int(0)
	if allowSmallBatch == true {
		num_batches = int(math.Ceil(proportion))
	} else {
		num_batches = int(math.Floor(proportion))
	}
	fmt.Println(num_batches)

	// creating waitor to keep track of how far we have gotten
	var epoch sync.WaitGroup
	epoch.Add(num_batches)
	// channel to collect individual batches ready for ordering
	batch_channel := make(chan []*D, workers)
	ordered_channel := make(chan []*D)
	// dispatcher goroutine to keep dispatching new jobs so workers are always busy
	go func() {
		for i := 0; i < num_batches; i++ {

		}
		// closing channels when epochs are complete
		// epoch.Wait()
		close(batch_channel)
		close(ordered_channel)
	}()
	return ordered_channel, nil
}

// getBatch by number and batch size this will also reference the mapping for precisely which indexes to use from the dataset
func getBatch[D dataset.DatasetCompat[L], L dataset.LattigoCompat](ch *chan []*D, ds dataset.Dataset[D, L], wg *sync.WaitGroup, batch int, batchSize int, mapping *[]int) {
	defer wg.Done()
	b := make([]*D, batchSize)
	for i := batch * batchSize; i < (batch+1)*batchSize; i++ {
		sample, _ := ds.Get((*mapping)[i])
		b[i-(batch*batchSize)] = sample
	}
	*ch <- b
}
