/* int i = 0; */
/* int j = 0; */
/* int k = 0; */

/* float X = 0.0; */
/* float Y = 0.1; */

/* void func1(int a, int b) { */
/*   char c1 = 'a'; */
/*   char c2 = 'a'; */

/*   float B = 0.0; */
/* } */

/* void func2(float Z1, float Z2, char charVar) { */
/*   int A1[15]; */
/*   long int A2[15][15]; */

/*   float B1 = 0.0; */
/*   float B2 = 0.0; */
/*   float B3 = 0.0; */
/* } */

/* int main() { */
/*   int m = 0; */
/*   int n = 0; */
/*   int p = 0; */
/*   int q = 0; */

/*   float V1 = 0.0; */
/*   float V2 = 0.0; */
/*   float V3 = 0.0; */

/*   long int T1[15][15]; */
/*   long int T2[15][15]; */





/*   return 0; */
/* } */

/* int i = 5; */

/* int func1(int n) { */
/*   int i = 7; */


/*   return i + n; */
/* } */

/* int func2(int m) { */
/*   return i + m; */
/* } */

/* int main() { */
/*   int k = 0; */
/*   int z = 0; */
/*   int i = 0; */

/*   k = func1(0); */
/*   z = func2(0); */


/*   return 0; */
/* } */

/* float eps = 0.001; */
/* /\* ... *\/ */
/* /\* double eps = 0.05; *\/ */

/* /\* ... *\/ */

/* int main() { */
/*   long int kLongInt1 = 0; */

/*   int kInt1 = 0; */





/*   return 0; */
/* } */

#include <stdio.h>

int licznik() {
  static int counter = 0;

  counter++;


  return counter;
}

int main() {
  int number = 0;

  printf("Pierwsze wywołanie licznik().\n");
  printf("number = %d.\n", licznik());

  printf("Drugie wywołanie licznik().\n");
  printf("number = %d.\n", licznik());

  printf("Trzecie wywołanie licznik().\n");
  printf("number = %d.\n", licznik());





  return 0;
}
