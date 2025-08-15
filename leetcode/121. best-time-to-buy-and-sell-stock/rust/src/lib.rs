use std::i32;

pub fn max_profit(prices: Vec<i32>) -> i32 {
    let mut max_profit = 0;
    let mut buying_price = i32::MAX;

    for selling_price in prices {
        if selling_price < buying_price {
            buying_price = selling_price;
            continue;
        }

        let profit = selling_price - buying_price;
        if profit > max_profit {
            max_profit = profit
        }
    }

    return max_profit;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = max_profit(vec![7, 1, 5, 3, 6, 4]);
        assert_eq!(result, 5);
    }

    #[test]
    fn it_works2() {
        let result = max_profit(vec![7, 6, 4, 3, 1]);
        assert_eq!(result, 0);
    }
}
