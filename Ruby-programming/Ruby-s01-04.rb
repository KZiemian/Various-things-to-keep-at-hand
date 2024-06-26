#!/home/kamil/.rbenv/shims/ruby

# array = [1, 2, 3, 4, 5]

# array.each do |i|
#   puts i
# end

# def array_copy(source)
#   destination = []

#   source.each do |elem|
#     if elem < 4
#       destination.push(elem)
#     end
#   end

#   return destination
# end

# puts array_copy([1, 2, 3, 4, 5])
# puts array_copy([-23, 12, 0, 19])

# student_ages = {
#   "Jack" => 10,
#   "Jill" => 12,
#   "Bob" => 14
# }

# puts student_ages

# restaurant_menu = {
#   "Ramen" => 3,
#   "Dal Makhani" => 4,
#   "Tea" => 2
# }

# puts restaurant_menu

# puts restaurant_menu["Ramen"]

# restaurant_menu = {}

# puts restaurant_menu

# restaurant_menu["Ramen"] = 3
# restaurant_menu["Dal Makhani"] = 4
# restaurant_menu["Tea"] = 2

# puts restaurant_menu

# restaurant_menu = { "Ramen" => 3, "Dal Makhani" => 4, "Coffee" => 2 }

# restaurant_menu.each do | item, price |
#   puts "#{item}: $#{price}"
# end

# restaurant_menu = { "Ramen" => 3, "Dal Makhani" => 4, "Coffer" => 3 }

# restaurant_menu.each do | item, price |
#   restaurant_menu[item] = 1.1 * price
# end

# puts restaurant_menu

# restaurant_menu = { "Ramen" => 3, "Dal Makhani" => 4, "Coffer" => 3 }

# puts restaurant_menu.keys

# normal = Hash.new
# was_not_there = normal[:zig]
# puts "Wasn't there:"
# p was_not_there

# usually_brown = Hash.new("brown")
# pretending_to_be_there = usually_brown[:zig]
# puts "Pretending to be there:"
# p pretending_to_be_there

# chuck_norris = Hash[:punch, 99, :kick, 98, :stop_bullets_with_hands, true]
# p chuck_norris

# def artax
#   a = [:punch, 0]
#   b = [:kick, 72]
#   c = [:stops_bullets_with_hands, false]

#   key_value_pairs = [a, b, c]

#   Hash[key_value_pairs]
# end

# p artax

# puts 1.class

# puts "".class

# puts [].class

puts 1.is_a?(Integer)
puts 1.is_a?(String)
