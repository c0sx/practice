const PER_HOUR: i32 = 221;
const MINUTES_PER_HOUR: f64 = 60.0;

pub fn production_rate_per_hour(speed: u8) -> f64 {
    let production = i32::from(speed) * PER_HOUR;
    let production = f64::from(production);

    return if speed < 5 {
        production
    } else if speed < 9 {
        production * 0.9
    } else {
        production * 0.77
    };
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    let production = production_rate_per_hour(speed);

    (production / MINUTES_PER_HOUR) as u32
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_correctly_calculates_production_rate_per_hour() {
        assert_eq!(production_rate_per_hour(6), 1193.4);
    }

    #[test]
    fn it_correctly_calculates_working_items_per_minute() {
        assert_eq!(working_items_per_minute(6), 19);
    }
}
