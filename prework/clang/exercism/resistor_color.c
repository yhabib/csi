#ifndef RESISTOR_COLOR_H
#define RESISTOR_COLOR_H

typedef enum
{
  BLACK,
  BROWN,
  RED,
  ORANGE,
  YELLOW,
  GREEN,
  BLUE,
  VIOLET,
  GREY,
  WHITE
} resistor_band_t;

resistor_band_t color_code(resistor_band_t);
const resistor_band_t *colors(void);

#endif

// #include "resistor_color.h"

// Notes:
//  static is a storage specifier: all_colors is scoped to this file and will not been outside
//    if used inside a function the variable will be kept in memory even though the function returns
//  const is a type qualifier: value of variable will not change after initialization
//  void as type parameter if no real parameters

resistor_band_t color_code(resistor_band_t color)
{
  return color;
}

const resistor_band_t *colors(void)
{
  static resistor_band_t colors[10] = {BLACK, BROWN, RED, ORANGE, YELLOW, GREEN, BLUE, VIOLET, GREY, WHITE};
  return colors;
}