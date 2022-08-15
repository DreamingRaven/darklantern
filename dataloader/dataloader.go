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
		// kick off all starting workers
		for i := 0; i < workers; i++ {
			go getBatch(&batch_channel, ds, &epoch, i, batchSize, &dsidx)
		}
		// while not reached the end of number of batches
		b := 0 + workers
		for b < num_batches {
			// if workers have not finished but the next in line is in cache pipe it to channel
			// if worker has finished but is not the next cache result and start next
			// if worker has finished and is the very next in line pipe into channel and start next
			time.Sleep(100 * time.Second)
			b++
		}
		// closing channels when this epoch is complete
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
