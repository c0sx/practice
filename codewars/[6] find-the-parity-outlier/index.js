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

module.exports = findOutlier;

