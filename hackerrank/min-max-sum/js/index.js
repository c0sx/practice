const minMaxSum = arr => {
    let min = 0;
    let max = 0;

    const ordered = [...arr].sort();
    for (let i = 0; i < ordered.length; ++i) {
        if (i < ordered.length - 1) {
            min += ordered[i];
        }

        if (i > 0) {
            max += ordered[i];
        }
    }

    return [min, max]
}
