R f(T *p);
R f(T const *p);

R g(T& r);
R g(T const& r);

constexpr static auto string = PETSCII("Hello World");

constexpr static auto string_table = make_string_table("side title",
						       "item1", "item2");

constexpr static auto slides =
  make_slide_set(
		 slide1_func,
		 slide2_func,
		 slide3_func);

constexpr static auto slide_numbers =
  make_slide_number<slides.size()>();
