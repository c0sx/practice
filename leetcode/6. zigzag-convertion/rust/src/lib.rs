enum Direction {
    Down,
    Up
}

pub fn convert(s: String, num_rows: i32) -> String {
    if num_rows == 1 {
        return s;
    }

    let mut result: Vec<Vec<char>> = vec![];
    for _i in 0..num_rows {
        result.push(vec![]);
    }

    let mut direction = Direction::Down;
    let mut row_index = 0;
    for c in s.chars() {
        result[row_index].push(c);

        match direction {
            Direction::Down => {
                if row_index < num_rows as usize - 1 {
                    row_index += 1;
                } else if row_index >= num_rows as usize - 1 {
                    direction = Direction::Up;
                    row_index -= 1;
                }
            },
            Direction::Up => {
                if row_index > 0 {
                    row_index -= 1;
                } else {
                    direction = Direction::Down;
                    row_index += 1;
                }
            }
        }
    }

    result.into_iter().map(|x| x.into_iter().collect::<String>()).collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works1() {
        let result = convert("PAYPALISHIRING".to_string(), 3);
        assert_eq!(result, "PAHNAPLSIIGYIR");
    }

    #[test]
    fn it_works2() {
        let result = convert("PAYPALISHIRING".to_string(), 4);
        assert_eq!(result, "PINALSIGYAHRPI");
    }

    #[test]
    fn it_works3() {
        let result = convert("A".to_string(), 1);
        assert_eq!(result, "A");
    }

    #[test]
    fn it_works4() {
        let result = convert("orangepi".to_string(), 4);
        assert_eq!(result, "opreiagn")
    }

    #[test]
    fn it_works5() {
        let result = convert("AB".to_string(), 1);
        assert_eq!(result, "AB");
    }
}
