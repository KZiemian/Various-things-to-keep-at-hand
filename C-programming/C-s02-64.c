#include <stdio.h>

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

  double x = 4.81;

  while (x <= 5.0) {
    printf("f(%.3f) = %.4f.\n", x, 0.000001 * x * x * x * x * x * x +
	   0.001 * x * x * x * x * x - 0.01 * x * x * x * x -
	   0.1 * x * x * x  + x * x + x - 10.0);

    x += 0.01;
  }






  return 0;
}
