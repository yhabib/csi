#include "vec.h"

data_t dotproduct(vec_ptr u, vec_ptr v)
{
   // Multiple accumulators -> parallelism
   data_t sum = 0, sum2 = 0;
   // Moves this out from the loop -> to avoid calculating it with each iteration
   int length = vec_length(u);
   // Pointer to the data so I cana access it directly
   data_t *data_u = get_vec_start(u);
   data_t *data_v = get_vec_start(v);

   for (long i = 0; i < length; i += 2)
   { // we can assume both vectors are same length
      sum += data_u[i] * data_v[i];
      sum2 += data_u[i+1] * data_v[i+1];
   }
   return sum + sum2;
}
