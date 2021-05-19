# Garbage Collection

## Commands

```bash
GODEBUG=gctrace=1 go run <path_to_program>
GODEBUG=gctrace=1,gcpacertrace=1 go run <path_to_program>
```


## Class

### Objectives
* How a "mark and sweep" garbage collector works
* OPeration and performance overhead of Go's concurrent collector
* Why the "ballast" trick from Twitch article helped performance

### Notes
* Types of Garbage Collectors
  * Mark and sweep
  * Reference counting(Python)
  * Stop and Copy 
    * Java?
  * Generational garbage collection
    * Takes andvantage of "lindy" effect
    * JavaScript??
* Reference book "The Garbage Collection Handbook"
  
* Simple Mark&Sweep 
  * Mark phase
    * Start at "root" node(s)
      * Any node that is not in the heap: 
        * Any pointers in some goroutine's stack
          * How do you find all the stacks in all the goroutines?
            * `g struct` tells us
        * Global variables. static data, registers, ...
    * Traverse all the nodes as a *graph* traversal
      * any algorithm that visits all nodes would do the job so both BFS and DFS
        * GO provably does some optimizations here  
      * while traversing mark each visited node with a bit 
    * When do you run this mark phase? ideas?
      * Timer to run GC every Xms
      * Inc counter after each allocation and run it after N 
      * Metrics on the Heap, how much it has been used.
      * Run it when running out of memory. if allocation fails then run it and try again
        ```c
        void * malloc(size) {
          // try to allocate
          if () {
            mark()
            sweep()
          }
          // try again, and if fails ask for more mem
          if (failed) {
          }
        }
        ```
  * Sweep phase:
    * Visits *every* node in the heap
      * if it is mark -> unmark it
      * if it is NOT mark -> free it(it is garbage)

* GO's collector:
  * Marking happens based on a Pacing algorithm
    * Based on the stress put into the Heap
      * Stress as number of allocations per unit of time
      * Prediction based on the GC Percentage(100% default)
  * Sweeping happens every time there is an allocation
  * Latency?
    * Goal: minimize "stop the world" time
    * Cost of "STW" pahses:
      * Setup phase ~10-30 microseconds
      * Termination phase ~60-90 microseconds
    * Cost of Concurrent phase:
      * 1/4 of your P(available goroutine)
      * Can also enlist other goroutines to help, "Mark Assist"
        * in case that there are too many allocations, then it marks other goroutines to assist
        * by doing this, all the marked ones stop allocating
    * Write Barrier
      * Go in this mode between the Setup and Termination phase -> Concurrency phase
      * Any pointer updates needs to do some extra work to let the GC know that some allocation is happening.
  * How can we ensure that the "ballast" doesn't actually cause an increase in the "Resident Set Size"?
    *Does Go support this? Falg in the future?
  * GC percentage flag, impact into pacing?
    * Counterfactual("start" of next mark phase)
      * GC finishes 10mb of heap is "live"
      * Target: 100% -> 20mb
      * Start next GC       
    * What really happens?
      * Pacer schedules next GC through predictions based on current pace, heap growth
      * Target: want to make sure there is at most 100% by the *end* of the next GC run
        * Something more happing there?
* Twitch solution with "ballast"
  * Instead of setting the GC percetage to 1000% they set initial Heap size to 10gb so a 100% increase would be 20GB
  * Too noisy, too unpredictable, could magnify small differences if we set it very high



