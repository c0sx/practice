use std::{collections::{HashMap, HashSet}};

fn link_sibling(hash_map: &mut HashMap<i32, Vec<i32>>, from: i32, to: i32) {
    let adjacents = hash_map.entry(from).or_insert(vec![]);

    adjacents.push(to);
}

fn traverse(key: i32, result: &mut Vec<i32>, visited: &mut HashSet<i32>, hash_map: &HashMap<i32, Vec<i32>>) {
    if visited.contains(&key) {
        return;
    }

    visited.insert(key);
    result.push(key);

    let adjacents = hash_map.get(&key).unwrap();

    for one in adjacents {
        traverse(*one, result, visited, hash_map)
    }
}

fn find_root(hash_map: &HashMap<i32, Vec<i32>>) -> i32 {
    let root = hash_map
        .into_iter()
        .find(|(_, vec)| vec.len() == 1)
        .unwrap();

    *root.0
}

pub fn restore_array(adjacent_pairs: Vec<Vec<i32>>) -> Vec<i32> {
    let mut hash_map = HashMap::<i32, Vec<i32>>::new();

    for pair in adjacent_pairs {
        let left = pair[0];
        let right = pair[1];

        link_sibling(&mut hash_map, left, right);
        link_sibling(&mut hash_map, right, left);
    }


    let root = find_root(&hash_map);
    let mut result = Vec::new();
    let mut visited = HashSet::new();

    traverse(root, &mut result, &mut visited, &hash_map);

    result
}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let input = vec![vec![2, 1], vec![3, 4], vec![3, 2]];
        let result = restore_array(input);
        
        assert_eq!(result, vec![1, 2, 3, 4]);
    }
}
