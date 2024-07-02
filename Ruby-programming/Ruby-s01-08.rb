#!/home/kamil/.rbenv/shims/ruby
# coding: utf-8

# friends = Array["Kevin", "Karen", "Oscar"]

# puts friends.length()
# puts friends.include? "Kevin"
# puts friends.include?("Karen")

# friends = Array["Kevin", "Karen", "Oscar", "Andy"]

# puts friends
# puts friends.sort()

# states = {
#   "Pennsylvania" => "PA",
#   "New York" => "NY",
#   "Oregon" => "OR",
# }

# puts states

# puts states["Oregon"]

# states = {
#   :Pensylvania => "PA",
#   :NewYork => "NY",
#   :Oregon => "OR",
# }

# puts states
# puts states[:Pensylvania]

# def sayhi
#   puts "Hello User"
# end

# sayhi

# def sayhi(name)
#   puts "Hello #{name}."
# end

# sayhi("Karolina")

# def sayhi(name, age)
#   puts "Hello #{name}, I'm #{age} years old."
# end

# sayhi("Karolina", 37)
# sayhi("Karolina")

# def sayhi(name, age = 37)
#   puts "Hello #{name}, I'm #{age} years old."
# end

# sayhi("Karolina", 27)
# sayhi("Karolina")

# def cube(number)
#   number * number * number
# end

# puts cube(2)

# def strangeCube(number)
#   return number * number * number, 27
# end

# puts strangeCube(2)
# puts strangeCube(2)[0]
# puts strangeCube(2)[1]
# valueReturn = strangeCube(2)
# puts valueReturn

# ismale = true
# ismale = false
# ismale = true

# if ismale
#   puts "You are male."
# end

# if ismale
#   puts "You are male."
# else
#   puts "You are female."
# end

# is_male = true
# is_tall = false

# if is_male and is_tall
#   puts "You are male and tall."
# else
#   puts "You are either not male or not tall."
# end

is_male = false
is_tall = false

if is_male and is_tall
  puts "You are male and tall."
elsif is_male and !is_tall
  puts "You are male but you are not tall."
else
  puts "You are not male."
end
