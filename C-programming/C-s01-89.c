#include <stdio.h>

int main() {
  char bufferChar[] = "lasso";
  char *ptrStrVal = NULL;
  const char *ptrString = "Rak";


  printf("ptrString: %s\n", ptrString);

  ptrStrVal = (char*)(ptrString);
  ptrString = (const char*)(bufferChar);

  printf("ptrString: %s\n", ptrString);
  printf("ptrStrVal: %s\n", ptrStrVal);



  return 0;
}
