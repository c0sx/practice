function mostFrequentDays(year) {
    const lastWeekdayOfPastYear = getLastWeekdayOfPastYears(year)
    const firstWeekdayOfCurrentYear = getFirstWeekdayOfCurrentYear(lastWeekdayOfPastYear);
    const daysInCurrentYear = getDaysInYear(year);
    const remainsDays = getRemainDays(daysInCurrentYear);
    const weekdays = getRemainWeekdays(firstWeekdayOfCurrentYear, remainsDays)

    return weekdays.sort().map(renderWeekday)
}

function getLastWeekdayOfPastYears(year) {
    let totalDays = 0
    for (let i = 1; i < year; ++i) {
        totalDays += getDaysInYear(i)
    }

    return getRemainDays(totalDays)
}

function getRemainDays(days) {
    const fullWeeks = Math.trunc(days / 7);
    return days - fullWeeks * 7;
}

function getDaysInYear(year) {
    if (isLeap(year)) {
        return 366
    }

    return 365
}

function getFirstWeekdayOfCurrentYear(weekday) {
    if (weekday === 7) {
        return 1;
    }

    return weekday + 1;
}

function getRemainWeekdays(firstDay, count) {
    const weekdays = [];
    for (let i = 0; i < count; ++i) {
        const next = firstDay + i;
        if (next > 7) {
            weekdays.push(Math.trunc(next / 7))
        }
        else {
            weekdays.push(next)
        }
    }

    return weekdays
}

function renderWeekday(weekday) {
    switch (weekday) {
        case 1: return "Monday";
        case 2: return "Tuesday";
        case 3: return "Wednesday";
        case 4: return "Thursday";
        case 5: return "Friday";
        case 6: return "Saturday";
        case 7: return "Sunday"
    }
}

function isLeap(year) {
    if (year % 400 === 0) {
        return true
    }

    if (year % 100 === 0) {
        return false;
    }

    return year % 4 === 0;
}

console.log(mostFrequentDays(1084)) // , ['Tuesday', 'Wednesday'], "Should be: ['Tuesday', 'Wednesday']");
console.log(mostFrequentDays(1167)) // , ['Sunday'], "Should be: ['Sunday']");
console.log(mostFrequentDays(1216)) // , ['Friday', 'Saturday'], "Should be: ['Friday', 'Saturday']");
console.log(mostFrequentDays(1492)) // , ['Friday', 'Saturday'], "Should be: ['Friday', 'Saturday']");
console.log(mostFrequentDays(1770)) // , ['Monday'], "Should be: ['Monday']");
console.log(mostFrequentDays(1785)) // , ['Saturday'], "Should be: ['Saturday']");
console.log(mostFrequentDays(212)) // , ['Wednesday', 'Thursday'], "Should be: ['Wednesday', 'Thursday']");
console.log(mostFrequentDays(1901)) //   , ['Tuesday'], "Should be: ['Tuesday']");
console.log(mostFrequentDays(2135)) // , ['Saturday'], "Should be: ['Saturday']");
console.log(mostFrequentDays(3043)) // , ['Sunday'], "Should be: ['Sunday']");
console.log(mostFrequentDays(2001)) // , ['Monday'], "Should be: ['Monday']");
console.log(mostFrequentDays(3150)) // , ['Sunday'], "Should be: ['Sunday']");
console.log(mostFrequentDays(3230)) // , ['Tuesday'], "Should be: ['Tuesday']");
console.log(mostFrequentDays(328)) // , ['Monday', 'Sunday'], "Should be: ['Monday', 'Sunday']");
console.log(mostFrequentDays(2016)) // , ['Friday', 'Saturday'], "Should be: ['Friday', 'Saturday']");
console.log(mostFrequentDays(334)) // , ['Monday'], "Should be: ['Monday']");
console.log(mostFrequentDays(1986)) // , ['Wednesday'], "Should be: ['Wednesday']");
console.log(mostFrequentDays(3361)) // , ['Thursday'], "Should be: ['Thursday']");
console.log(mostFrequentDays(5863)) // , ['Thursday'], "Should be: ['Thursday']");
console.log(mostFrequentDays(1910)) // , ['Saturday'], "Should be: ['Saturday']");
console.log(mostFrequentDays(1968)) // , ['Monday', 'Tuesday'], "Should be: ['Monday', 'Tuesday']");
console.log(mostFrequentDays(7548)) // , ['Thursday', 'Friday'], "Should be: ['Thursday', 'Friday']");
console.log(mostFrequentDays(8116)) // , ['Wednesday', 'Thursday'], "Should be: ['Wednesday', 'Thursday']");
console.log(mostFrequentDays(9137)) // , ['Friday'], "Should be: ['Friday']");
console.log(mostFrequentDays(1794)) // , ['Wednesday'], "Should be: ['Wednesday']");
