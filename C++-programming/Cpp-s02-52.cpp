#include <iostream>

int main() {
  int x = 3;
  int const *const p = &x;

  std::cout << "*p: " << *p << '\n';



  return 0;
}
