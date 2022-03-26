use nth_prime::{nth_prime};

fn main() {
    println!("find a prime with index {} => {}", 0, nth_prime(0));
    println!("find a prime with index {} => {}", 1, nth_prime(1));
    println!("find a prime with index {} => {}", 2, nth_prime(2));
    println!("find a prime with index {} => {}", 10001, nth_prime(10001));
}
