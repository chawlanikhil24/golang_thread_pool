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

// import (
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"reflect"
// 	"time"

// 	thread_pool "github.com/chawlanikhil24/golang_thread_pool"
// )

// ----------------------------------------------------------------

// type StructExample struct {
// 	objId      *int
// 	objId_str  *string
// 	threadPool *thread_pool.ThreadPool

// 	sErr           chan error
// 	callback_type1 *func(chan error, *int, int)
// }

// func StructExamplePtr() *StructExample {
// 	obj := new(StructExample)
// 	obj.objId = nil
// 	obj.objId_str = nil
// 	obj.sErr = nil
// 	return obj
// }

// func (ob *StructExample) SetCallBack(cb interface{}) error {
// 	var err error = nil
// 	switch reflect.TypeOf(cb) {

// 	case reflect.TypeOf(ob.callback_type1):
// 		ob.callback_type1 = cb.(*func(chan error, *int, int))

// 	default:
// 		err = fmt.Errorf("Invalid CB function")
// 	}
// 	return err
// }

// func (ob *StructExample) Execute(thread_id int) {
// 	if ob.callback_type1 != nil {
// 		f := *ob.callback_type1
// 		f(ob.GetErrChan(), ob.GetObjectIdInt(), thread_id)
// 	}
// 	return
// }

// func (ob *StructExample) GetErrChan() chan error {
// 	return ob.sErr
// }

// func (ob *StructExample) SetErrChan(errChan chan error) {
// 	ob.sErr = errChan
// 	return
// }

// func (ob *StructExample) GetObjectIdStr() *string {
// 	return ob.objId_str
// }

// func (ob *StructExample) SetObjectIdStr(obid *string) {
// 	ob.objId_str = obid
// 	return
// }

// func (ob *StructExample) GetObjectIdInt() *int {
// 	return ob.objId
// }

// func (ob *StructExample) SetObjectIdInt(obid *int) {
// 	ob.objId = obid
// 	return
// }

// func (ob *StructExample) AllocateThreadPool(tp *thread_pool.ThreadPool) {
// 	ob.threadPool = tp
// 	return
// }

// func (ob *StructExample) GetThreadPool() *thread_pool.ThreadPool {
// 	return ob.threadPool
// }

// func randomVal(upperLimit, lowerLimit int) int {
// 	i := rand.Intn(upperLimit-lowerLimit) + lowerLimit
// 	return i
// }

// func InvokeCallback(ii int, cErr chan error, tp *thread_pool.ThreadPool) {
// 	id := ii
// 	sObj := StructExamplePtr()
// 	sObj.SetErrChan(cErr)
// 	sObj.SetObjectIdInt(&id)
// 	sObj.AllocateThreadPool(tp)
// 	callback := func(err chan error, id_ *int, thread_id int) {
// 		log.Println("Assigned thread id: ", thread_id, " to object id: ", *id_)
// 		if err == nil {
// 			log.Fatalf("Error channel is not declared.")
// 		}
// 		if id_ == nil {
// 			log.Fatalf("Int object id is nil.")
// 		}
// 		if tp == nil {
// 			log.Fatalf("Thread pool is nil.")
// 		}
// 		wait := randomVal(4, 1)
// 		log.Println("Reporting from callback function with id: ", *id_, " Waiting for ", wait, " seconds.")
// 		time.Sleep(time.Duration(wait) * time.Second)

// 		tp.FreeThread(thread_id)
// 		err <- fmt.Errorf("Done with this struct id: %d", *id_)
// 	}

// 	err := sObj.SetCallBack(&callback)
// 	if err != nil {
// 		fmt.Println("SetCallBack failed with err: ", err)
// 		return
// 	}

// 	tp.Push(sObj)
// }

// // ----------------------------------------------------------------

// func Fake_main() {

// 	UpperBound := 10000

// 	cErr := make(chan error, UpperBound)
// 	defer close(cErr)

// 	threadpool := thread_pool.NewThreadPool(1000) /* Thread pool initialisation with worker threads*/
// 	threadpool.SetInFlightRequestThreshold(2000)
// 	threadpool.RunSchedulerInBackground() /* Kick start the scheduler in background*/

// 	for ii := 1; ii <= UpperBound; ii++ {
// 		InvokeCallback(ii, cErr, threadpool)
// 	}

// 	for ii := 1; ii <= UpperBound; ii++ {
// 		c := <-cErr
// 		log.Println(ii, " Waiting for object to end: ", c)
// 	}
// 	return
// }

// // ----------------------------------------------------------------
