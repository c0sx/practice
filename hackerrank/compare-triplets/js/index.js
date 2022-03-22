const compare = (aScore, bScore) => {
    const result = [0, 0];

    for (let i = 0; i < 3; ++i) {
        const a = aScore[i];
        const b = bScore[i];

        if (a > b) {
            result[0] += 1;
        } else if (b > a) {
            result[1] += 1;
        }
    }

    return result;
}
