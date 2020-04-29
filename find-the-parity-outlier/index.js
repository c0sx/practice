function findOutlier(integers){
    let twoSteps = Math.abs(integers[0] % 10) % 2;
    let oneSteps = Math.abs(integers[1] % 10) % 2;
    const outlier = integers[0];

    for (let index = 2; index < integers.length; ++index) {
        const currentStep = Math.abs(integers[index] % 10) % 2;
        if (currentStep !== oneSteps) {
            if (currentStep !== twoSteps) {
                return integers[index];
            }

            return integers[index - 1];
        }

        twoSteps = oneSteps;
        oneSteps = currentStep;
    }

    return outlier;
}

console.log(findOutlier([2, 4, 0, 100, 4, 11, 2602, 36])); // Should return: 11 (the only odd number)
console.log(findOutlier([160, 3, 1719, 19, 11, 13, -21])) // Should return: 160 (the only even number)
