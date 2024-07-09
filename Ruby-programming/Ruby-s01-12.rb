#!/home/kamil/.rbenv/shims/ruby
# coding: utf-8

# puts "1.class.class -> #{1.class.class}"

# puts "Object.new -> #{Object.new}"

# class Rectangle
#   def initialize(length, breadth)
#     @length = length
#     @breadth = breadth
#   end

#   def perimeter
#     2 * (@length + @breadth)
#   end

#   def area
#     @length * @breadth
#   end
# end

# puts "Rectangle.new.perimeter -> #{Rectangle.new.perimeter}"

# puts "1.next -> #{1.next}"

# puts "1.method(\"next\") -> #{1.method("next")}"

# next_method_object = 1.method("next")
# puts "next_method_object.call -> #{next_method_object.call}"

# def reverse_sign(an_integer)
#   return 0 - an_integer
# end

# puts "reverse_sign(100) -> #{reverse_sign(100)}"
# puts "reverse_sign(-5) -> #{reverse_sign(-5)}"

# def do_nothing
# end

# puts "do_nothing.class -> #{do_nothing.class}"

# def add(a_number, another_number)
#   a_number + another_number
# end

# puts "add(1, 2) -> #{add(1, 2)}"

# def add(a_number, another_number, yet_another_number)
#   a_number + another_number + yet_another_number
# end

# puts "add(1, 2, 3) -> #{add(1, 2, 3)}"

# def show_numbers(*numbers)
#   return "numbers -> #{numbers}"
# end

# puts "show_numbers(1) -> #{show_numbers(1)}"
# puts "show_numbers(1, 2) -> #{show_numbers(1, 2)}"
# puts "show_numbers(1, 2, 3) -> #{show_numbers(1, 2, 3)}"

# def add(*numbers)
#   numbers.inject(0) { |sum, number| sum + number }
# end

# puts "add(1) -> #{add(1)}"
# puts "add(1, 2) -> #{add(1, 2)}"
# puts "add(1, 2, 3) -> #{add(1, 2, 3)}"
# puts "add(1, 2, 3, 4) -> #{add(1, 2, 3, 4)}"

# def add(a_number, another_number, yet_another_number)
#   a_number + another_number + yet_another_number
# end

# numbers_to_add = [1, 2, 3]
# puts "add(*numbers_to_add) -> #{add(*numbers_to_add)}"

# def add(*numbers)
#   numbers.inject(0) { |sum, number| sum + number }
# end

# def add_with_message(message, *numbers)
#   "#{message}: #{add(*numbers)}"
# end

# puts "add_with_message(\"The sum is\", 1, 2, 3) -> #{add_with_message("The sum is", 1, 2, 3)}"

def add(a_number, another_number, options = {})
  sum = a_number + another_number
  sum = sum.abs if options[:absolute]
  sum = sum.round(options[:precision]) if options[:round]

  return sum
end

puts "add(1.0134, -5.568) -> #{add(1.0134, -5.568)}"
puts "add(1.0134, -5.568, absolute: true) -> #{add(1.0134, -5.568, absolute: true)}"
puts "add(1.0134, -5.568, absolute: true, round: true, precision: 2) -> #{add(1.0134, -5.568, absolute: true, round: true, precision: 2)}"
