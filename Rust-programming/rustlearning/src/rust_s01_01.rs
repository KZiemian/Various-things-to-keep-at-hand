// fn main() {
//     println!("Hello, world!");
// }

// use std::io;

// fn main() {
//     println!("Guess the number!");

//     println!("Plase input your guess.");

//     let mut guess = String::new();

//     io::stdin()
// 	.read_line(&mut guess)
// 	.expect("Failed to read line");

//     println!("You guessed: {guess}.");
// }

// use rand::Rng;
// use std::cmp::Ordering;
// use std::io;

// fn main() {
//     println!("Guess the number!");

//     let secret_number = rand::thread_rng().gen_range(1..=100);

//     loop {
// 	println!("Please input your guess.");

// 	let mut guess = String::new();

// 	io::stdin()
// 	    .read_line(&mut guess)
// 	    .expect("Failed to read line.");

// 	let guess: u32 = match guess.trim().parse() {
// 	    Ok(num) => num,
// 	    Err(_) => continue,
// 	};

// 	println!("You guessed: {guess}.");

// 	match guess.cmp(&secret_number) {
// 	    Ordering::Less => println!("Too small!"),
// 	    Ordering::Greater => println!("Too big!"),
// 	    Ordering::Equal => {
// 		println!("You win!");

// 		break;
// 	    }
// 	}
//     }
// }

// fn main() {
//     println!("Hello, World!");
// }

// fn main() {
//     // let x: i32 = 5;

//     // println!("x = {x}");

//     let x: i32 = 5;

//     println!("x = {x}");

//     x = 7;

//     println!("x = {x}")
// }

// fn main() {
//     let mut x: i32 = 5;

//     println!("x = {x}");

//     x = 6;

//     println!("x = {x}")
// }

// fn main() {
//     const DURATION: i32 = 3;

//     println!("DURATION = {DURATION}");
// }

fn main() {
    const DURATION: i32 = 3;

    println!("DURATION = {DURATION}");

    DURATION = 3;

    println!("DURATION = {DURATION}");
}
