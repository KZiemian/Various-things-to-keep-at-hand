#include <stdio.h>
/* #include <stdlib.h> */
/* #include <string.h> */


/* int a = 1; */

/* void func1() { */
/*   int x = 2; */

/*   static int y = 0; */

/*   x--; */
/*   y++; */
/*   a *= 2; */

/*   printf("x = %d, y = %d, a = %d.\n", x, y, a); */
/* } */

/* int main() { */
/*   func1(); */
/*   func1(); */
/*   func1(); */





/*   return 0; */
/* } */

/* int a = 0; */

/* void func1() { */
/*   a++; */
/*   /\* b++; *\/ */
/* } */

/* int b = 0; */

/* void func2() { */
/*   a++; */
/*   b++; */

/*   func1(); */
/* } */

/* int main() { */
/*   printf("a = %d, b = %d.\n", a, b); */

/*   func1(); */

/*   printf("a = %d, b = %d.\n", a, b); */

/*   func2(); */

/*   printf("a = %d, b = %d.\n", a, b); */





/*   return 0; */
/* } */

/* int main() { */
/*   char arrayChar1[] = "Ala ma "; */
/*   char arrayChar2[] = "kota"; */
/*   char *text = NULL; */
/*   char *temp = NULL; */

/*   text = (char*)malloc(strlen(arrayChar1) + 1); */

/*   if (!text) { */
/*     printf("Error 1.\n"); */
/*   } */

/*   strcpy(text, arrayChar1); */

/*   printf("text: %s.\n", text); */

/*   temp = (char*)malloc(strlen(text) + strlen(arrayChar2) + 1); */
/*   strcpy(temp, text); */
/*   strcat(temp, arrayChar2); */

/*   free(text); */

/*   text = temp; */

/*   printf("temp: %s.\n", temp); */

/*   free(temp); */





/*   return 0; */
/* } */

int main() {
  int intVar = 0;
  int *ptrInt = &intVar;

  printf("ptrInt = %p.\n", ptrInt);
  printf("*ptrInt = %d.\n", *ptrInt);
  printf("&intVar = %p.\n", &intVar);





  return 0;
}
