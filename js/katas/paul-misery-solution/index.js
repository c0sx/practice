function paul(x){
    const misery = calculateMisery(x);
    const happiness = getHappines();

    const point = happiness.find(one => misery <= one.finish);
    if (!point) {
        return "Miserable!"
    }

    return point.value;
}

const calculateMisery = x => {
    return x.reduce((acc, current) => {
        switch (current) {
            case "kata": return acc + 5;
            case "Petes kata": return acc + 10;
            case "life": return acc + 0;
            case "eating": return acc + 1;
            default: return acc;
        }
    }, 0);
}

const getHappines = () => {
    const distribution = [
        { value: "Super happy!", finish: 39 },
        { value: "Happy!", finish: 69 },
        { value: "Sad!", finish: 99 },
        { value: "Miserable!", finish: 100 }
    ];

    let left = 0;
    return distribution.map(value => {
        const start = left;
        const finish = value.finish + 1;
        left = finish;

        return {
            ...value,
            start,
            finish,
        };
    });
}
