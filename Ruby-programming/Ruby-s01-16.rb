#!/home/kamil/.rbenv/shims/ruby
# coding: utf-8

# number = 127
# number_str = number.to_s

# puts "number_str.length -> #{number_str.length}"

# def kaprekar?(k)
#   number_str = k.to_s
#   n = number_str.length

#   square_k_str = (k**2).to_s
#   length_of_square_k = square_k_str.length

#   number_01 = square_k_str[0..(length_of_square_k - n - 1)].to_i
#   number_02 = square_k_str[(length_of_square_k - n)..(-1)].to_i

#   puts number_01
#   puts number_02

#   if (number_01 + number_02) == k
#     return true
#   else
#     return false
#   end
# end

# puts kaprekar?(9)

# puts (-1).abs

# for i in 0...5
#   puts "i -> #{i}"
# end

# str_test = "abcdefg"

# puts str_test
# puts str_test[0, str_test.length - 1]


# puts "Enter first number: "

# num1 = gets.chomp().to_f

# puts "Enter operator: "

# operatorVar = gets.chomp()

# puts "Enter second number: "

# num2 = gets.chomp().to_f

# if operatorVar == "+"
#   puts (num1 + num2)
# elsif operatorVar == "-"
#   puts (num1 - num2)
# elsif operatorVar == "*"
#   puts (num1 * num2)
# elsif operatorVar == "/"
#   puts (num1 / num2)
# else
#   puts "Unknow operator: #{operatorVar}"
# end


# def get_day_name(day)
#   day_name = ""

#   case day
#   when "mon"
#     day_name = "Monday"

#   when "tue"
#     day_name = "Tuesday"

#   when "wed"
#     day_name = "Wednsday"

#   when "thu"
#     day_name = "Thursday"

#   when "fri"
#     day_name = "Friday"

#   when "sat"
#     day_name = "Saturday"

#   when "sun"
#     day_name = "Sunday"

#   else
#     puts "Ivalid abbreviation: #{day}."

#     day_name = "Get invalid abbreviation."
#   end

#   return day_name
# end

# puts "mon -> #{get_day_name("mon")}"
# puts "tue -> #{get_day_name("tue")}"
# puts "wed -> #{get_day_name("wed")}"
# puts "thu -> #{get_day_name("thu")}"
# puts "fri -> #{get_day_name("fri")}"
# puts "sat -> #{get_day_name("sat")}"
# puts "sun -> #{get_day_name("sun")}"

# puts "I.R.S. -> #{get_day_name("I.R.S.")}"
