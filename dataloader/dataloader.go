package dataloader

import (
	"fmt"

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
func SimpleDataloader[D dataset.DatasetCompat[L], L dataset.LattigoCompat](ds dataset.Dataset[D, L], workers int, batchSize int, shuffle bool) {
	fmt.Println("dataset", ds)
	fmt.Println("workers", workers)
	fmt.Println("batchSize", workers)
}
