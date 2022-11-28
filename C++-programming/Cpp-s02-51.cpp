#include <iostream>

int main() {
  int x = 0;
  // int y = 1;

  int *const p = &x;


  std::cout << " x: " << x << '\n';

  std::cout << "*p: " << *p << '\n';


  *p = 5;


  std::cout << "*p: " << *p << '\n';

  std::cout << " x: " << x << '\n';


  // p = &y;


  // std::cout << "*p: " << *p << '\n';




  return 0;
}
