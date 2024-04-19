#!/usr/bin/ruby

# books = ["BCD", "EFG", "MNO", "DEF", "ABC", "CDE"]

# books.sort!

# puts books
# print books, "\n"


# string_1 = "ABC"
# string_2 = "BCD"

# puts string_1

# puts "(string_1 <=> string_2) = #{string_1 <=> string_2}"

# # puts string_1 <=> string_2

# def cos()
#   puts "Dziala"
# end

# cos
# cos()

# print "\n\nTablice hashowe\n"

# lisp_hash = Hash.new("Lisp")

# puts lisp_hash["cos"]
# puts lisp_hash["cos innego"]

# print "\n\nSymbole i stringi\n"

# puts "string".object_id
# puts "string".object_id

# puts :symbol.object_id
# puts :symbol.object_id

# print "\n\nSymbole i stringi\n"

# napis = "string"
# symbol_przyk = :symbol

# puts napis
# # puts symbol_przyk

# puts "String to symbol: #{napis} -> #{napis.to_sym}"
# puts "Symbol to string: #{symbol_przyk} -> #{symbol_przyk.to_s}"

# print "\n\nTablice haszowe i symbole\n"

# tablica_hashowa_4 = {
#   one: 1,
#   two: 2,
#   three: 3
# }

# puts tablica_hashowa_4

# print "\n\nCzas dla symboli i stringow\n"

# require 'benchmark'

# string_AZ = Hash[("a".."z").to_a.zip((1..26).to_a)]
# symbol_AZ = Hash[(:a..:z).to_a.zip((1..26).to_a)]

# string_time = Benchmark.realtime do
#   100_000.times { string_AZ["r"] }
# end

# symbol_time = Benchmark.realtime do
#   100_000.times { symbol_AZ[:r] }
# end

# puts "String time: #{string_time} seconds"
# puts "Symbol time: #{symbol_time} seconds"

# print "\n\nMetoda select\n"

# movie_ratings = { memento: 3, primer: 3.5, the_matrix: 5, truman_show: 4,
#                   red_dawn: 1.5, skyfall: 4, alex_corss: 2, uhf: 1,
#                   lion_king: 3.5 }

# good_movies = movie_ratings.select { |movie, grade| grade > 3 }

# puts good_movies

# print "\n\nMetody each_key, each_value\n"

# movie_ratings.each_key { |key| puts key }
# movie_ratings.each_value { |value| puts value }
