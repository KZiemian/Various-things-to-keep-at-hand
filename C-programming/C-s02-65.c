#include <stdio.h>
#include <math.h>

int main() {
  double x = 9.8;

  while (x <= 10.0) {
    /* printf("f(%.3f) = %.4f.\n", x, cos(x) * exp(-x)); */
    /* printf("f(%.3f) = %.4f.\n", x, x * exp(-(x - 2.0))); */
    /* printf("f(%.3f) = %.4f.\n", x, x * x * exp(-x)); */
    /* printf("f(%.3f) = %.4f.\n", x, x * x * exp(-(x - 1.0))); */
    /* printf("f(%.3f) = %.4f.\n", x, pow(4.0, x)); */
    printf("f(%.3f) = %.4f.\n", x, pow(5.0, x));

    x += 0.01;
  }

  /* printf("1/8 == %.5f.\n", 1.0/8.0); */

  /* double result = -5.0 * sqrt(2.0) / 2.0 + 1.5 + 2.0 * sqrt(3.0); */
  /* double result = -5.0 * sqrt(2.0) / 2.0 + 3.0 * sqrt(3.0) / 2.0 + 2.0; */

  /* printf("result = %.3f.\n", result); */

  /* printf("1.42 / 1.06 = %.3f.\n", 1.42 / 1.06); */

  /* double x = 12.4; */

  /* while (x <= 12.6) { */
  /*   printf("f(%.3f) = %.4f.\n", x, cos(x) - x); */

  /*   x += 0.01; */
  /* } */





  return 0;
}

/* 59:10 https://www.youtube.com/watch?v=a6GTCDRToXo */
