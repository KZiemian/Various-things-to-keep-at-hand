#include <stdio.h>
#include <math.h>

int main() {
  double x = 4.01;

  while (x <= 4.1) {
    printf("f(%.3f) = %.4f.\n", x, x * x * x * x - x * x * x - x * x + x);

    x += 0.01;
  }





  return 0;
}
