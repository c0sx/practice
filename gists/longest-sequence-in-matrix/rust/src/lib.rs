#[derive(Debug)]
struct Point {
    i: usize,
    j: usize,
    value: i32,
    id: u32
}

pub fn longest_sequence<const N: usize>(matrix: [[i32; N]; N]) -> u32 {
    let mut sequence = 0;
    let mut points = Vec::with_capacity(N);

    for (i, row) in matrix.iter().enumerate() {
        points.push(Vec::with_capacity(N));

        for (j, cell) in row.iter().enumerate() {
            let id = if *cell == 1 {
                let id = get_id(i, j, &points);
                if id == 0 {
                    sequence += 1;
                    sequence
                } else {
                    id
                }
            } else {
                0
            };

            let point = Point {
                i,
                j,
                value: *cell,
                id
            };

            points[i].push(point);
        }
    }

    // дан двумерный массив из нулей и единиц,
    // нужно посчитать количество последовательностей единиц в этой матрице,
    // то есть если единицы соседствуют по вертикали или горизонтали
    // (но не по диагонали) - это одна последовательность

    sequence
}

fn get_id(i: usize, j: usize, points: &Vec<Vec<Point>>) -> u32 {
    let left = if i > 0 {
        points.get(i - 1).and_then(|r| r.get(j))
    } else {
        None
    };

    let top = if j > 0 {
        points.get(i).and_then(|r| r.get(j - 1))
    } else {
        None
    };

    let value = vec![top, left]
        .into_iter()
        .filter(|x| x.is_some())
        .map(|x| x.unwrap())
        .find(|x| x.value == 1);

    match value {
        Some(p) => p.id,
        None => 0
    }
}

#[cfg(test)]
mod tests {
    use super::longest_sequence;

    #[test]
    fn it_works() {
        let value1 = (
            [
                [1, 0, 1],
                [0, 1, 0],
                [1, 0, 1],
            ],
            5
        );

        let value2 = (
            [
                [1, 1, 1],
                [0, 0, 0],
                [1, 1, 1]
            ],
            2
        );

        let value3 = (
            [
                [1, 0, 1],
                [1, 0, 1],
                [1, 1, 1]
            ],
            1
        );

        let value4 = (
            [
                [1, 0, 1, 0, 1],
                [1, 0, 1, 0, 1],
                [1, 0, 1, 0, 1],
                [1, 0, 1, 0, 1],
                [1, 1, 1, 1, 1]
            ],
            1
        );

        assert_eq!(longest_sequence(value1.0), value1.1);
        assert_eq!(longest_sequence(value2.0), value2.1);
        assert_eq!(longest_sequence(value3.0), value3.1);
        assert_eq!(longest_sequence(value4.0), value4.1);
    }
}
