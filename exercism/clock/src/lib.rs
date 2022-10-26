use std::fmt::{Debug, Display, Formatter};

pub struct Clock {
    hours: i32,
    minutes: i32,
}

impl Debug for Clock {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}", self.to_string())
    }
}

impl Display for Clock {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(f, "{:0>2}:{:0>2}", self.hours, self.minutes)
    }
}

impl PartialEq for Clock {
    fn eq(&self, other: &Self) -> bool {
        self.hours == other.hours && self.minutes == other.minutes
    }

    fn ne(&self, other: &Self) -> bool {
        !self.eq(other)
    }
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let (h, m) = Self::normalize(hours, minutes);

        Clock {
            hours: h,
            minutes: m
        }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        let (h, m) = Self::normalize(self.hours, self.minutes + minutes);

        Clock {
            hours: h,
            minutes: m
        }
    }

    fn normalize(hours: i32, minutes: i32) -> (i32, i32) {
        let total_minutes = hours * 60 + minutes;
        let hours = (total_minutes as f32 / 60.0).floor();
        let minutes = total_minutes as f32 - hours * 60.0;

        let hours = hours % 24.0;
        let hours = if hours < 0.0 {
            24.0 + hours
        } else {
            hours
        };

        (hours as i32, minutes as i32)
    }
}


#[cfg(test)]
mod tests {
    use crate::Clock;

    #[test]
    fn should_add_minutes() {
        let clock = Clock::new(15, 32);
        let clock2 = clock.add_minutes(10);
        let clock3 = clock.add_minutes(30);

        assert_eq!(clock2.hours, 15);
        assert_eq!(clock2.minutes, 42);
        assert_eq!(clock3.hours, 16);
        assert_eq!(clock3.minutes, 2);
    }

    #[test]
    fn should_subtract_minutes() {
        let clock = Clock::new(15, 32);
        let clock2 = clock.add_minutes(-10);
        let clock3 = clock.add_minutes(-70);

        assert_eq!(clock2.hours, 15);
        assert_eq!(clock2.minutes, 22);
        assert_eq!(clock3.hours, 14);
        assert_eq!(clock3.minutes, 22);
    }

    #[test]
    fn should_omit_days() {
        let clock = Clock::new(26, 0);

        assert_eq!(clock.hours, 2);
        assert_eq!(clock.minutes, 0);

        let clock2 = Clock::new(-25, 0);

        assert_eq!(clock2.hours, 23);
        assert_eq!(clock2.minutes, 0);
    }

    #[test]
    fn should_equal_clocks() {
        let clock1 = Clock::new(10, 0);
        let clock2 = Clock::new(10, 0);

        assert!(clock1 == clock2);
    }

    #[test]
    fn should_stringify_value() {
        assert_eq!(Clock::new(8, 15).to_string(), "08:15");
        assert_eq!(Clock::new(1, -4820).to_string(), "16:40");
    }
}
