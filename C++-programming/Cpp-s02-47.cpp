#include <iostream>

int const n = 3;

struct foo {
  int foo1: 2 * n;
  int foo2: 4;
};

int main() {
  struct foo fooVar;

  std::cout << "foo.foo1: " << fooVar.foo1 << '\n';



  return 0;
}
