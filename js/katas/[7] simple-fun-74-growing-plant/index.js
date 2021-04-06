const growingPlant = (upSpeed, downSpeed, desiredHeight) => {
    let days = 1;
    let height = 0;

    while ((height += upSpeed) < desiredHeight) {
        days++;
        height -= downSpeed;
    }

    return days;
}

module.exports = growingPlant;
