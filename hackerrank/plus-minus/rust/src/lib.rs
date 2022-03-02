pub fn plus_minus<const N: usize>(arr: [i32; N]) -> [f64; 3] {
    let mut positive = Vec::<i32>::new();
    let mut negative = Vec::<i32>::new();
    let mut zero = Vec::<i32>::new();

    for item in arr.iter() {
        if *item > 0 {
            positive.push(*item);
        } else if *item < 0 {
            negative.push(*item)
        } else {
            zero.push(*item)
        }
    }

    let length = N as f64;

    return [
        positive.len() as f64 / length,
        negative.len() as f64 / length,
        zero.len() as f64 / length
    ]
}

#[cfg(test)]
mod tests {
    use super::plus_minus;

    #[test]
    fn it_works() {
        let input = [1, 1, 0, -1, -1];
        let expect = [0.4, 0.4, 0.2];

        assert_eq!(plus_minus(input), expect);
    }
}
