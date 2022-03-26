pub fn expected_minutes_in_oven() -> i32 {
    40
}

pub fn remaining_minutes_in_oven(m: i32) -> i32 {
    return expected_minutes_in_oven() - m;
}

pub fn preparation_time_in_minutes(layers: i32) -> i32 {
    return layers * 2;
}

pub fn elapsed_time_in_minutes(layers: i32, m: i32) -> i32 {
    return preparation_time_in_minutes(layers) + m;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    pub fn it_works() {
        assert_eq!(expected_minutes_in_oven(), 40);
        assert_eq!(remaining_minutes_in_oven(30), 10);
        assert_eq!(preparation_time_in_minutes(2), 4);
        assert_eq!(elapsed_time_in_minutes(3, 20), 26);
    }
}
