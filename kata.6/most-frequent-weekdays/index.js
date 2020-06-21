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

module.exports = mostFrequentDays;
