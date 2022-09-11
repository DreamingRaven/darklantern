package dataloader

import (
	"fmt"
	"math"
	"math/rand"
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
	// var epoch sync.WaitGroup
	// epoch.Add(num_batches)
	// channel to collect individual batches ready for ordering
	batch_channel := make(chan knownBatch[D, L], workers)
	ordered_channel := make(chan []*D)

	// dispatcher goroutine to keep dispatching new jobs so workers are always busy
	go func() {
		b := 0
		bCache := make([]*[]*D, num_batches)
		fmt.Println("Launching initial workers...")
		// launch full set of initial workers
		for i := 0; i < workers; i++ {
			// goroutine + channel wrapped worker
			fmt.Println("Worker:", i, "launch")
			go func(batch int) {
				batch_channel <- getBatch(ds, batch, batchSize, &dsidx)
				// epoch.Done()
			}(i)
		}
		nextWorker := workers

		fmt.Println("Managing workers and worker channel...")
		// watch channel for worker outputs and re-schedule completed workers
		// TODO workers are overunning we are queueing too many
		for kb := range batch_channel {
			// exit from infinite loop
			fmt.Printf("loop b:%v, cache:%v\n", b, bCache)
			// fmt.Println(kb)
			if b == num_batches {
				break
			}
			// if the next batch in the channel is the next batch in sequence
			if kb.metadata.id == b {
				// return in channel
				ordered_channel <- *kb.batch
				// requeue worker as next worker
				go func(batch int) {
					batch_channel <- getBatch(ds, batch, batchSize, &dsidx)
					// epoch.Done()
				}(nextWorker)
				nextWorker += 1
				b += 1
			} else if kb.metadata.id < num_batches {
				// insert into cache
				bCache[kb.metadata.id] = kb.batch
				// do not requeue worker
			}
			// now check if we have any thing(s) in cache waiting to be sent off
			for range bCache {
				if b == num_batches {
					break
				}
				if bCache[b] != nil {
					// take from cache into channel
					ordered_channel <- *bCache[b]
					// nillify cache position
					bCache[b] = nil
					// requeue worker as next worker
					go func(batch int) {
						batch_channel <- getBatch(ds, batch, batchSize, &dsidx)
						// epoch.Done()
					}(nextWorker)
					nextWorker += 1
					b += 1
				} else {
					// early exit from this batch cache scan as the successor is not in the cache
					break
				}
			}

			// // we dont want to bomb resources so sleep until next loop
			// time.Sleep(1 * time.Second)
		}
		// closing channels when this epoch is complete
		// epoch.Wait()
		// time.Sleep(10 * time.Second)
		fmt.Println("CLOSING CHANNELS")
		close(batch_channel)
		close(ordered_channel)
	}()
	fmt.Println("Begin Channeling")
	return ordered_channel, nil
}

type workerMetadata struct {
	id int
}

type knownBatch[D dataset.DatasetCompat[L], L dataset.LattigoCompat] struct {
	metadata workerMetadata
	batch    *[]*D
}

// getBatch by number and batch size this will also reference the mapping for precisely which indexes to use from the dataset
func getBatch[D dataset.DatasetCompat[L], L dataset.LattigoCompat](ds dataset.Dataset[D, L], batch int, batchSize int, mapping *[]int) knownBatch[D, L] {
	b := make([]*D, batchSize)
	for i := batch * batchSize; i < (batch+1)*batchSize && i < len(*mapping); i++ {
		sample, _ := ds.Get((*mapping)[i])
		b[i-(batch*batchSize)] = sample
	}
	kb := knownBatch[D, L]{metadata: workerMetadata{id: batch}, batch: &b}
	return kb
}
