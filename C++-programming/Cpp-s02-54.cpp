T *p;
T const *pc;
T volatile *pv;

pc = p;
pv = p;

int i = 37;
int const ci = 2;

void wi(int& r);
void wci(int const& rc);

wi(i);
wci(ci);

R foo(T const *p) {
  // ...

  for (; *p != v; ++p) {
    if (*p == one_thing) {
      // *p = another_thing;
    }
  }

  // ...
}

class widget {
public:
  widget(widget const& w);
  widget &operator=(widget const& w);
};

class string {
public:
  // ...
  size_t size();
  void clear();
  // ...


private:
  char* text;
  size_t stored_size;
};

R foo(string const& s) {
  // ...
  for (size_t i = 0; i < s.size(); ++i) {
    // ...
  }
  // ...
}

class string {
public:
  // ...
  size_t size() const;
  void clear();
  // ...



private:
  char* text;
  size_t stored_size;
};

R foo(string const& s) {
  // ...
  for (size_t i = 0; i < s.size(); ++i) {
    if (s[i] == something) {
      // ...
    }
  }
  // ...
}

class string {
public:
  // ...
  char &operator[](size_t i) const {
    return text[i];
  }
  // ...



private:
  char *text;
  size_t stored_size;
};

class string {
public:
  char &operator[](size_t i) {
    return text[i];
  }

  char const &operator[](size_t i) const {
    return text[i];
  }
  // ...
};
