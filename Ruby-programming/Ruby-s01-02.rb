#!/home/kamil/.rbenv/shims/ruby

# puts self

# puts 2.even?

# puts 1.next

# puts 1.methods

# puts 1.methods.sort

# puts ['rock', 'paper', 'scissors'].index('paper')

# puts 2.between?(2, 3)

# puts 4.+(3)

# puts 1+2

# puts 1 + 2

# words = ["foo", "bar", "baz"]
# puts words[1]
# puts words.[](1)

# puts "Ruby Monk"

# puts "Ruby Monk".length

# a = 1
# b = 4
# puts "The number #{a} is less than #{b}."

# def string_length_interpolater(incoming_string)
#   "The string you just gave me has a length of #{incoming_string.length}."
# end

# puts string_length_interpolater("Hi")
# puts string_length_interpolater("Good")

# puts "[Luke:] I can't believe it. [Yoda:] That is why you fail.".include?("Yoda")

# puts "Ruby is a beautiful language".start_with?("Ruby")

# puts "I can't work with any other language but Ruby".end_with?("Ruby")

# puts "I am a Rubyist".index("R")

# puts "i am in lowercase".upcase

# puts "This is Mixed CASE".downcase

# puts "ThiS iS A vErY ComPLEx SenTeNcE".swapcase

# "Fear it the path to the dark side".split(" ")

# puts "Fear is the path to the dark side".split(" ")

# puts "Ruby" + "Monk"

# puts "Ruby".concat("Monk")

# puts "Ruby" << "Monk"

# puts "I should look into your problem when I get time".sub("I", "We")

# puts "I should look into your problem when I get time".gsub("I", "We")

# puts "RubyMonk".gsub(/[aeiou]/, "1")

# puts "RubyMonk Is Pretty Brilliant".gsub(/[RMIPB]/, "0")

# puts "RubyMonk Is Pretty Brilliant".match(/ ./)

# puts "RubyMonk Is Pretty Brilliant".match(/ ./, 10)

# name = "Bob"

# puts name == "Bob"

# age = 37

# puts age <= 35

# puts age >= 23 && (name == "Bob" || name == "Jill")

# puts !(name == "Bob")

# def check_sign(number)
#   if number > 0
#     "#{number} is positive"
#   else
#     "#{number} is negative"
#   end
# end

# puts check_sign(1)
# puts check_sign(-1)

def check_sign(number)
  if number > 0
    "#{number} is positive"
  elsif number < 0
    "#{number} is negative"
  else
    "#{number}"
  end
end

puts check_sign(1)
puts check_sign(-1)
puts check_sign(0)
