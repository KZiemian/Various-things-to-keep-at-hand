#!/home/kamil/.rbenv/shims/ruby
# coding: utf-8

# module Dojo
#   A = 4

#   module Kata
#     B = 8

#     module Roulette
#       class ScopeIn
#         def push
#           15
#         end
#       end
#     end
#   end
# end

# A = 16
# B = 23
# C = 42

# puts "A -> #{A}"
# puts "Dojo::A -> #{Dojo::A}"

# puts

# puts "B -> #{B}"
# puts "Dojo::Kata::B -> #{Dojo::Kata::B}"

# puts

# puts "C -> #{C}"
# puts "Dojo::Kata:Roulette::ScopeIn.new.push -> #{Dojo::Kata::Roulette::ScopeIn.new.push}"

# fd = IO.sysopen("new-fd.txt", "w")

# p IO.new(fd)

# io_streams = Array.new
# ObjectSpace.each_object(IO) { |x| io_streams << x }

# p io_streams

# p STDOUT.class
# p STDOUT.fileno

# puts "STDOUT.class -> #{STDOUT.class}"
# puts "STDOUT.fileno -> #{STDOUT.fileno}"

# p STDIN.class
# p STDIN.fileno

# puts "STDIN.class -> #{STDIN.class}"
# puts "STDIN.fileno -> #{STDIN.fileno}"

# p STDERR.class
# p STDERR.fileno

# puts "STDERR.class -> #{STDERR.class}"
# puts "STDERR.fileno -> #{STDERR.fileno}"

# p $stdin.object_id
# p STDIN.object_id

# puts "$stdin.object_id -> #{$stdin.object_id}"
# puts "STDIN.object_id -> #{STDIN.object_id}"

# p $stdout.object_id
# p STDOUT.object_id

# puts "$stdout.object_id -> #{$stdout.object_id}"
# puts "STDOUT.object_id -> #{STDOUT.object_id}"

# p $stderr.object_id
# p STDERR.object_id

# puts "$stderr.object_id -> #{$stderr.object_id}"
# puts "STDERR.object_id -> #{STDERR.object_id}"

# puts "$stdin.object_id -> #{$stdin.object_id}"
# puts "STDIN.object_id -> #{STDIN.object_id}"

# puts

# puts "$stdout.object_id -> #{$stdout.object_id}"
# puts "STDOUT.object_id -> #{STDOUT.object_id}"

# puts

# puts "$stderr.object_id -> #{$stderr.object_id}"
# puts "STDERR.object_id -> #{STDERR.object_id}"

# mode = "w+"
# file = File.open("friend-list.txt", mode)

# puts file.inspect
# puts file.read

# file.close
