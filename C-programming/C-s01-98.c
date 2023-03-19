#include <stdio.h>
#include <string.h>

int main(int argc, char *argv[]) {
  if (argc >= 3) {
    if (strcmp("alfa", argv[2]) == 1) {
      printf("Trzeci argument to \"alfa\".\n");
    } else {
      printf("Trzeci argument jest rozny od \"alfa\".\n");
    }
  } else {
    printf("Argumentow jest co najwyzej 2.\n");
  }



  return 0;
}
