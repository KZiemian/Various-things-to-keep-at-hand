#include <stdio.h>

int main() {
  /* int celsius = 0; */
  /* int fahr = 0; */

  /* int lower = -17; */
  /* int upper = 150; */
  /* int step = 11; */

  /* celsius = lower; */

  /* while (celsius <= upper) { */
  /*   printf("Celsius: %d;    Fahrenhit: %d.\n", celsius, 9*celsius / 5 + 32); */

  /*   celsius += step; */
  /* } */

  /* double celsius = 0.0; */
  /* double fahr = 0.0; */

  /* double lower = -17.0; */
  /* double upper = 150.0; */
  /* double step = 11.0; */

  /* celsius = lower; */

  /* while (celsius <= upper) { */
  /*   printf("Celsius: %.1f;    Fahrenhit: %.3f.\n", celsius, (9.0/5.0) * */
  /* 	   celsius + 32.0); */

  /*   celsius += step; */
  /* } */

  /* double celsius = 0.0; */
  /* double fahr = 0.0; */
  /* double valueCalculated = 0.0; */

  /* double lower = 0.0; */
  /* double upper = 300.0; */
  /* double step = 20.0; */

  /* fahr = lower; */

  /* while (fahr <= 300) { */
  /*   valueCalculated = 5.0 * (fahr - 32.0) / 9.0; */
  /*   /\* printf("fahr = %.1f;    valueCalculated = %.3f.\n", fahr, *\/ */
  /*   /\* 	   valueCalculated); *\/ */
  /*   valueCalculated = 9.0 * valueCalculated / 5.0 + 32.0; */
  /*   printf("fahr = %.1f;    valueCalculated = %.3f.\n", fahr, */
  /* 	   valueCalculated); */

  /*   fahr += step; */
  /* } */

  int i = 1;
  int factorialValue = 1;


  for (i = 1; i <= 20; i++) {
    factorialValue *= i;

    printf("i = %d;   factorialValue = %d.\n", i, factorialValue);
  }





  return 0;
}
