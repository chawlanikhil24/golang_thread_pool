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
	"math/rand"
	"time"
)

// A utility function to generate random values between the provided upper bounds
func randomVal(upperLimit, lowerLimit int) int {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(upperLimit-lowerLimit) + lowerLimit
	return i
}
