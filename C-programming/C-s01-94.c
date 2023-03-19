#include <stdio.h>
#include <string.h>

int main() {
  const char arrayConst[] = "Rak";
  char *ptrCharArray = NULL;

  printf("arrayConst: %s\n", arrayConst);

  ptrCharArray = (char*)(arrayConst);

  strcpy(ptrCharArray, "fa");

  printf("arrayConst: %s\n", arrayConst);
  printf("ptrCharArray: %s\n", ptrCharArray);



  return 0;
}
