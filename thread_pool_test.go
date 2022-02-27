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
	"testing"
)

type testStruct struct {
	input_data int
	expected_data int

	t *testing.T
	cb *func(chan error)
	errChan chan error
}

func (tS *testStruct) Execute(thread_id int) {
	tS.t.Log("Worker thread id: ", thread_id)
	cb := *tS.cb
	defer cb(tS.errChan)

	output := tS.checkDataValues()
	if !output {
		tS.t.Fatalf("Data Corruption in the thread execution. Expected %d, Got %d", tS.expected_data, tS.input_data)
	}
}

func (tS *testStruct) setExpectedData(expected_data int) {
	tS.expected_data = expected_data
}

func testStructConstructor(data int, errChan chan error, t *testing.T) *testStruct {
	object := new(testStruct)
	object.input_data = data
	object.t = t
	object.errChan = errChan
	return object
}

func (tS *testStruct) checkDataValues() bool {
	 if tS.expected_data == tS.input_data {
		 return true
	 }
	 return false
}

func (tS *testStruct) setCallback(f *func(chan error)) {
	tS.cb = f
}

func TestUnitExection(t *testing.T){

	errChan := make(chan error, 1)
	threadpool := NewThreadPool(1)
	threadpool.SetInFlightRequestThreshold(1)
	threadpool.RunSchedulerInBackground()

	sample_data :=  testStructConstructor(10, errChan, t)
	sample_data.setExpectedData(10)
	cb := func(errChan chan error) {
		errChan <- nil
	}

	sample_data.setCallback(&cb)
	threadpool.Push(sample_data)

	<-errChan
}

func TestScaleExection(t *testing.T){

	numObjects := 100
	errChan := make(chan error, numObjects)
	threadpool := NewThreadPool(100)
	threadpool.SetInFlightRequestThreshold(100)
	threadpool.RunSchedulerInBackground()

	for ii := 1; ii <= numObjects ; ii++ {
		sample_data :=  testStructConstructor(ii, errChan, t)
		sample_data.setExpectedData(ii)
		cb := func(errChan chan error) {
			errChan <- nil
		}

		sample_data.setCallback(&cb)
		threadpool.Push(sample_data)
	}

	for ii := 1; ii <= numObjects ; ii++ {
		<- errChan
	}

}

