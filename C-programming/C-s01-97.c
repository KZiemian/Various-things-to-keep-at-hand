#include <stdio.h>

int main(int argc, char **argv) {
  if (argc >= 3) {
    printf("Liczba przeslanych argumentow jest wieksza lub rowna 3.\n");
  } else {
    printf("Liczba przeslanych argumentow wynosi co najwyzej 2.\n");
  }



  return 0;
}
