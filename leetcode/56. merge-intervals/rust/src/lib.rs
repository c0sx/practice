use std::cmp;

pub fn merge(intervals: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
    if intervals.len() == 1 {
        return intervals;
    }

    let mut intervals = intervals;
    intervals.sort();

    let mut result = vec![];
    let mut prev = intervals[0].to_vec();
    let mut index = 1;

    while index < intervals.len() {
        let current = intervals[index].to_vec();

        if current[0] <= prev[1] {
            prev[1] = cmp::max(prev[1], current[1]);
        } else {
            result.push(prev.to_vec());
            prev = current.to_vec();
        }

        index += 1
    }

    result.push(prev.to_vec());

    result.to_vec()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = merge(vec![vec![1, 3], vec![2, 6], vec![8, 10], vec![15, 18]]);
        assert_eq!(result, [[1, 6], [8, 10], [15, 18]]);
    }

    #[test]
    fn it_works2() {
        let result = merge(vec![vec![1, 4], vec![4, 5]]);
        assert_eq!(result, [[1, 5]]);
    }

    #[test]
    fn it_works3() {
        let result = merge(vec![vec![1, 3]]);
        assert_eq!(result, [[1, 3]]);
    }

    #[test]
    fn it_works4() {
        let result = merge(vec![vec![1, 4], vec![0, 4]]);
        assert_eq!(result, [[0, 4]]);
    }

    #[test]
    fn it_works5() {
        let result = merge(vec![vec![1, 4], vec![0, 0]]);
        assert_eq!(result, [[0, 0], [1, 4]]);
    }

    #[test]
    fn it_works6() {
        let result = merge(vec![vec![1, 4], vec![0, 2], vec![3, 5]]);
        assert_eq!(result, [[0, 5]]);
    }
}
