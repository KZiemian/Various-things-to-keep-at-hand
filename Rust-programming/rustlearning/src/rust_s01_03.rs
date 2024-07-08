// fn main() {
//     let x: i32 = 5;

//     println!("x = {x}");

//     let x: i32 = 7;

//     println!("x = {x}")
// }

// fn main() {
//     let x: i32 = 5;

//     println!("x is {x}");

//     let x: &str = "Hello world!";

//     println!("x = {x}");
// }

// fn main() {
//     let a: i32 = -5;
//     let b: u32 = 6;
//     let c: f64 = 1.23;
//     let d: bool = true;
//     let e: char = 'a';
//     let f: &str = "hello";

//     println!("a = {a}");
//     println!("b = {b}");
//     println!("c = {c}");
//     println!("d = {d}");
//     println!("e = {e}");
//     println!("f = {f}");
// }

// fn main() {
//     let a: i32 = 2 + 2;
//     let b: i32 = 4 - 1;
//     let c: i32 = 2 * 3;
//     let d: i32 = 12 / 4;
//     let e: i32 = 5 % 2;

//     println!("a = {a}");
//     println!("b = {b}");
//     println!("c = {c}");
//     println!("d = {d}");
//     println!("e = {e}");
// }

// fn main() {
//     let tup: (&str, i32, f64) = ("hello", 5, 1.2);
//     // let (x: &str, y: i32, z: f64) = tup;


//     // println!("x -> {x}, y -> {y}, z -> {z}");
//     let x: &str = tup.0;
//     let y: i32 = tup.1;
//     let z: f64 = tup.2;

//     println!("x -> {x}");
//     println!("y -> {y}");
//     println!("z -> {z}");
// }

// fn main() {
//     let a: [i32; 5] = [1, 2, 3, 4, 5];
//     let b: [&str; 2] = ["hello", "world"];
//     let c: [i32; 5] = [3; 5];

//     println!("a[0] -> {}", a[0]);
//     println!("b[1] -> {}", b[1]);
//     println!("c[2] -> {}", c[2]);
// }

// fn main() {
//     hello();
// }

// fn hello() {
//     println!("Hello World!");
// }

// fn main() {
//     hello("Yann");
// }

// fn hello(name: &str) {
//     println!("Hello {name}.");
// }

// fn main() {
//     hello("Yann", 63);
// }

// fn hello(name: &str, age: i32) {
//     println!("Hello {name}, you are {age}.");
// }

// fn main() {
//     let y: i32 = five();

//     println!("{y}");
// }

// fn five() -> i32 {
//     return 5;
// }
