#!/home/kamil/.rbenv/shims/ruby
# coding: utf-8

# def calculate(*numbers)
#   if numbers[-1].is_a?(Hash)
#     numbers_arg = numbers
#     options = numbers_arg[-1]
#     numbers_arg.delete(options)

#     if options[:subtract]
#       subtract(numbers_arg)
#     elsif options[:add]
#       add(numbers_arg)
#     else
#       add(numbers_arg)
#     end
#   end

#   add(numbers)
# end

# def add(*numbers)
#   numbers.inject(0) { |sum, number| sum + number }
# end

# def subtract(number, *numbers)
#   sum = numbers.inject(0) { |sum, number| sum + number }
#   number - sum
# end

# # puts "add(4, 5)"
# # puts add(4, 5)

# # puts "add(-10, 2, 3)"
# # puts add(-10, 2, 3)

# puts "add(0, 0, 0, 0)"
# puts add(0, 0, 0, 0)

# l = lambda { "Do or do not" }

# puts "l.call = #{l.call}"

# l = lambda do |string|
#   if string == "try"
#     return "There's no such thing"
#   else
#     return "Do or do not."
#   end
# end

# puts "l.call(\"try\") = #{l.call("try")}"

# def demonstrate_block(number)
#   yield(number)
# end

# puts demonstrate_block(1) { |number| number + 1 }

# module WarmUp
#   def push_ups
#     "Phew, I need a break!"
#   end
# end

# class Gym
#   include WarmUp

#   def preacher_curls
#     "I'm building my biceps."
#   end
# end

# class Dojo
#   include WarmUp

#   def tai_kyo_kyu
#     "Look at my stance!"
#   end
# end

# puts Gym.new.push_ups
# puts Dojo.new.push_ups

# module WarmUp
# end

# puts "WarmUp.class -> #{WarmUp.class}"
# puts "Class.superclass -> #{Class.superclass}"
# puts "Module.superclass -> #{Module.superclass}"

module Perimeter
  class Array
    def initialize
      @size = 400
    end
  end
end

our_array = Perimeter::Array.new
ruby_array = Array.new

p our_array.class
p ruby_array.class
