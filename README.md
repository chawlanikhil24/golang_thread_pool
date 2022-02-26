# golang_thread_pool
This Golang package, enhances the experience of asynchronous programming in golang programs,by implementing a thread_pool concept, with staged event driven architecture based programming.

Now, to walk you through **staged event driven architecture**, according to [wikipedia](https://en.wikipedia.org/wiki/Staged_event-driven_architecture), it is an approach in event driven programming which decomposes the programing in to connected stages. By applying an **admission control** to the worker queue, it helps to maintain a steady load on hardware.

Admission control is a pre-validation check that ensures the safer check before submitting the request into the event queue.

## How to use this library ?

* import the libary.
* Create the threadpool object using contructor, pass the worker thread count as parameter.
* Set the inflight request count, **Default: 5**.
* Run the scheduler.
* Form any object of threadpool.AsyncProgInterface type, which implements only 1 function , Execute(thread_id int).

Here is the code,

`threadpool := thread_pool.NewThreadPool(1000) /* Thread pool initialisation with worker threads*/
threadpool.SetInFlightRequestThreshold(2000) /* Set how much requests should be submitted in queue at max */
threadpool.RunSchedulerInBackground() /* Kick start the scheduler in background*/
threadpool.Push(/* thread_pool.AsyncProgInterface */) /*In easy words, any struct that implements an "Execute(thread_ID int) Method"*/`

Feel free to contribute and help to improve.