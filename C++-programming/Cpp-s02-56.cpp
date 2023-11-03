// R f(T *p);
// R f(T const *p);

// R g(T& r);
// R g(T const& r);

// constexpr static auto string = PETSCII("Hello World");

// constexpr static auto string_table = make_string_table("side title",
// 						       "item1", "item2");

// constexpr static auto slides =
//   make_slide_set(
// 		 slide1_func,
// 		 slide2_func,
// 		 slide3_func);

// constexpr static auto slide_numbers =
//   make_slide_number<slides.size()>();

// #include <iostream>

// #include <cmath>

// #include <iomanip>



// int main() {
  // std::cout << std::fixed << std::setprecision(2);

  // std::cout << "pow(2, 0.5) = " << pow(2, 0.5) << ".\n";
  // std::cout << "pow(3, 0.5) = " << pow(3, 0.5) << ".\n";
  // std::cout << "pow(5, 0.5) = " << pow(5, 0.5) << ".\n";

  // std::cout << std::fixed << std::setprecision(3);

  // std::cout << "pow(2, 0.5) = " << pow(2, 0.5) << ".\n";
  // std::cout << "pow(3, 0.5) = " << pow(3, 0.5) << ".\n";
  // std::cout << "pow(5, 0.5) = " << pow(5, 0.5) << ".\n";
  // const double sqrt2 = 1.4142136523;

  // std::cout << "sqrt(2) = " << sqrt2 << ".\n";
  // std::cout << "sqrt(2) = " << std::setw(8) << sqrt2 << ".\n";
  // std::cout << std::setw(12) << "sqrt(2) = " << sqrt2 << ".\n";
  // std::cout << std::setw(12) << sqrt2 << ", " << std::setw(12)
  // 	    << sqrt2 << ".\n";




//   return 0;
// }

#include <iostream>

#include <iomanip>



int main() {
  // const double CANS_PER_PACK = 6.0;

  // double packPrice = 0.0;
  // double canVolume = 0.0;
  // double packVolume = 0.0;
  // double pricePerOunce = 0.0;

  // std::cout << "Proszę wprowadzić cenę sześciopaku: ";
  // std::cin >> packPrice;

  // std::cout << "Proszę wprowadzić pojemność puszki (w uncjach): ";
  // std::cin >> canVolume;

  // packVolume = canVolume * CANS_PER_PACK;

  // pricePerOunce = packPrice / packVolume;

  // std::cout << std::fixed << std::setprecision(2);
  // std::cout << "Cena za uncję: " << pricePerOunce << ".\n";

  // double varDouble = 12.345678;

  // std::cout << "varDouble = " << varDouble << ".\n";
  // std::cout << std::fixed << std::setprecision(2)
  // 	    << "varDouble = " << varDouble << ".\n";
  // std::cout << "varDouble = " << varDouble << ".\n";
  // std::cout << ":" << std::setw(6) << varDouble << ".\n";
  // std::cout << ":" << std::setw(6) << 12 << ".\n";
  // std::cout << "|" << std::setw(6) << ":" << 12 << ".\n";

  const int PENNIES_PER_DOLLAR = 100;
  const int PENNIES_PER_QUARTER = 25;

  int billValue = 0;
  int itemPrice = 0;
  int changeDue = 0;
  int dollarCoins = 0;
  int quarters = 0;



  std::cout << "Wprowadź nominał banknotu (1 = banknot 1 dol., 5 = 5 dol. etc.): ";
  std::cin >> billValue;

  std::cout << "Wprowadź cenę artykułu w centach: ";
  std::cin >> itemPrice;

  changeDue = PENNIES_PER_DOLLAR * billValue - itemPrice;
  dollarCoins = changeDue / PENNIES_PER_DOLLAR;
  changeDue = changeDue % PENNIES_PER_DOLLAR;
  quarters = changeDue / PENNIES_PER_QUARTER;

  std::cout << "Jednodolarówki:  " << std::setw(6) << dollarCoins << ".\n";
  std::cout << "Ćwierćdolarówki: " << std::setw(6) << quarters << ".\n";





  return 0;
}
