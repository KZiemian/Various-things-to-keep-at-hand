#!/home/kamil/.rbenv/shims/ruby
# coding: utf-8

# puts "Ruby Monk Is Pretty Brilliant".gsub(/[RMIPB]/, "0")

# puts "Ruby Monk Is Pretty Brilliant".match(/ ./)

# puts "Ruby Monk Is Pretty Brilliant".match(/ ./, 10)

# puts "true and true -> #{true and true}"
# puts "false and true -> #{false and true}"
# puts "true and false -> #{true and false}"
# puts "false and false -> #{false and false}"

# puts "true && true -> #{true && true}"
# puts "false && true -> #{false && true}"
# puts "true && false -> #{true && false}"
# puts "false && false -> #{false && false}"

# puts "true or true -> #{true or true}"
# puts "false or true -> #{false or true}"
# puts "true or false -> #{true or false}"
# puts "false or false -> #{false or false}"

# puts "true || true -> #{true || true}"
# puts "false || true -> #{false || true}"
# puts "true || false -> #{true || false}"
# puts "false || false -> #{false || false}"

# puts "true == true -> #{true == true}"
# puts "false == true -> #{false == true}"
# puts "true == false -> #{true == false}"
# puts "false == false -> #{false == false}"

# puts "1 < 1 -> #{1 < 1}"
# puts "1 < 2 -> #{1 < 2}"
# puts "1 > 1 -> #{1 > 1}"
# puts "2 > 1 -> #{2 > 1}"

# puts "1 <= 1 -> #{1 <= 1}"
# puts "1 <= 2 -> #{1 <= 2}"
# puts "2 <= 1 -> #{2 <= 1}"

# puts "1 >= 1 -> #{1 >= 1}"
# puts "1 >= 2 -> #{1 >= 2}"
# puts "2 >= 1 -> #{2 >= 1}"

# name = "Jill"

# puts "!(name == \"Jill\") = #{!(name == "Jill")}"

# def check_sign(number)
#   if number > 0
#     "#{number} is positive"
#   else
#     "#{number} is negative"
#   end
# end

# puts "check_sign(1) = #{check_sign(1)}"
# puts "check_sign(0) = #{check_sign(0)}"
# puts "check_sign(-1) = #{check_sign(-1)}"

# def check_sign(number)
#   if number > 0
#     "#{number} is positive"
#   elsif number < 0
#     "#{number} is negative"
#   else
#     "#{number}"
#   end
# end

# puts "check_sign(1) = #{check_sign(1)}"
# puts "check_sign(0) = #{check_sign(0)}"
# puts "check_sign(-1) = #{check_sign(-1)}"

# def check_sign(number)
#   number > 0 ? "#{number} is positive" : "#{number} is negative"
# end

# puts "check_sign(1) = #{check_sign(1)}"
# puts "check_sign(0) = #{check_sign(0)}"
# puts "check_sign(-1) = #{check_sign(-1)}"

# arrayVar = Array.new

# arrayVar[0] = 0
# arrayVar[1] = 1

# puts "arrayVar -> #{arrayVar}"

# puts "[1, 2, 3, 4, 5][2] = #{[1, 2, 3, 4, 5][2]}"
# puts "[1, 2, 3, 4, 5, 6, 7][4] = #{[1, 2, 3, 4, 5, 6, 7][4]}"

# puts "[1, 2, 3, 4, 5][-1] = #{[1, 2, 3, 4, 5][-1]}"

# puts "[1, 2, 3, 4, 5].map{ |i| i + 1 } = #{[1, 2, 3, 4, 5].map{ |i| i + 1 }}"

puts "[1, 2, 3, 4, 5, 6].select{ |number| number % 2 == 0} = #{[1, 2, 3, 4, 5, 6].select{ |number| number % 2 == 0 }}"
