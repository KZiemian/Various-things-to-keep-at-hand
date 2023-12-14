#include <stdio.h>
#include <stdbool.h>

bool searchFor(int id) {
  int indexMid = 0;
  int elemMid = 0;
  int start = 0;
  int end = 15;
  bool found = false;
  int numberOfLoops = 0;

  int arrayInt[16] = {1, 3, 6, 8, 9, 12, 13, 16, 17, 22, 27, 30, 33, 40,
    45, 50};


  while((start <= end) && (found != true) && (numberOfLoops < 100)) {
    numberOfLoops++;

    indexMid = (start + end) / 2;
    elemMid = arrayInt[indexMid];

    if (elemMid == id) {
      printf("numberOfLoops = %d.\n", numberOfLoops);



      return true;
    } else if (elemMid < id) {
      start = indexMid + 1;
    } else {
      end = indexMid - 1;
    }
  }

  printf("numberOfLoops = %d.\n", numberOfLoops);



  return false;
}

int main() {
  /* int intVar = 0; */
  /* int *ptrInt = &intVar; */

  /* *ptrInt = 2; */

  /* printf("ptrInt = %p.\n", ptrInt); */
  /* printf("*ptrInt = %d.\n", *ptrInt); */

  /* ptrInt = 16; */

  /* printf("ptrInt = %p.\n", ptrInt); */
  /* printf("*ptrInt = %d.\n", *ptrInt); */

  /* int intVar = 0; */
  /* double doubleVar = 0.5; */

  /* int *ptrInt = &intVar; */
  /* double *ptrDouble = &doubleVar; */

  /* printf("ptrInt = %p.\n", ptrInt); */
  /* printf("*ptrInt = %d.\n", *ptrInt); */
  /* printf("ptrDouble = %p.\n", ptrDouble); */
  /* printf("*ptrDouble = %.3f.\n", *ptrDouble); */

  /* ptrInt = &doubleVar; */

  /* printf("ptrInt = %p.\n", ptrInt); */
  /* printf("*ptrInt = %d.\n", *ptrInt); */
  /* printf("doubleVar = %.3f.\n", doubleVar); */

  /* int arrayInt[5] = {1, 2, 3, 4, 5}; */
  /* int *ptrInt = arrayInt; */

  /* printf("&arrayInt[0] = %p.\n", &arrayInt[0]); */
  /* printf("arrayInt = %p.\n", arrayInt); */
  /* printf("*arrayInt = %d.\n", *arrayInt); */
  /* printf("*(arrayInt + 2) = %d.\n", *(arrayInt + 2)); */
  /* printf("*arrayInt = %d.\n", *arrayInt); */
  /* arrayInt += 3; */
  /* printf("*arrayInt = %d.\n", *arrayInt); */

  /* printf("arrayInt = %p.\n", arrayInt); */

  /* ptrInt++; */

  /* printf("ptrInt = %p.\n", ptrInt); */
  /* printf("*ptrInt = %d.\n", *ptrInt); */
  /* ptrInt += 2; */

  /* printf("*ptrInt = %d.\n", *ptrInt); */
  /* bool boolVar1 = true; */
  /* bool boolVar2 = false; */
  /* int intVar = 0; */

  /* intVar = 7 * boolVar1 + 3 * boolVar2; */

  /* printf("intVar = %d.\n", intVar); */

  printf("searchFor(3) = %d.\n", searchFor(3));
  printf("searchFor(2) = %d.\n", searchFor(2));
  printf("searchFor(40) = %d.\n", searchFor(40));
  printf("searchFor(43) = %d.\n", searchFor(43));




  return 0;
}
