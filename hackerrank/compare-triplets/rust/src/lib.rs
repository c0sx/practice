pub fn compare_triplets(a_score: [i32; 3], b_score: [i32; 3]) -> [i32; 2] {
    let mut result = [0, 0];

    let mut i = 0;
    while i < 3 {
        let a = a_score[i];
        let b = b_score[i];

        if a > b {
            result[0] += 1
        } else if b > a {
            result[1] += 1
        }

        i += 1;
    }

    result
}

#[cfg(test)]
mod tests {
    use super::compare_triplets;

    #[test]
    fn it_works() {
        assert_eq!(compare_triplets([3, 2, 1], [1, 2, 3]), [1, 1]);
        assert_eq!(compare_triplets([1, 1, 1], [1, 1, 1]), [0, 0]);
        assert_eq!(compare_triplets([3, 3, 3], [1, 1, 1]), [3, 0]);
        assert_eq!(compare_triplets([1, 1, 1], [3, 3, 3]), [0, 3]);
    }
}
