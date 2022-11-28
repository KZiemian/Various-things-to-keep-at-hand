map<string, string>::iterator i = begin(dict);

auto i = f(1, 2, 3) * 42.0;

auto w = factory();

auto i = begin(dict);

auto strVar = "Hello";

auto w = get_widget();

auto e = employee{ empid };

auto w = widget{ 12, 34 };

auto w = make_unique<widget>();

auto f(double x) -> int;

auto f(double x) {
  // ...
}

auto e = employee{ empid };

auto w = get_widget();

auto x = 42;

auto x = 42.f; // float

auto x = 42; // unsigned int

auto x = "42"s; // std::string

auto x = 42ns; // chrono::nanoseconds

auto f(double x) -> int;

auto f = [=](double) {
  // ...
};

using dict = set<string>;

template<class T>
using myvec = vector<T,myalloc>;

category name = type [initializer];

auto e = employee{ empid };
auto w = get_widget();
auto x = 42;
auto x = 42.f;
auto x = 42.ul;
auto x = "42"s;
auto x = 1.2ns;
auto func(double x) -> int;
auto func = [=](double x) {
  // ...
};
using dict = set<string>;
template<class T>
using myvec = vector<T,myalloc>;
auto x = type{value};

// auto lock = lock_guard<mutex>{ m };
// auto ai = atomic<int>{};
// auto a = array<int,50>{};

auto x = {1};
auto x = 1;
auto a = matrix{...}, b = matrix{...};
auto ab = a * b;
auto c = matrix{ a * b };
auto x = {long long}{ 42 };
auto y = class X{1, 2, 3};

base* pb = new derived();

unique_ptr<base> pb = make_uqniue<derived>();

auto pb = unique_ptr<base>{ make_unique<derived>() };

int const number = 8675309;
char const msg[] = "hello";

widget *const cpw = ...; // const pointer to widget
widget *const *pcpw = ...; // pointer to const pointer to widget
widget **const cpw = ...; // const pointer to pointer to widget

uint32_t volatile *const x[N];

T *p; // pointer
T const *pc; // pointer to const

void wp(T *q); // wants pointer
void wpc(T const *qc); // wants pointer to const

wp(pc);
wpc(p);
