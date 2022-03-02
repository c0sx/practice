const plusMinus = (arr) => {
    let positive = 0;
    let negative = 0;
    let zeros = 0;

    arr.forEach(item => {
        if (item > 0) {
            positive += 1;
        } else if (item < 0) {
            negative += 1
        } else {
            zeros += 1
        }
    });

    return [
        positive / arr.length,
        negative / arr.length,
        zeros / arr.length
    ]
}
