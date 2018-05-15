#include "lib.h"

#include <stdlib.h>
#include <stdio.h>
#include <math.h>


float poly(float x) {
  float mu = 10.0;
  float s;
  int j;
  float pol[100];

  for (j=0; j<100; j++) {
      pol[j] = mu = (mu + 2.0) / 2.0;
    }
  s = 0.0;
  for (j=0; j<100; j++) {
      s = x*s + pol[j];
    }
  return s;
}

