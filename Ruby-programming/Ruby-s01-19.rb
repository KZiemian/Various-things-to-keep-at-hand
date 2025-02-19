#!/home/kamil/.rbenv/shims/ruby
# coding: utf-8

# class Book
#   attr_accessor :title, :author, :pages

#   def initialize(name)
#     # puts "Creating Book"
#     puts ("Hello " + name)
#   end
# end

# book1 = Book.new("Rowling")
# book1.title = "Harry Potter"
# book1.author = "J.K. Rowling"
# book1.pages = 400

# book2 = Book.new("Tolkien")
# book2.title = "Lord of the Rings"
# book2.author = "Tolkien"
# book2.pages = 500

# class Book
#   attr_accessor :title, :author, :pages

#   def initialize(title, author, pages)
#     @title = title
#     @author = author
#     @pages = pages
#   end

# end

# book1 = Book.new("Tolkien", "Lord of the Rings", 500)
# book2 = Book.new("J.K. Rowling", "Harry Potter", 400)

# puts book1.title
# puts book2.author

# class Student
#   attr_accessor :name, :major, :gpa

#   def initialize(name, major, gpa)
#     @name = name
#     @major = major
#     @gpa = gpa
#   end

#   def has_honors
#     if @gpa >= 3.5
#       return true
#     else
#       return false
#     end
#   end
# end

# student1 = Student.new("Jim", "Buisness", 2.6)
# student2 = Student.new("Pam", "Art", 3.6)

# puts student1.has_honors
# puts student2.has_honors

class Question
  attr_accessor :prompt, :answer

  def initialize(prompt, answer)
    @prompt = prompt
    @answer = answer
  end
end

p1 = "What color are apples?\n(a) Red.\n(b) Purple.\n(c) Orange.\n"
p2 = "What color are bananas?\n(a) Pink.\n(b) Green.\n(c) Yellow.\n"
p3 = "What color are pears?\n(a) Yellow.\n(b) Green.\n(c) Orange.\n"

questions = [
  Question.new(p1, "a"),
  Question.new(p2, "c"),
  Question.new(p3, "b")
]

def run_test(questions)
  answer = ""
  score = 0

  for question in questions
    puts question.prompt
    answer = gets.chomp()

    if answer == question.answer
      score += 1

    end
  end

  puts ("You got " + score.to_s + "/" + questions.length().to_s)
end

run_test(questions)
