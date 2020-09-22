const sum = value => {
    let number = 0;

    const add = (v) => {
        number += v;

        add.valueOf = () => {
            return number;
        }

        return add;
    }

    return add(value);
}
