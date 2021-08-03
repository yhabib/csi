# Roaring Bitmaps

## Bitmap

A bitmap is a mapping from one system such as integers to bits. It is also known as bitmap index or a bit array.

The memory is divided into units for bitmap. These units may range from a few bytes to several kilobytes. Each memory unit is associated with a bit in the bitmap. If the unit is occupied, the bit is 1 and if it is empty, the bit is zero.

### API

* set(uint32 x): It sets the "nth" bit to true, where nth is the provided input. Eg: x = 32 -> it will set the bit 32 to true
* get(uint32 x): It checks the presence of the value in the structure. 