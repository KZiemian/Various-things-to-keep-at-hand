#include <iostream>

int main() {
  int x = 3;
  // int y = 5;

  int const *const p = &x;


  std::cout << "*p: " << *p << '\n';

  // *p = &y;
  // *p = 8;
  x = 4;

  std::cout << "*p: " << *p << '\n';




  return 0;
}
