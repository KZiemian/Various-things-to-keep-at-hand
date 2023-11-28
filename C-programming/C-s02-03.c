#include <stdio.h>
/* #include <time.h> */

/* int main() { */
/*   int i = time(NULL); */
/*   const time_t *timeVar = (const time_t*)(&i); */

/*   printf("i = %d\n", i); */
/*   printf("ctime: %s\n", ctime(timeVar)); */





/*   return 0; */
/* } */

static char *strings[] = {
  "this is string one",
  "this is string two",
  "this is string three",
};

const int string_no = (sizeof strings) / (sizeof strings[0]);

int main() {
  printf("string_no = %d.\n", string_no);





  return 0;
}
