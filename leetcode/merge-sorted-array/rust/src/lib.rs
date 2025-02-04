pub fn merge(nums1: &mut Vec<i32>, m: i32, nums2: &mut Vec<i32>, n: i32) {
    let mut m = m as usize;
    let mut n = n as usize;

    while n > 0 {
        let w = n + m - 1;

        if m > 0 && nums1[m - 1] > nums2[n - 1] {
            nums1[w] = nums1[m - 1];
            m -= 1;
        } else {
            nums1[w] = nums2[n - 1];
            n -= 1;
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let mut nums1 = vec![1, 2, 3, 0, 0, 0];
        let mut nums2 = vec![2, 5, 6];

        merge(&mut nums1, 3, &mut nums2, 3);

        assert_eq!(nums1, [1, 2, 2, 3, 5, 6])
    }

    #[test]
    fn it_works2() {
        let mut nums1 = vec![1];
        let mut nums2 = vec![];

        merge(&mut nums1, 1, &mut nums2, 0);

        assert_eq!(nums1, [1])
    }

    #[test]
    fn it_works3() {
        let mut nums1 = vec![0];
        let mut nums2 = vec![1];

        merge(&mut nums1, 0, &mut nums2, 1);

        assert_eq!(nums1, [1])
    }
}
