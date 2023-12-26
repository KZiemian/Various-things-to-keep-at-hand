// fn main() {
//     println!("Hello, world!");
// }

use std::io;

fn main() {
    println!("Guess the number!");

    println!("Plase input your guess.");

    let mut guess = String::new();

    io::stdin()
	.read_line(&mut guess)
	.expect("Failed to read line");

    println!("You guessed: {guess}.");
}
