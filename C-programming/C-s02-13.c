#include <stdio.h>

int main() {
  /* printf("Hello World!"); */
  /* double doubleVar = 2.0 * 0.7071067811; */

  /* printf("doubleVar = %.3f.\n", doubleVar); */

  /* doubleVar *= doubleVar; */

  /* printf("doubleVar = %.3f.\n", doubleVar); */

  /* double doubleVar = 3.0 * 0.4714045207; */

  /* printf("doubleVar = %.3f.\n", doubleVar); */

  /* doubleVar *= doubleVar; */

  /* printf("doubleVar = %.3f.\n", doubleVar); */

  /* int i = 0; */
  /* int factorialOf10 = 1; */

  /* for (i = 1; i <= 10; i++) { */
  /*   factorialOf10 *= i; */

  /*   printf("i = %d, factorialOf10 = %d.\n", i, factorialOf10); */
  /* } */

  /* int i = 0; */
  /* int sum = 0; */

  /* for (i = 1; i <= 20; i++) { */
  /*   sum += i; */

  /*   printf("i = %d, sum = %d.\n", i, sum); */
  /* } */

  /* int i = 0; */

  /* for (i = 1; i < 20; i++) { */
  /*   printf("%d^2 = %d.\n", i, i * i); */
  /* } */

  /* int i = 0; */

  /* for (i = 1; i <= 20; i++) { */
  /*   printf("%d^3 = %d.\n", i, i * i * i); */
  /* } */

  /* int i = 0; */

  /* for (i = 1; i <= 20; i++) { */
  /*   printf("%d^4 = %d.\n", i, i * i * i * i); */
  /* } */

  /* int i = 0; */
  /* int sum = 0; */

  /* for (i = 1; i <= 20; i++) { */
  /*   sum += i * i; */

  /*   printf("i = %d, sum = %d.\n", i, sum); */
  /* } */

  /* for (i = 1; i <= 20; i++) { */
  /*   sum += i * i * i; */

  /*   printf("i = %d, sum = %d.\n", i, sum); */
  /* } */

  /* int celsius = 0; */
  /* int fahr = 0; */

  /* int lower = 0; */
  /* int upper = 300; */
  /* int step = 20; */

  /* fahr = lower; */

  /* while (fahr <= upper) { */
    /* printf("Celsius: %d;    Fahrenheit: %d.\n", 5 * (fahr - 32) / 9, fahr); */
  /*   printf("Fahrenheit: %d;    Celsius: %d.\n", fahr, 5 * (fahr - 32) / 9); */

  /*   fahr += step; */
  /* } */

  double celsius = 0.0;
  double fahr = 0.0;

  double lower = 0.0;
  double upper = 300.0;
  double step = 20.0;

  fahr = lower;

  while (fahr <= upper) {
    printf("Fahrenheit: %.1f;   Celsius: %.5f.\n", fahr,
	   (5.0/9.0) * (fahr - 32.0));

    fahr += step;
  }





  return 0;
}
