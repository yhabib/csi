# Roaring Bitmaps

## Bitmap

A bitmap is a mapping from one system such as integers to bits. It is also known as bitmap index or a bit array.

The memory is divided into units for bitmap. These units may range from a few bytes to several kilobytes. Each memory unit is associated with a bit in the bitmap. If the unit is occupied, the bit is 1 and if it is empty, the bit is zero.

### API

* set(uint32 x): It sets the "nth" bit to true, where nth is the provided input. Eg: x = 32 -> it will set the bit 32 to true
* get(uint32 x): It checks the presence of the value in the structure. 
* union(uncompressedBitmap)uncompressedBitmap : 
* intersect(uncompressedBitmap)uncompressedBitmap : 

## Notes

* Bitmap use cases
  * Bloom filters
  * Inverted index(in any information retrieval system)
    * Search index where you can search for documents containing some words
      * A bitmap per word, where each index represents a document
  * Bitmap heap scan in Postgres
    * how do you use two index?
      * e.g. movies
        * index on title
        * index on year
        * want to do a query like title > ... and title < ... and year > ... and year < ...
          * use one index,, get some of page ids
          * use second index, get set of page ids
          * intersect them
* Alternatives to implement sets of integers:
  * bitmap
  * hastable(key integer, value = some canonical value to indicate presence, True, etc)
    * Java HashSet, probaly most languages do this as their default "Set" implementation
  * array: put everything in it
    * sort or not to allow binary search
* Different enconding schemes:
  * example input: [0, 0, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 0]
  * Basic RLE
    * [(4, 0), (1, ,1), (1, 0), (1,0), ...]
    * [(1, 1), (1, ,0), (1, 1), (1,0), ...] - worst case
  * Word-aligned hybrid:
    * we **can't** operate directly in the copmressed version
    * for each word:
      * 1st bit tells you if the word is "filled" or not
      * if not:
        * just the literal bits for input
      * if filled:
        * 2nd bit represents the value(either 0 or 1)
        * remaining bits represents the lenght of the run
      * example:
        * [   0, 0, 0, 0, 1, 0, 1,     1] [1, 1, 0, 1, 0, 0, 0,      1, 1]
        * [0, 0, 0, 0, 0, 1, 0, 1] [0, 1,  1, 1, 0, 1, 0, 0, 0]  [0, 1, 1]
  * ~Concise~
  * Roaring
    * we **can** operate directly in the copmressed version
    * Key ideas
      * Nested List
        * Array
        * Bitmap
        * Run length
    * Basic operations
      * Get
        * Figure out wich bucket is in, top level
          * Get most significant 16bits (that is your index)
          * Binary search the array of containers to look for that value
        * Based on type of bucket:
          * If array, binary search
          * If bitmap, find right word
      * Set (mark bit as on)
        * Top Level: binary search to find right container
        * Second level: if it's an array, binary search and do linear insert
          * Possibly convert it to a bitmap
          * If it's a bitmap just ad it
      * What kind of performance can you expect?
        * Get: worst, binary on  an array of 4096, 2^12, and binary search at the top level to figure out right container
        * Set: worst, binary search + linear at tope level,, binary search + linear insert or convert to bitmap
    * Implementation details:
      * Generate key:
        * If you want to add 123874612 to the set, what is they key?
          * 123874612 >> 16 = 1890
          * 123874612 & ((1 << 16) - 1) = 11572
      * When do you convert from array to bitmap?
        * Paper says as soon as you hit 4097 elements -> bitmap