# Log Structured Merge Tree

## Design

* What type of format file?
  * Binary
* How to do structure the Index?

| keylen  | valuelen | key            | value          |
|---------|----------|----------------|----------------|
| uint16b | uint16b  | String (utf-8) | String (utf-8) |

* Block?
  * Atomic Read -> We read a chund of data, this should define the block size. So when we read a block we read it as a one operation.
  * Why do we need a block?
    * Index gives me a block
    * Load bloack into memory and search there for a key.

* First iteration
  * Rows of K/V pairs encoded in binary with a prefix for the number of elements. Eg: Hello World
    * 04 00 05 05 05 00 68 65 6c 6c 6f 77 6f 72 6c 64 ....
    * To find a key we need to go through the whole file until we find it
* Second iteration
  * With index -> To help locate a key in the data section

## Notes

* Serialization performance?
  * Measure using bytes/second (e.g. for a fixed dataset, or benchmark suite of datasets)
* Why do we need to define blocks?
  * To chunk data and be able to load only subsets of it into memory
* Design
  * Text format file -> Has some downsides, because it is diffictult to go direct to a section of a file
  * Binary format -> Easier to read from an offset.
* Optimization:
  * Bloom Filter
    * We could have a filter for the whole file
    * Check key in filter, if not there end
    * If present use index to look up block
    * Load block into memory and search for key
  * Compress Blocks w/ zip or whatever as the data is potential very compressable
  * 