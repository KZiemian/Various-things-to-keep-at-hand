#include <stdio.h>
#include <string.h>

int main() {
  /* char arrayChar[] = "C string"; */

  /* printf("arrayChar: %s.\n", arrayChar); */
  /* printf("sizeof(arrayChar) = %ld.\n", sizeof(arrayChar)); */
  /* printf("strlen(arrayChar) = %ld.\n", strlen(arrayChar)); */

  /* char charVar = 'a'; */
  /* char charVar1 = "d"; */

  /* printf("charVar = %c.\n", charVar); */
  /* printf("charVar = %d.\n", charVar); */
  /* printf("charVar1 = %c.\n", charVar1); */

  /* charVar = 100; */

  /* printf("charVar = %c.\n", charVar); */
  /* printf("charVar = %d.\n", charVar); */

  char charVar = '\0';

  printf("charVar = %c.\n", charVar);
  printf("charVar = %d.\n", charVar);

  char stringVar[14] = {'H', 'e', 'l', 'l', 'o', ',', ' ', 'W', 'o', 'r',
    'l', 'd', '!', '\0'};

  printf("stringVar = %s\n", stringVar);
  printf("strlen(stringVar) = %ld.\n", strlen(stringVar));
  printf("sizeof(stringVar) = %ld.\n", sizeof(stringVar));



  return 0;
}
