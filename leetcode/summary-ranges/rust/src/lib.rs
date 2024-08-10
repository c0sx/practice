pub fn summary_ranges(nums: Vec<i32>) -> Vec<String> {
    if nums.len() == 0 {
        return vec![];
    }

    let mut result: Vec<String> = vec![];
    let mut start = nums[0];
    let mut end = nums[0];
    let mut index = 1;

    while index < nums.len() {
        let current = nums[index];
        let prev = nums[index - 1];
        if current == prev + 1 {
            end += 1;
        } else {
            let str = if start == end {
                format!("{}", start)
            } else {
                format!("{}->{}", start, end)
            };

            result.push(str);
            start = current;
            end = current;
        }

        index += 1;
    }

    let str = if start == end {
        format!("{}", start)
    } else {
        format!("{}->{}", start, end)
    };

    result.push(str);

    return result;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = summary_ranges(vec![0, 1, 2, 4, 5, 7]);
        let expected = vec!["0->2", "4->5", "7"];

        assert_eq!(result, expected);
    }

    #[test]
    fn it_works2() {
        let result = summary_ranges(vec![0, 2, 3, 4, 6, 8, 9]);
        let expected = vec!["0", "2->4", "6", "8->9"];

        assert_eq!(result, expected);
    }

    #[test]
    fn it_works3() {
        let result = summary_ranges(vec![]);
        let expected: Vec<String> = vec![];

        assert_eq!(result, expected);
    }
}
