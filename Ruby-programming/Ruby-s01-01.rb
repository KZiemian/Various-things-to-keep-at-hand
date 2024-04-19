#!/usr/bin/ruby

# print "Wypisywanie na wyjscie\n"

# # Komentarz.
# print "Hello "
# puts "World!"
# puts "Hello World!".length
# puts "Kamil".reverse

# =begin
# 1
# 2
# 3
# =end

# print "Podaj swoje imie: "
# imie = gets.chomp
# puts imie

# print "\n\nInstrukcje warunkowe\n"

# liczba = 1

# if (liczba == 0)
#   puts "Rowne 0"
# elsif (liczba == 1)
#   puts "Rowne 1"
# else
#   puts "Jakas inna liczba"
# end

# zmienna_bool = false

# unless zmienna_bool
#   puts "Rozumiem"
# else
#   puts "Nie rozumiem"
# end

# print "\n\nOperacje na stringach\n"

# print "Thring, plesthe!: "
# user_input = gets.chomp


# if user_input.include? "s"
#   user_input.gsub!(/s/, "th")
#   if user_input.include? "S"
#     user_input.gsub!(/S/, "Th")
#   end
# end

# puts "Wynik to: #{user_input}"

# print "\n\nPetla while\n"

# counter = 1

# while counter < 11
#   puts counter
#   counter += 1
# end

# print "\n\nPetla untile\n"

# counter = 1

# until counter == 8
#   puts counter
#   counter += 1
# end

# puts "\n\nPierwsza petla for\n"

# for num in 1..10
#   puts num
# end

# puts "\n\nDruga petla for\n"

# for num in 1...10
#   puts num
# end

# print "\n\nPetla loop, I przyklad\n"

# i = 20

# loop do
#   i -= 1
#   puts "i = #{i}"
#   break if i <= 0
# end

print "Tablice\n"

my_array = [1, 2, 3, 4, 5]

print "\n\nPierwszy each\n"

my_array.each do |x|
  x += 10
  print "#{x} "
end
puts ""
