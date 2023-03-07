#include <stdio.h>

struct przyklad {
  int i;
  float f;
};

int main() {
  struct przyklad structVar1;
  struct przyklad structVar2 = { 1, 1.5 };

  structVar1.i = 7;
  structVar1.f = 7.25;


  printf("structVar1.i: %d, structVar1.f: %f\n\n", structVar1.i, structVar1.f);
  printf("structVar2.i: %d, structVar2.f: %f\n", structVar2.i, structVar2.f);


  return 0;
}
