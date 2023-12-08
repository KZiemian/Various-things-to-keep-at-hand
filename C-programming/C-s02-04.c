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

/* int main() { */
/*   float testVar1 = 262144.0; */
/*   float testVar2 = testVar1 + 0.01; */

/*   printf("testVar1: %f.\n", testVar1); */
/*   printf("testVar2: %f.\n", testVar1); */

/*   printf("testVar1 == testVar2: %d.\n", testVar1 == testVar2); */





/*   return 0; */
/* } */

/* int printf(const char* format, ...); */

int main() {
  /* printf("Hello, World!\n"); */

  /* unsigned long long unsignedLongLongVar = 0; */
  /* unsigned long long int unsignedLongLongIntVar = 0; */

  /* printf("sizeof(unsignedLongLongVar) = %ld.\n", sizeof(unsignedLongLongVar)); */
  /* printf("sizeof(unsignedLongLongIntVar) = %ld.\n", */
  /* 	 sizeof(unsignedLongLongIntVar)); */

  /* float floatVar = 0.0; */
  /* double doubleVar = 0.0; */
  /* long double longDoubleVar = 0.0; */

  /* printf("sizeof(floatVar) = %ld.\n", sizeof(floatVar)); */
  /* printf("sizeof(doubleVar) = %ld.\n", sizeof(doubleVar)); */
  /* printf("sizeof(longDoubleVar) = %ld.\n", sizeof(longDoubleVar)); */

  /* const int i = 5; */
  /* int arrayVar[2] = {1, 2, 3}; */

  /* arrayVar[0] = 1; */
  /* arrayVar[1] = 2; */
  /* arrayVar[2] = 3; */
  /* arrayVar[3] = 4; */
  /* arrayVar[4] = 5; */


  /* printf("arrayVar[1] = %d.\n", arrayVar[1]); */

  /* for (int i = 0; i < 4; i++) { */
  /*   printf("arrayVar[%d] = %d.\n", i, arrayVar[i]); */
  /* } */

  int arrayInt[6] = {1, 2, 3, 4, 5, 6};

  int i = 0;
  int intVar = 0;

  for (i = 0; i < 6; i++) {
    printf("arrayInt[%d] = %d.\n", i, arrayInt[i]);
  }

  intVar = arrayInt[3];
  printf("intVar = %d.\n", intVar);

  intVar = *(arrayInt + 2);

  printf("intVar = %d.\n", intVar);

  *(arrayInt + 2) = 8;

  for (i = 0; i < 6; i++) {
    printf("arrayInt[%d] = %d.\n", i, arrayInt[i]);
  }





  return 0;
}
