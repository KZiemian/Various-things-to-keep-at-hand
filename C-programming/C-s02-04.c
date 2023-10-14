#include<stdio.h>

/* int main() { */
/*   float meters = 0.0; */
/*   int iterations = 1000000000; */

/*   int i = 0; */

/*   for (i = 0; i < iterations; i++) { */
/*     meters += 0.01; */
/*   } */

/*   printf("Expected: %f km.\n", 0.01 * iterations / 1000.0); */
/*   printf("Got: %f km.\n", meters / 1000.0); */





/*   return 0; */
/* } */

int main() {
  float testVar1 = 262144.0;
  float testVar2 = testVar1 + 0.01;

  printf("testVar1: %f.\n", testVar1);
  printf("testVar2: %f.\n", testVar1);

  printf("testVar1 == testVar2: %d.\n", testVar1 == testVar2);



  return 0;
}
