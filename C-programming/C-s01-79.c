#include <stdio.h>

int main() {
  int n1 = 0;
  int n2 = 0;
  
  for (n1 = 1, n2 = 1; n1 < 5 && n2 < 6; ++n1, ++n2) {
    printf("n1 = %d, n2 = %d\n", n1, n2);
  }



  return 0;
}


