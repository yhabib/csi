// #include "darts.h"---------
#define SCORE_INNER_CIRCLE 10
#define SCORE_MIDDLE_CIRCLE 5
#define SCORE_OUTER_CIRCLE 1
#define SCORE_OUT_OF_BOUNDS 0

typedef struct
{
  float x;
  float y;
} coordinate_t;

int score(coordinate_t);
// ---------------------------------------

// Notes: By using the square of the distance no need of root
// Exercism: Define function and structs there

int score(coordinate_t pos)
{
  float d = pos.x * pos.x + pos.y * pos.y;
  if (d <= 1)
    return SCORE_INNER_CIRCLE;
  else if (d <= 25)
    return SCORE_MIDDLE_CIRCLE;
  else if (d <= 100)
    return SCORE_OUTER_CIRCLE;
  return SCORE_OUT_OF_BOUNDS;
}
