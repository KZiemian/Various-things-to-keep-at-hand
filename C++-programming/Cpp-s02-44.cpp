for (auto& e : c) {
  use(e);
 }

for (e : c) {
  use(e);
 }

unique_ptr<widget> factory();

void caller() {
  auto w = factory();
  auto g = make_unique<gadget>();

  use(*w, *g);
}

void f(widget& w) {
  use(w);
}

void g(widget* w) {
  if (w) {
    use(*w);
  }
}

auto upw = make_unique<widget>();

f(*upw);

auto spw = make_shared<widget>();

g(spw.get());

unique_ptr<widget> factory();

void sink(unique_ptr<widget> upw);

void reseat(unique_ptr<widget>& upw);

shared_ptr<widget> factory();

void share(shared_ptr<widget> spw);

void reseat(shared_ptr<widget>& spw);

void may_share(const shared_ptr<widget>& cspw);

shared_ptr<widget> g_p;

void f(widget& w) {
  g();
  use(w);
}

void g() {
  g_p = ...;
}

void my_code() {
  auto pin = g_p;
  f(*pin);
}

void my_code() {
  auto pin = g_p;
  f(*pin);
  pin->foo();
}

auto var = init;

auto var = type{ init };

type var{ init };

template<class Container, class Value>
void someFunction(Container& c, Value& v) {
  if (find(begin(c), end(c), v) == end(c)) {
    c.push_back(move(v));
  }

  assert(c.empty() == false );
}

void f(const vector<int>& v) {
  vector<int>::const_iterator i = v.begin();
  // ...
}

void f(contst vector<int>& v) {
  auto i = v.begin();
  // ...
}

int i = f(1, 2, 3) * 42;
int i = f(1, 2, 3) * 42.0; // ?!?!

widget w = factory();
