#include <stdio.h>
/* #include <stdlib.h> */

#include <math.h>

/* #include <unistd.h> */

/* int main() { */
/*   int i = 0; */
/*   char bufferChar[10] = "Karakorum"; */
/*   FILE *ptrFile = fopen("data_1.dat", "w+"); */

/*   if (ptrFile == NULL) { */
/*     printf("Blad otwarcia pliku.\n"); */
/*   } */

/*   fprintf(ptrFile, "Eureka...\n"); */
/*   fprintf(ptrFile, "Eureka, drugi raz...\n"); */

/*   fflush(ptrFile); */

/*   printf("fileno(ptrFile) = %d\n", fileno(ptrFile)); */

/*   write(fileno(ptrFile), bufferChar, 7); */
/*   fprintf(ptrFile, "\n"); */





/*   return 0; */
/* } */

/* const double PI = 3.1415; */

/* #define PI 3.1415 */

int main() {
  /* double radius = 2; */

  /* double areaOfCircle = PI * radius * radius; */

  /* int sizeOfType = sizeof(short int); */

  /* printf("sizeOfType = %d.\n", sizeOfType); */

  /* sizeOfType = sizeof(int); */

  /* printf("sizeOfType = %d.\n", sizeOfType); */

  /* int intVar = 5000000000; */

  /* printf("intVar = %d.\n", intVar); */

  /* printf("PI = %.4f.\n", PI); */

  /* PI = 3.0; */

  /* printf("areaOfCircle = %.4f.\n", areaOfCircle); */

  /* double doubleVar = 0.25; */

  /* printf("doubleVar = %.2f.\n", doubleVar); */

  /* doubleVar++; */

  /* printf("doubleVar = %.2f.\n", doubleVar); */

  /* int intVar = 0; */
  /* int intVarTest = 0; */

  /* printf("intVar = %d.\n", intVar); */

  /* intVar++; */

  /* printf("intVar = %d.\n", intVar); */

  /* intVar++; */
  /* intVar++; */

  /* printf("intVar = %d.\n", intVar); */

  /* intVar--; */

  /* printf("intVar = %d.\n", intVar); */

  /* intVarTest = intVar++; */

  /* printf("intVar = %d.\n", intVar); */
  /* printf("intVarTest = %d.\n", intVarTest); */

  /* int i = 0; */

  /* do { */
  /*   printf("i = %d.\n", i); */

  /*   i++; */
  /* } while (i < 0); */



  /* double x_0 = 2.0; */
  /* double x_n = x_0; */

  /* double x_0 = 2.0; */
  /* double x_n = x_0; */

  /* x_n = 0.5 * (x_n + x_0 / x_n); */

  /* printf("x_n = %.4f.\n", x_n); */

  /* x_n = 0.5 * (x_n + x_0 / x_n); */

  /* printf("x_n = %.5f.\n", x_n); */

  /* x_n = 0.5 * (1.4167 + x_0 / 1.4167); */

  /* printf("x_n = %.5f.\n", x_n); */

  /* x_n = 0.5 * (1.4142 + x_0 / 1.4142); */

  /* printf("x_n = %.5f.\n", x_n); */

  /* x_n = 0.5 * (1.41421 + x_0 / 1.41421); */

  /* printf("x_n = %.5f.\n", x_n); */

  double x_0 = 1000.0;
  double x_n = x_0;
  double distance = fabs(x_0 - x_n * x_n);

  int numberOfSteps = 0;

  while (distance > 0.001) {
    x_n = 0.5 * (x_n + x_0 / x_n);

    distance = fabs(x_0 - x_n * x_n);
    /* printf("x_n = %.5f.\n", x_n); */
    /* printf("distance = %.5f.\n", distance); */

    numberOfSteps++;
  }

  printf("x_n = %.4f.\n", x_n);
  printf("x_n * x_n = %.4f.\n", x_n * x_n);
  printf("numberOfSteps = %d.\n", numberOfSteps);

  /* distance = abs(10.0 - 3.196 * 3.196); */

  /* printf("distance = %.5f.\n", distance); */

  /* printf("3.196 * 3.196 = %.4f.\n", 3.196 * 3.196); */





  return 0;
}
