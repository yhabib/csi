# Bloom Filter

## Run project

```sh
go run *.go
```

## References

* [Tutorial](https://llimllib.github.io/bloomfilter-tutorial/)
* [Doc](https://michaelnielsen.org/ddi/why-bloom-filters-work-the-way-they-do/)
* [Wiki](https://en.wikipedia.org/wiki/Bloom_filter#Examples)

## Class Notes

* What problem is a Bloom filter trying to solve?
  * Existence in a set, "is an item included in a set"
  * Fast and memory efficient
    * Faster?
      * Constant time lookups
      * But could be slower than a HashSet
    * Memory efficient
      * One bit per element to save the data
  * But it is "crappier" than a "set" you can get false positives

* Alternatives of Bloom filters?
  * Hash Sets
  * Tree, linked list, trie
  * Array?

* How does a Bloom filter work?
  * Bit array
  * Key, run it through hash functions
  * "Turn on" those bits

* What kind of error can occur?
  * A false positive
  * No false negative -> Deterministic hash functions

* How can we have "n different hash functions" ?
  * "seed" on hash function
    * different values of seed  -> different outputs
  * prefix key to have different values
    * h("1 alice")
    * h("2 alice")
    * h("3 alice")

* Where does a Bloom filter fit in?
  * Used as in memory DS
  * Placed before a DB to avoid doing lookups
  * If accurancy is high ~98% it means that will access the DB only in 2% of the cases

* Which usage patterns should we tunr on /turn off Bloom filters?
  * If many keys might be present: can skip a LOT of disks reads(98%)
  * If most of our keys are present: it's useless
  * Range Scans: it's useless

* Question:
  * Bloom filter data is stored on disk
  * At some point, yuo are probably going to load that into memory, WHEND do you do that?
    * At boot time?
      * How much memory would this potentially take up?
        * 1 bit for each entry?
    * Last recently used?
      * first load bloom filter into memory
      * check and if needed go back again to memory for the value
  * What kind of knobs can we tune, what kind of tradeoffs can we make(time, space, accuracy, etc...)
    * How larger (larger = more accuracy)
    * Number of hash functions (accurancy vs. time)

* Bloom filter applications
  * Bitcoin

* Additional resources:
  * [Postgress](https://www.percona.com/blog/2019/06/14/bloom-indexes-in-postgresql/)
