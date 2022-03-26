use rust::multiples_3_or_5;

pub fn main() {
    let size = 1_000;
    let r = multiples_3_or_5(size);

    println!("{}", r);
}
