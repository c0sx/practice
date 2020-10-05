const regex = /^(\d{4}|\d{6})$/;
function validatePIN (pin) {
    return regex.test(pin);
}

module.exports = validatePIN
