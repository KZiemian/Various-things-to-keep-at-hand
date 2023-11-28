#include <stdio.h>

/* #include <math.h> */

/* int main() { */
/*   int i = 0; */
/*   char charVar = 'a'; */
/*   FILE *ptrFile = fopen("dane.txt", "r"); */

/*   while (feof(ptrFile) == 0) { */
/*     charVar = getc(ptrFile); */

/*     printf("%c, %d, feof(ptrFile) = %d, ferror(ptrFile) = %d\n", charVar, */
/* 	   charVar, feof(ptrFile), ferror(ptrFile)); */
/*   } */





/*   return 0; */
/* } */

/* #define MATH_PI 3.1415 */

/* const double MATH_PI = 3.1415; */

int main() {
  /* char charVar = 'A'; */
  /* int intVar = 65; */

  /* printf("Podaj literę: "); */
  /* scanf("%c", &charVar); */

  /* printf("charVar: %c.\n", charVar); */
  /* printf("charVar: %i.\n\n", charVar); */

  /* printf("Podaj liczbę całkowitą: "); */
  /* scanf("%i", &intVar); */

  /* printf("intVar: %i.\n", intVar); */
  /* printf("intVAr: %c.\n", intVar); */

  /* int a = 0; */
  /* int i = 0; */

  /* printf("Podaj wartość a: "); */
  /* scanf("%d", &a); */

  /* printf("Czynniki pierwsze liczby %d to: ", a); */

  /* i = 2; */

  /* while (a >= 2) { */
  /*   while ((a % i) == 0) { */
  /*     printf("%d, ", i); */

  /*     a = a / i; */
  /*   } */
  /*   i++; */
  /* } */

  /* printf("Pi = %f.\n", MATH_PI); */

  /* double valueOfSin = sin(M_PI/2); */

  /* printf("Sinus pi/2 radianów wynosi: %f.\n", valueOfSin); */

  /* int break = 0; */

  /* printf("break = %d.\n", break); */

  /* int intVar; */

  /* printf("intVar = %d.\n", intVar); */

  /* int arrayInt[3] = {1, 2, 3}; */
  /* int arrayInt[] = {-1, -2, -3, -4}; */
  /* int i = 0; */

  /* for (i = 0; i < 4; ++i) { */
  /*   /\* printf("arrayInt[%d] = %d.\n", i, arrayInt[i]); *\/ */
  /*   printf("arrayInt[%d] = %d.\n", i, arrayInt[i]); */
  /* } */

  /* printf("sizeof(char): %ld.\n", sizeof(char)); */
  /* printf("sizeof(short int): %ld.\n", sizeof(short int)); */
  /* printf("sizeof(int): %ld.\n", sizeof(int)); */
  /* printf("sizeof(long int): %ld.\n", sizeof(long int)); */

  int sizeOfType = sizeof (int);
  unsigned int sizeOfType1 = sizeof (unsigned short int);
  size_t sizeOfType2 = sizeof(long int);

  printf("sizeOfType = %d.\n", sizeOfType);
  printf("sizeOfType1 = %d.\n", sizeOfType1);
  printf("sizeOfType2 = %ld.\n", sizeOfType2);





  return 0;
}
