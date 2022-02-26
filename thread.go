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


package thread_pool

// ----------------------------------------------------------------
type ThreadInfo struct {
	busy bool
	id   int

	upperBoundForId int
	lowerBoundForId int
}

func NewThread() *ThreadInfo {
	th := new(ThreadInfo)
	th.upperBoundForId = 999999
	th.lowerBoundForId = 999
	th.id = randomVal(th.upperBoundForId, th.lowerBoundForId)
	th.busy = false
	return th
}

func (th *ThreadInfo) IsBusy() bool {
	return th.busy
}

func (th *ThreadInfo) SetBusy() {
	th.busy = true
	return
}

func (th *ThreadInfo) SetFree() {
	th.busy = false
	return
}

func (th *ThreadInfo) GetId() int {
	return th.id
}

// ----------------------------------------------------------------
