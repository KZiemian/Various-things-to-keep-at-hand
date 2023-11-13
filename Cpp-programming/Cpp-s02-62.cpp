#include <iostream>

// #include <string>

// const double RATE1 = 0.17;
// const double RATE2 = 0.32;
// const double RATE1_SINGLE_LIMIT = 85000;
// const double RATE1_MARRIED_LIMIT = 170000;

int main() {
  // std::cout << "0.17 * 85000 = " << 0.17 * 85000 << ".\n";
  // std::cout << "0.17 * 170000 = " << 0.17 * 170000 << ".\n";


  // double tax = 0.0;
  // double income = 0.0;

  // std::cout << "Proszę podać swój dochód: ";
  // std::cin >> income;

  // std::cout << "Rozliczanie somadzielne - s, z małżonkiem - m: ";

  // std::string martialStatus = "";

  // std::cin >> martialStatus;


  // if (martialStatus == "s") {
  //   if (income <= RATE1_SINGLE_LIMIT) {
  //     tax = RATE1 * income;
  //   } else {
  //     tax = 14450 + RATE2 * (income - RATE1_SINGLE_LIMIT);
  //   }
  // } else {
  //   if (income <= RATE1_MARRIED_LIMIT) {
  //     tax = RATE1 * income;
  //   } else {
  //     tax = 28900 + RATE2 * (income - RATE1_MARRIED_LIMIT);
  //   }
  // }

  // std::cout << "Podatek wynosi " << tax << ".\n";

  // double temp = -10.0;

  // if ((temp > 0.0) && (temp < 100.0)) {
  //   std::cout << "Stan cieły.\n";
  // }

  // if ((temp <= 0) || (temp >= 100)) {
  //   std::cout << "Stan inny niż ciekły.\n";
  // }

  // bool boolVar = 0 < 200 < 100;

  // std::cout << "0 < 200 < 100 = " << boolVar << ".\n";

  // boolVar = -10;

  // std::cout << "-10 = " << boolVar << ".\n";

  // int x = -2;

  // boolVar = 0 < x && x < 100 || x == -1;

  // std::cout << "boolVar = " << boolVar << ".\n";

  int floor = 0;

  std::cout << "Podaj wartość zmiennej floor: ";
  std::cin >> floor;

  std::cout << "floor = " << floor << ".\n";

  std::cout << "std::cin.fail() = " << std::cin.fail() << ".\n";

  if(std::cin.fail()) {
    std::cout << "Błąd: nie podano liczby całkowitej.\n";

    return 1;
  }






  return 0;
}
