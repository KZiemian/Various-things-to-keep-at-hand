// fn main() {
//     let y: i32 = four();

//     println!("{y}")
// }

// fn four() -> i32 {
//     4
// }

// fn main() {
//     let tup_var: (i32, &str, char) = (3, "Hello", 'a');

//     let (x, y, z) = tup_var;

//     println!("x -> {x}, y -> {y}, z -> {z}");
// }

// fn main() {
//     println!("One.");
//     println!("Two."); // println!("A")
//     println!("Three.");
//     // println!("Four.");
//     println!("Five.");
// }

// fn main() {
//     let number: i32 = 5;

//     if number < 6 {
// 	println!("{number} is less than 6.");
//     }

//     let number: i32 = 7;

//     if number < 6 {
// 	println!("We see that {number} is less than 6.");
//     }
// }

// fn main() {
//     // let number: i32 = 5;
//     let number: i32 = 7;

//     if number < 6 {
// 	println!("{number} is less than 6.");
//     } else {
// 	println!("{number} is greater or equal than 6.");
//     }
// }

// fn main() {
//     let number: i32 = 2;

//     if number == 1 {
// 	println!("number is 1");
//     } else if number == 2 {
// 	println!("number is 2")
//     } else {
// 	println!("number is diffrent from 1 and 2");
//     }
// }

// fn main() {
//     let condition: bool = true;
//     let a_int: i32 = if condition { 5 } else { 6 };

//     println!("condition -> {condition}.");
//     println!("a_int -> {a_int}.");
// }

fn main() {
    let mut condition: bool = true;
    let mut a_int: i32 = if condition { 10 } else { 13 };

    println!("condition -> {condition}.");
    println!("a_int -> {a_int}");

    condition = false;
    a_int = if condition { 11 } else { 12 };

    println!("condition -> {condition}.");
    println!("a_int -> {a_int}.");
}
