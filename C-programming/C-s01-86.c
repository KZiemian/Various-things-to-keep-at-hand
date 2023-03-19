#include <stdio.h>

int main() {
  int intVar = 12345;
  int printfRetVal = 0;

  printfRetVal = printf("Hello, World!\n");
  printf("printfRetVal: %i\n\n", printfRetVal);
  
  printfRetVal = printf("Tyle \n");
  printf("printfRetVal: %i\n\n", printfRetVal);

  printfRetVal = printf("Tyle %i\n", intVar);
  printf("printfRetVal: %i\n", printfRetVal);



  return 0;
}
