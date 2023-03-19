#include <stdio.h>

int main() {
  char bufferChar[] = "lasso";
  char *ptrStrVal = NULL;
  char *ptrString = "Rak";

  printf("ptrString: %s\n", ptrString);

  ptrStrVal = (char*)(ptrString);
  ptrString = (char*)(bufferChar);

  *ptrString = 'S';

  printf("ptrString: %s\n", ptrString);
  printf("ptrStrVal: %s\n", ptrStrVal);



  return 0;
}
