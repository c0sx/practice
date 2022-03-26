#[derive(Clone, PartialEq, Debug)]
pub enum LogLevel {
    Info,
    Warning,
    Error,
}

fn log(level: LogLevel, message: &str) -> String {
    let mut log = match level {
        LogLevel::Info => String::from("[INFO]: "),
        LogLevel::Warning => String::from("[WARNING]: "),
        LogLevel::Error => String::from("[ERROR]: ")
    };

    log.push_str(message);

    log
}
pub fn info(message: &str) -> String {
    log(LogLevel::Info, message)
}
pub fn warn(message: &str) -> String {
    log(LogLevel::Warning, message)
}
pub fn error(message: &str) -> String {
    log(LogLevel::Error, message)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(info("Timezone changed"), "[INFO]: Timezone changed");
    }
}
