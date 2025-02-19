#!/home/kamil/.rbenv/shims/ruby
# coding: utf-8

# File.open("tornado-of-souls.txt", "r") do |fileVar|
#   puts fileVar.readline()
#   puts fileVar.readline()
# end

# File.open("tornado-of-souls.txt", "r") do |fileVar|
#   puts fileVar.readchar()
#   puts fileVar.readchar()
# end

# File.open("tornado-of-souls.txt", "r") do |fileVar|
#   for line in fileVar.readlines()
#     puts line
#   end
# end

# fileVar = File.open("tornado-of-souls.txt", "r")

# puts fileVar.read

# fileVar.close()


# File.open("tornado-of-souls.txt", "a") do |fileVar|
#   fileVar.write("Band: Megadeth")
# end

# File.open("tornado-of-souls.txt", "w") do |fileVar|
#   fileVar.write("Band: Megadeth")
# end

# File.open("tornado-of-souls-New.txt", "w") do |fileVar|
#   fileVar.write("Band: Megadeth\n")
#   fileVar.write("Title: Tornado of souls")
# end

# File.open("tornado-of-souls-03.txt", "r+") do |fileVar|
#   fileVar.readline
#   fileVar.write("Band: Megadeth")
# end

# num = 10 / 0

# some_nums = [1, 3, 9, 12, 24]

# puts "some_nums[0] = #{some_nums[0]}"
# puts "some_nums[10] = #{some_nums[dog]}"
# puts some_nums[10]

# begin
#   num = 10 / 0
# rescue
#   puts "Code rescued"
# end

# some_nums = [1, 3, 6, 9, 12]

# begin
#   some_nums[dog]
# rescue
#   puts "We rescue code."
# end

# some_nums = [1, 3, 6, 9, 12]

# begin
# # num = 10 / 0
#   some_nums[dog]
# rescue ZeroDivisionError
#   puts "Division by zero error"
# rescue NameError => exceptionVar
#   puts "Wrong name"
#   puts exceptionVar
# end

class Book
  attr_accessor :title, :author, :pages
end

book1 = Book.new()

book1.title = "Harry Potter"
book1.author = "J.K. Rowling"
book1.pages = 400

puts book1.title
puts book1.author
puts book1.pages

book2 = Book.new()
book2.title = "Lord of the Rings"
book2.author = "J.R.R. Tolkien"
book2.pages = 500

puts book2.title
puts book2.author
puts book2.pages
