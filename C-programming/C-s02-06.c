#include <stdio.h>
/* #include <stdlib.h> */
/* #include <stdbool.h> */

/* void wypiszTabliceInt(int arrayInt[3]) { */
/*   printf("Tablica: [%d, %d, %d].\n", arrayInt[0], arrayInt[1], arrayInt[2]); */
/* } */

/* void wypiszTabliceInt(int arrayInt[], int arraySize) { */
/*   int i = 0; */

/*   printf("[%d, ", arrayInt[0]); */

/*   for (i = 1; i < (arraySize - 1); i++) { */
/*     printf("%d, ", arrayInt[i]); */
/*   } */

/*   printf("%d].\n", arrayInt[arraySize - 1]); */
/* } */

/* void wypiszZerowyElementInt(int arrayInt[], int size) { */
/*   printf("Zerowy element tablicy: %d.\n", arrayInt[0]); */
/* } */

/* void ampersandNaItem(int& intVar) { */
/*   printf("Działa.\n"); */
/* } */

/* double* podwojTabliceDouble(double arrayDouble[], int arraySize) { */
/*   int i = 0; */
/*   double* arrayReturn = (double*)malloc(arraySize); */

/*   for (i = 0; i < arraySize; i++) { */
/*     arrayReturn[i] = 2.0 * arrayDouble[i]; */
/*   } */


/*   return arrayReturn; */
/* } */

int main() {
  /* int arrayInt1[3] = {1, 2, 3}; */
  /* int arrayInt2[5] = {-1, -2, -3, -4, -5}; */

  /* double arrayDouble[4] = {0.5, 0.75, 1.0, 1.25}; */

  /* bool boolVar = false; */

  /* int intVar = 3; */
  /* int *ptrVar = &intVar; */

  /* int intVar = 3 + 4 * true; */

  /* printf("boolVar = %d.\n", boolVar); */
  /* printf("intVar = %d.\n", intVar); */

  /* wypiszZerowyElementInt(arrayInt1, 3); */
  /* wypiszZerowyElementInt(arrayDouble, 4); */


  /* int intVar = 3; */
  /* arrayInt1 = &intVar; */

  /* wypiszTabliceInt(arrayInt1); */
  /* printf("Wypisujemy tablice arrayInt1.\n"); */
  /* wypiszTabliceInt(arrayInt1, 3); */
  /* printf("Wypisujemy tablice arrayInt2.\n"); */
  /* wypiszTabliceInt(arrayInt2, 5); */
  /* wypiszTabliceInt(arrayDouble, 8); */

  /* ampersandNaItem(intVar); */

  /* printf("*arrayInt1 = %d.\n", *arrayInt1); */

  /* printf("ptrVar: %ls.\n", ptrVar); */
  /* printf("ptrVar: %ls.\n", ptrVar); */
  /* printf("&intVar: %ls.\n", &intVar); */

  /* double arrayDouble[3] = {0.5, 1.0, 1.5}; */
  /* double* arrayValue; */

  /* arrayValue = podwojTabliceDouble(arrayDouble, 3); */

  /* printf("arrayValue[0] = %.3f.\n", arrayValue[0]); */

  int *ptrInt, intVar;
  int intVar1 = 0;

  /* intVar = &intVar1; */
  ptrInt = &intVar1;

  /* printf("*intVar = %d.\n", *intVar); */
  printf("*ptrVar = %d.\n", *ptrInt);




  return 0;
}
