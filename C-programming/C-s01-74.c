#include <stdio.h>
#include <stdlib.h>

void errorExit(void) {
  fprintf(stderr, "Blad krytyczny: Zatrzymanie pracy programu!\n");

  exit(8);
}

int main() {
  int value = 1;
  /* value = -1; */

  if (value < 0) {
    errorExit();
  }

  printf("To jeszcze nie koniec.\n");



  return 0;
}
