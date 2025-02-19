#include <stdio.h>
#include <math.h>

int main() {
  double x = 9.51;
  double y = 0.25 * x * x * x * x - x * x * x + x - 5.0 * cos(x) + 5.0;


  printf("f(%.3f) = %.4f.\n", x, y);

  x = 9.52;

  while (x <= 10.0) {
    y = 0.25 * x * x * x * x - x * x * x + x - 5.0 * cos(x) + 5.0;

    printf("f(%.3f) = %.4f.\n", x, y);

    x += 0.01;
  }





  return 0;
}
