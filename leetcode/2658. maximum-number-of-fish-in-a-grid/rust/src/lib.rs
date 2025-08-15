fn count_fishes(grid: &mut Vec<Vec<i32>>, i: usize, j: usize) -> i32 {
    if i >= grid.len() || j >= grid[0].len() || grid[i][j] == 0 {
        return 0;
    }

    let mut amount = grid[i][j];
    grid[i][j] = 0;

    if i > 0 {
        amount += count_fishes(grid, i - 1, j)
    }

    amount += count_fishes(grid, i + 1, j);

    if j > 0 {
        amount += count_fishes(grid, i, j - 1)
    }

    amount += count_fishes(grid, i, j + 1);

    return amount;
}

pub fn find_max_fish(grid: Vec<Vec<i32>>) -> i32 {
    if grid.len() == 0 {
        return 0;
    }

    let mut result: i32 = 0;
    let mut grid = grid;

    for i in 0..grid.len() {
        for j in 0..grid[0].len() {
            let value = grid[i][j];
            if value == 0 {
                continue;
            }

            let amount = count_fishes(&mut grid, i, j);
            result = result.max(amount);
        }
    }

    return result;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works1() {
        let input = vec![
            vec![0, 2, 1, 0],
            vec![4, 0, 0, 3],
            vec![1, 0, 0, 4],
            vec![0, 3, 2, 0],
        ];

        let result = find_max_fish(input);
        assert_eq!(result, 7);
    }

    #[test]
    fn it_works2() {
        let input = vec![
            vec![1, 0, 0, 0],
            vec![0, 0, 0, 0],
            vec![0, 0, 0, 0],
            vec![0, 0, 0, 1],
        ];

        let result = find_max_fish(input);
        assert_eq!(result, 1);
    }
}
