fn spin_words(sentence: &str) -> String {
    let words = sentence.split(" ").collect::<Vec<&str>>();
    let mut result: Vec<String> = vec![];

    for word in words {
        if word.chars().count() >= 5 {
            let reversed = word.chars().rev().collect::<String>();
            result.push(reversed);
        } else {
            result.push(String::from(word));
        }
    }

    return result.join(" ");
}

#[cfg(test)]
mod tests {
    use super::spin_words;

    #[test]
    fn examples() {
        assert_eq!(spin_words("Welcome"), "emocleW");
        assert_eq!(spin_words("Hey fellow warriors"), "Hey wollef sroirraw");
        assert_eq!(spin_words("This is a test"), "This is a test");
        assert_eq!(spin_words("This is another test"), "This is rehtona test");
        assert_eq!(
            spin_words("You are almost to the last test"),
            "You are tsomla to the last test"
        );
        assert_eq!(
            spin_words("Just kidding there is still one more"),
            "Just gniddik ereht is llits one more"
        );
        assert_eq!(
            spin_words("Seriously this is the last one"),
            "ylsuoireS this is the last one"
        );
    }
}
