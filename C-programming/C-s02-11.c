#include <stdio.h>

/* int main() { */
/*   int intVar = 10; */
/*   int *ptrInt = &intVar; */

/*   printf("ptrInt = %p.\n", ptrInt); */
/*   printf("sizeof(int) = %ld.\n", sizeof(int)); */
/*   printf("ptrInt + 1 = %p.\n", ptrInt + 1); */





/*   return 0; */
/* } */

/* int main() { */
/*   int intVar = 1025; */
/*   int *ptrInt = &intVar; */
/*   char *ptrChar = NULL; */

/*   printf("sizeof(int) = %ld.\n", sizeof(int)); */
/*   printf("ptrInt = %p, *ptrInt = %d.\n", ptrInt, *ptrInt); */

/*   ptrChar = (char*)(ptrInt); */

/*   printf("sizeof(char) = %ld.\n", sizeof(char)); */
/*   printf("ptrChar = %p, *ptrChar = %c.\n", ptrChar, *ptrChar); */
/*   printf("(ptrChar + 1) = %p, *(ptrChar + 1) = %d.\n", (ptrChar + 1), */
/* 	 *(ptrChar + 1)); */




/*   return 0; */
/* } */

void increment(int *p) {
  *p = (*p) + 1;
}

int main() {
  /* int intVar = 1025; */
  /* int *ptrInt = &intVar; */
  /* void *ptrVoid = ptrInt; */

  /* printf("sizeof(int) = %ld.\n", sizeof(int)); */
  /* printf("ptrInt = %p, *ptrInt = %d.\n\n", ptrInt, *ptrInt); */

  /* printf("ptrVoid = %p.\n", ptrVoid); */
  /* printf("*ptrVoid = %d.\n", *ptrVoid); */

  /* int intVar = 5; */
  /* int *ptrInt = &intVar; */
  /* int *(*ptrPtrInt) = &ptrInt; */
  /* int *(*(*ptrPtrPtrInt)) = &ptrPtrInt; */

  /* printf("intVar = %d, &intVar = %p.\n", intVar, &intVar); */
  /* printf("*ptrInt = %d, ptrInt = %p.\n", *ptrInt, ptrInt); */
  /* printf("&ptrInt = %p.\n", &ptrInt); */
  /* printf("*ptrPtrInt = %p.\n", *ptrPtrInt); */
  /* printf("ptrPtrInt = %p.\n", ptrPtrInt); */
  /* printf("ptrPtrPtrInt = %p.\n", ptrPtrPtrInt); */
  /* printf("*ptrPtrPtrInt = %p\n", *ptrPtrPtrInt); */

  /* printf("*(*ptrPtrInt) = %d.\n\n", *(*ptrPtrInt)); */

  /* *(*(*ptrPtrPtrInt)) = 17; */

  /* printf("intVar = %d.\n", intVar); */

  int intVar = 0;

  intVar = 10;

  increment(&intVar);

  printf("intVar = %d.\n", intVar);





  return 0;
}
