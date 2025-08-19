#include <stdio.h>
#include <math.h>

/* const double pi_const = 3.1415926535; */
/* const double e_const = 2.7182818284; */

int main() {
  /* double x = 5.3; */

  /* while (x <= 5.5) { */
  /*   printf("f(%.3f) = %.4f.\n", x, 0.001 * x * x * x * x - 0.01 * x * x * x - */
  /* 	   0.1 * x * x - x); */

  /*   x += 0.01; */
  /* } */

  /* printf("U_2N_3:\n"); */
  /* printf("476.0 / 518.0 = %.4f.\n", 476.0 / 518.0); */

  /* printf("PaCl_5:\n"); */
  /* printf("231.0 / 406.0 = %.4f.\n", 231.0 / 406.0); */

  /* printf("0.24 * pi = %.5f.\n", 0.24 * pi_const); */
  /* printf("1.0/e_const = %.5f.\n", 1.0/e_const); */

  /* printf("3.6 * 26 = %.4f.\n", 3.6 * 26); */
  /* printf("2.0 * 3.1415 * 0.0595 = %.4f.\n", 2.0 * 3.1415 * 0.0595); */

  /* double x = 20.0; */

  /* printf("f(%.3f) = %.4f.\n", x, 0.000001 * x * x * x * x * x * x + */
  /* 	 0.001 * x * x * x * x * x - 0.01 * x * x * x * x - */
  /* 	 0.1 * x * x * x  + x * x + x - 10.0); */

  /* while (x <= 20.0) { */
  /*   printf("f(%.3f) = %.4f.\n", x, 0.000001 * x * x * x * x * x * x + */
  /* 	   0.001 * x * x * x * x * x - 0.01 * x * x * x * x - */
  /* 	   0.1 * x * x * x  + x * x + x - 10.0); */

  /*   x += 0.01; */
  /* } */

  /* double x = 9.61; */

  /* while (x <= 10.0) { */
  /*   printf("f(%.3f) = %.4f.\n", x, x * x * x * x - 20.0 * x * x * x + */
  /* 	   33.75 * x * x - 235.0 * x + 89.0625); */

  /*   x += 0.01; */
  /* } */

  /* x^4 - 20x^3 + 33.75x^2 - 235x + 89.0625 */

  /* double x = 9.61; */

  /* while (x <= 10.0) { */
  /*   printf("f(%.3f) = %.4f.\n", x, x * x * x * x * x - */
  /* 	   19.222 * x * x * x * x + 69.903 * x * x * x - 10.34 * x * x - */
  /* 	   15.722 * x + 9.177); */

  /*   x += 0.01; */
  /* } */

  /* x = 1.3; */
  /* printf("f(%.3f) = %.4f.\n", x, 0.001 * x * x * x * x - */
  /* 	 0.001 * x * x * x - 0.1 * x * x - x); */

  /* double x = 0.0; */

  /* while (x <= 1.0) { */
  /*   printf("f(%.3f) = %.4f.\n", x, 4.0 * x * (1.0 - x)); */

  /*   x += 0.01; */
  /* } */

  /* x = 1.76; */
  /* printf("f(%.3f) = %.4f.\n", x, 0.001 * x * x * x * x - */
  /* 	 0.001 * x * x * x - 0.1 * x * x - x); */

  /* int intVar1 = 1; */
  /* int intvar1 = 2; */

  /* printf("intVar1: %d.\n", intVar1); */
  /* printf("intvar1: %d.\n", intvar1); */

  /* double x = 20.0; */

  /* printf("f(%.3f) = %.4f.\n", x, sin(x) / x); */

  double x = 25.0;

  while (x <= 25.2) {
    /* printf("f(%.3f) = %.4f.\n", x, x * exp(-x)); */
    /* printf("f(%.3f) = %.4f.\n", x, x * exp(-x + 1.0)); */
    printf("f(%.3f) = %.4f.\n", x, sin(x) / x);
    /* printf("f(%.3f) = %.4f.\n", x, acos(x)); */

    x += 0.01;
  }

  x = 25.2;

  printf("f(%.3f) = %.4f.\n", x, sin(x) / x);

  /* x = 1.0; */

  /* printf("f(%.3f) = %.4f.\n", x, acos(x)); */





  return 0;
}
