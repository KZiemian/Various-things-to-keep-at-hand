#!/home/kamil/.rbenv/shims/ruby

# "Hello World!"

# puts "Hello World!"
# puts 3**2
# puts Math.sqrt(3)

# def hi
#   puts "Hello World!"
# end

# hi()

# class Greeter
#   def initialize(name = "World")
#     @name = name
#   end

#   def say_hi
#     puts "Hi #{@name}!"
#   end

#   def say_bye
#     puts "Bye #{@name}, come back soon."
#   end
# end

# greeter = Greeter.new("Pat")

# greeter.say_hi
# greeter.say_bye

# puts greeter.@name

# puts Greeter.instance_methods
# puts "Druga lista metod."
# puts Greeter.instance_methods(false)

# puts greeter.respond_to?("name")
# puts greeter.respond_to?("say_hi")
# puts greeter.respond_to?("to_s")

# class Greeter
#   attr_accessor :name
# end

# greeter = Greeter.new("Andy")
# puts greeter.respond_to?("name")
# puts greeter.respond_to?("name=")

# greeter.say_hi

# greeter.name="Betty"

# puts greeter

# puts greeter.name

# greeter.say_hi

class MegaGreeter
  attr_accessor :names

  # Create the object
  def initialize(names = "World")
    @names = names
  end

  # Say hi to everybody
  def say_hi
    if @names.nil?
      puts "..."
    elsif @names.respond_to?("each")
      # @names is a list of some kind, iterate!
      @names.each do |name|
        puts "Hello #{name}!"
      end
    else
      puts "Hello #{@names}!"
    end
  end

    # Say bye to everybody
  def say_bye
    if @names.nil?
      puts "..."
    elsif @names.respond_to?("join")
      # Join the list elements with commas.
      puts "Goodbye #{@names.join(", ")}. Come back soon!"
    else
      puts "Goodbye #{@names}. Come back soon!"
    end
  end
end

if __FILE__ == $0
  mg = MegaGreeter.new
  mg.say_hi
  mg.say_bye

  # Change name to be "Zeke"
  mg.names = "Zeke"
  mg.say_hi
  mg.say_bye

  # Change the name to an array of names
  mg.names = ["Albert", "Brenda", "Charles",
              "Dave", "Englebert"]
  mg.say_hi
  mg.say_bye

  # Change to nil
  mg.names = nil
  mg.say_hi
  mg.say_bye
end
