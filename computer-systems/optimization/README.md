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


## Resources
- [Instruments](https://stackoverflow.com/questions/11445619/profiling-c-on-mac-os-x)