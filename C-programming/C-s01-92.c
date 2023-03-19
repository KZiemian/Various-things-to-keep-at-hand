#include <stdio.h>

int main() {
  const char arrayConst[] = "Rak";

  printf("arrayConst: %s\n", arrayConst);

  arrayConst[0] = 'S';

  printf("arrayConst: %s\n", arrayConst);



  return 0;
}


