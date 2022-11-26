#include <iostream>

int main() {
  int const x = 3;
  int const y = 4;

  int const *p = &x;


  std::cout << "*p: " << *p << '\n';
  std::cout << " x: " << x << '\n';
  std::cout << " y: " << y << '\n';

  p = &y;

  std::cout << "*p: " << *p << '\n';

  // *p = 7;



  return 0;
}
