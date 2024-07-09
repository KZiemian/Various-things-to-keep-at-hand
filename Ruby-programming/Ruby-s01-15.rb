#!/home/kamil/.rbenv/shims/ruby
# coding: utf-8

# what_am_i = File.open("clean-slate.txt", "w") do |file|
#   file.puts "Call me Ishmael."
# end

# puts "p what_am_i"
# p what_am_i

# File.open("clean-slate.txt", "r") { |file| puts file.read }

# file = File.open("master", "r+")

# p file.read
# file.rewind

# buffer = ""
# p file.read(23, buffer)
# p buffer

# file.close

# p File.read("tornado-of-souls.txt")

# puts

# File.open("tornado-of-souls.txt") do |fileLoc|
#   fileLoc.seek(20, IO::SEEK_SET)
#   p fileLoc.read(10)
# end

# lines = File.readlines("tornado-of-souls.txt")

# p lines

# puts

# p lines[0]

# puts Array.methods.sort()

# def palindrome?(sentence)
#   sentence_iner = sentence.downcase
#   sentence_iner = sentence_iner.split(" ").join("")

#   lengthLoc = sentence_iner.length

#   n = lengthLoc / 2
#   index = 0

#   while index < n
#     if sentence_iner[index] != sentence_iner[lengthLoc - 1 - index]
#       return false
#     end

#     index += 1
#   end

#   return true
# end

# puts "palindrome?(\"Race fast safe car\") = #{palindrome?("Race fast safe car")}"

# array_01 = [1, 2, 3]
# array_02 = array_01

# puts "array_01 -> #{array_01}"
# puts "array_02 -> #{array_02}"

# array_01.delete(1)

# puts "array_01 -> #{array_01}"
# puts "array_02 -> #{array_02}"

# array_01 = Array.new
# puts "array_01.include?(0) -> #{array_01.include?(0)}"

def non_duplicated_values(values)
  values_iner = values.clone
  result_array = Array.new
  double_elem = Array.new

  n = values_iner.length
  index = 0

  while index < n
    elem = values_iner[index]

    if !(result_array.include?(elem)) and !(double_elem.include?(elem))
      index_comp = index + 1

      while index_comp < n
        if elem == values_iner[index_comp]
          double_elem.push(elem)
          break
        else
          index_comp += 1
        end
      end

      if index_comp == n
        result_array.push(elem)
      end
    end

    index += 1
  end

  return result_array
end

puts non_duplicated_values([1, 2, 2, 3, 3, 4, 5])

puts

puts non_duplicated_values([1, 2, 2, 3, 4, 4])
