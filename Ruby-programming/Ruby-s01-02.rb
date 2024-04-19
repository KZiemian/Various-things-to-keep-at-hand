#!/usr/bin/ruby

# print "\n\nDrugi each\n"

# my_array.each { |item| puts "#{item}" }

# print "\n\nTimes\n"

# n = 3

# n.times { puts "Ruby" }

# print "\n\nPetla while do\n"

# i = 1

# while i <= 10 do
#   puts i
#   i += 1
# end

# print "\n\nHashes, czesc I\n"

# tablica_hashowa_1 = { 1 => 'jeden', 'dwa' => 2 }

# puts "tablica_hashowa_1[1] = #{tablica_hashowa_1[1]}"
# puts "tablica_hashowa_1[\"dwa\"] = #{tablica_hashowa_1["dwa"]}"

# print "\n\nHashes, czesc II\n"

# tablica_hashowa_2 = Hash.new

# tablica_hashowa_2["kot"] = "rudy"
# tablica_hashowa_2["pies"] = "bialy"

# puts "tablica_hashowa_2[\"kot\"] = #{tablica_hashowa_2["kot"]}"
# puts "tablica_hashowa_2[\"pies\"] = #{tablica_hashowa_2["pies"]}"

# print "\n\nIteracja po tablicy haszowej. Porzadek czy nie?"

# tablica_hashowa_3 = { 4 => 'cztery', 1 => 'jeden', 2 => 'dwa', 3 => 'trzy'}

# tablica_hashowa_3.each do |haslo, wartosc|
#   puts "#{haslo} => #{wartosc}"
# end

# print "\n\nSortowanie tablic hashowych"

# tablica_hashowa_3 = { 4 => 'cztery', 1 => 'jeden', 2 => 'dwa', 3 => 'trzy'}

# tablica_hashowa_3 = tablica_hashowa_3.sort_by do |liczba, slowo|
#   liczba
# end

# tablica_hashowa_3.each do |haslo, wartosc|
#   puts "#{haslo} => #{wartosc}"
# end
# print "\n"

# colors = { "blue" => 3, "green" => 2, "red" => 1 }

# colors.each do |haslo, wartosc|
#   puts "#{haslo} => #{wartosc}"
# end
# puts ""

# colors = colors.sort_by do |color, numer|
#   numer
# end

# colors.each do |haslo, wartosc|
#   puts "#{haslo} => #{wartosc}"
# end

# puts "Metody\n"

# def hello_world_1
#   puts "Hello World 1!"
# end

# hello_world_1()
# hello_world_1

# def hello_world_2() puts "Hello World 2!"
# end

# hello_world_2()
# hello_world_2

# print "\n\nBloki\n"

# 1.times do
#   puts "I'm a code block!"
# end

# 1.times { puts "As am I!" }

# [1, 2, 3, 4, 5].each { |i| puts 2 * i }

# [3, 2, 1, 5, 9, 7, 8, 4, 10, 6].sort!
