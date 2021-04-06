# Optimization

## Pagecount

Optimizations:
1. Replace division with right shift
2. Unroll the for loop with the three calculations at the same time

### Stretch goals

1.
```
PAGE_NUM = ((MEM >> PS) >> 2) + ((MEM >> PS) >> 4) + 1
PS = ((PAGE_SIZE >> 4) + (PAGE_SIZE >> 2) + 1)
```


## Vectod Dot Product

Optimizations:
1. Move `vec_length` outside of for loop O(n2)-> O(n)
2. Removes procedure calls and instead uses pointer to the data.
3. Dont' use memory for intermediate values
4. Loop rolling(2x1)
5. Multiple accumulators

## Resources
- [Instruments](https://stackoverflow.com/questions/11445619/profiling-c-on-mac-os-x)
- [Godbolt](https://godbolt.org/)