/* #include <stdio.h> */
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

/* int main() { */
/*   int intVar = 0; */
/*   int *ptrInt = &intVar; */

/*   printf("ptrInt = %p.\n", ptrInt); */
/*   printf("*ptrInt = %d.\n", *ptrInt); */
/*   printf("&intVar = %p.\n", &intVar); */





/*   return 0; */
/* } */

#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#include <math.h>

int main() {
  double x = 10.0;
  double result = 0.0;
  int randomInt = 0;

  srand(time(NULL));

  /* while (x <= 20.0) { */
  /*   randomInt = rand() % 101; */
  /*   /\* printf("randomInt = %d.\n", randomInt); *\/ */
  /*   result = exp(-x * x) + ((double)(randomInt)) / 100.0; */

  /*   printf("%.1f    %.3f\n", x, result); */

  /*   x += 0.1; */
  /* } */

  /* while (x <= 0.0) { */
  /*   randomInt = rand() % 101; */
  /*   result = exp(x) + ((double)(randomInt)) / 100.0; */

  /*   printf("%.1f    %.3f\n", x, result); */

  /*   x += 0.1; */
  /* } */

  /* while (x <= 20.0) { */
  /*   randomInt = rand() % 101; */
  /*   result = 1.0 + 0.000001 * exp(x) + ((double)(randomInt)) / 100.0; */

  /*   printf("%.1f    %.3f\n", x, result); */

  /*   x += 0.1; */
  /* } */

  while (x <= 20.0) {
    result = 0.000001 * x * x * x * x * x * x +
      0.001 * x * x * x * x * x - 0.01 * x * x * x * x -
      0.1 * x * x * x  + x * x + x - 9.25;
    randomInt = rand() % 101;

    result += (double)(randomInt) / 100.0;

    printf("%.3f    %.4f\n", x, result);

    x += 0.01;
  }




  return 0;
}
