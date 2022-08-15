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

	// creating waitor to keep track of how far we have gotten
	var epoch sync.WaitGroup
	epoch.Add(num_batches)
	// channel to collect individual batches ready for ordering
	batch_channel := make(chan []*D, workers)
	ordered_channel := make(chan []*D)
	// dispatcher goroutine to keep dispatching new jobs so workers are always busy
	go func() {
		b := 0
		// kick off all starting workers
		for i := 0; i < workers; i++ {
			// guard against more workers than batches
			if i < num_batches {
				go func() {
					batch_channel <- getBatch(ds, i, batchSize, &dsidx)
					epoch.Done()
				}()
				b++
			}
		}
		// while not reached the end of number of batches
		for b < num_batches {
			next := <-batch_channel
			ordered_channel <- next
			time.Sleep(10 * time.Second)
			b++
			fmt.Println("processing batch ", b)
		}
		// closing channels when this epoch is complete
		// epoch.Wait()
		time.Sleep(10 * time.Second)
		fmt.Println("CLOSING CHANNELS")
		close(batch_channel)
		close(ordered_channel)
	}()
	fmt.Println("Begin the channeling 2")
	return ordered_channel, nil
}

// getBatch by number and batch size this will also reference the mapping for precisely which indexes to use from the dataset
func getBatch[D dataset.DatasetCompat[L], L dataset.LattigoCompat](ds dataset.Dataset[D, L], batch int, batchSize int, mapping *[]int) []*D {
	b := make([]*D, batchSize)
	for i := batch * batchSize; i < (batch+1)*batchSize; i++ {
		// in case we are to populate a partial batch check we have
		// not exceeded the bounds of mapping
		if i > len(*mapping) {
			break
		}
		sample, _ := ds.Get((*mapping)[i])
		b[i-(batch*batchSize)] = sample
	}
	return b
}
