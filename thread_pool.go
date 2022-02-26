//    Copyright 2022 Nikhil Chawla

//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at

//        http://www.apache.org/licenses/LICENSE-2.0

//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package golang_thread_pool

import (
	"log"
	"sync"
)

// ----------------------------------------------------------------
type AsyncProgInterface interface {
	Execute(thread_id int)
}

// ----------------------------------------------------------------

type ThreadPool struct {
	/*Just a Queue Representation*/
	waitQueue []AsyncProgInterface

	/*To maintain consistency in queue data*/
	mutex *sync.Mutex

	/* Number of worker threads*/
	numWorkerThreads int

	/* A counter to maintain number of user workers at any point of time*/
	threadUsed int

	/*
		This is an important part of Thread pooling, since, we don't want to accept all requests upfront,
		which could cause pressue on memory, we maintain a certain count of inflight request to be pushed
		into the queue, if some more requests are getting produced, we safely block the producer before
		pushing them into the queue.
	*/
	inFlightRequestThreshold int

	/*
		needSignal is boolean alert, that alerts the pop operation that producer is stuck at pushing,
		and since popping is complete, it can unblock the push operation from producer.
	*/
	needSignal bool

	// signalChan is used to send the signal from pop operation that unblocks push operation
	signalChan chan bool

	// These are the threads information maintaing DS
	threadsAttr  map[int]*ThreadInfo
	threadIdsArr []int
}

// INTERNAL MECHANICS --------------------------------------------------------------------

// blockinitiateBlockUntil establishes the block for push operation and also initiates a pulse,
// communication channel with pop operation.
func (tp *ThreadPool) block() {
	tp.mutex.Lock()
	defer tp.mutex.Unlock()

	tp.needSignal = true
	tp.signalChan = make(chan bool)
}

// checkAdmissionControl is the important aspect of Staged Event Driven Architecture(SEDA), which
// basically executed the validations before enqueing, if its safe to push to the queue.
func (tp *ThreadPool) checkAdmissionControl() bool {
	if len(tp.waitQueue) < tp.inFlightRequestThreshold {
		return true
	}
	return false
}

// Typical queue pop operation
func (tp *ThreadPool) pop() AsyncProgInterface {

	tp.mutex.Lock()
	size := len(tp.waitQueue)
	if size > 0 {
		item := tp.waitQueue[size-1]
		tp.waitQueue = tp.waitQueue[:size-1]

		// pop operation is complete, check if producer is blocked
		// and needs a signal to unblock push operation.
		if tp.needSignal {
			// handling the event by sending a message through channel
			tp.signalChan <- true
			tp.needSignal = false
		}

		tp.mutex.Unlock() // until this pop operation reaches the termination point, we have maintained the lock on queue
		return item
	}
	return nil
}

// WaitForSignal blocks the push operation
func (tp *ThreadPool) waitForSignal() {

	<-tp.signalChan // Don't need the lock, because we are at receiving end.
	tp.mutex.Lock()
	close(tp.signalChan) // We need the lock to close this channel, preventing any leaks.
	tp.mutex.Unlock()
	return
}

// EXPOSED Methods --------------------------------------------------------------------

// FreeThread releases the thread that can be reused/consumed further.
func (tp *ThreadPool) FreeThread(thread_id int) {
	tp.mutex.Lock()
	defer tp.mutex.Unlock()

	log.Println("Releasing the thread id...", thread_id)
	tp.threadsAttr[thread_id].SetFree()
	tp.threadUsed -= 1
}

// ThreadPool Constructor
func NewThreadPool(numWorkerThreads int) *ThreadPool {

	tp := new(ThreadPool)
	tp.waitQueue = []AsyncProgInterface{}
	tp.numWorkerThreads = numWorkerThreads
	tp.mutex = &sync.Mutex{}

	tp.threadsAttr = map[int]*ThreadInfo{}
	tp.threadIdsArr = make([]int, numWorkerThreads)
	tp.threadUsed = 0

	tp.inFlightRequestThreshold = 5 /* mutable through SetInFlightRequestThreshold() */

	for ii := 0; ii < numWorkerThreads; ii++ {
		th := NewThread()
		thID := th.GetId()
		tp.threadsAttr[thID] = th
		tp.threadIdsArr[ii] = thID
	}

	// Signal handler for pop and push
	tp.needSignal = false

	return tp
}

// Typical queue push operation
func (tp *ThreadPool) Push(item AsyncProgInterface) {

	// Check if its safe to push.
	if tp.checkAdmissionControl() == false {
		tp.block()
		tp.waitForSignal() // since, the signal block is initiated, we need to wait for release
	}

	tp.mutex.Lock()
	tp.waitQueue = append(tp.waitQueue, item)
	tp.mutex.Unlock()
}

// Round Robin Scheduler
func (tp *ThreadPool) RunSchedulerInBackground() {
	go func(){
		idx := 0
		for {
			if len(tp.waitQueue) > 0 {
				thid := tp.threadIdsArr[idx]
				threadObject := tp.threadsAttr[thid]
				if threadObject.IsBusy() == false {
					item := tp.pop()
					tp.mutex.Lock()
					threadObject.SetBusy()
					tp.threadUsed += 1
					tp.mutex.Unlock()
					go item.Execute(threadObject.GetId())
				}
			}
			idx++
			if idx >= tp.numWorkerThreads {
				idx = 0
			}
		}
	}()
}

// SetInFlightRequestThreshold allows to mutate the inflight requests threshold
func (tp *ThreadPool) SetInFlightRequestThreshold(count int) {
	tp.inFlightRequestThreshold = count
	return
}

// ----------------------------------------------------------------
