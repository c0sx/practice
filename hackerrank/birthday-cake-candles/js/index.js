const birthdayCakeCandles = input => {
    let value = 0;
    let counter = 0;

    for (const i of input) {
        if (i > value) {
            value = i;
            counter = 1;
        } else if (i === value) {
            counter += 1;
        }
    }

    return counter;
}
