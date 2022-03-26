pub fn even_fibonacci_numbers(limit: i64) -> i64 {
    let mut sum: i64 = 0;
    let mut sequence = vec![1, 1];

    while sum < limit {
        let index = sequence.len() - 1;
        let next = sequence[index] + sequence[index - 1];

        if next > limit {
            break;
        }

        if next % 2 == 0 {
            sum += next;
        }

        sequence.push(next);
    }

    println!("{:?}", sequence);
    sum
}
