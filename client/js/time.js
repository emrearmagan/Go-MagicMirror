function getTime(){
    let time = new Date()

    var hours = time.getHours();
    var minutes = time.getMinutes();
    var seconds = time.getSeconds();

    hours = checkForLength(hours);
    minutes = checkForLength(minutes);
    seconds = checkForLength(seconds);

    let hours_minutes = hours + ":" + minutes;

    document.getElementById("hours-minutes").innerHTML = hours_minutes;
    document.getElementById("seconds").innerHTML = seconds.toString();
}

function getDate() {
    const week = new Array('Sonntag', 'Montag', 'Dienstag', 'Mittwoch', 'Donnerstag', 'Freitag', 'Samstag');
    const months = new Array('Januar', 'Februar', 'März', 'April', 'Mai', 'Juni', 'Juli', 'August', 'September', 'Oktober', 'November', 'Dezember');

    var date = new Date();
    var currentDay = week[date.getDay()];
    var month = months[date.getMonth()]

    var day = date.getDate();
    var year = date.getFullYear();

    let todaysDate = currentDay + ", " + day + ". " + month + " " + year;
    document.getElementById("date").innerHTML = todaysDate;
}

function checkForLength(x) {
    if(x < 10){
        x = "0" + x;
    }

    return x;
}

getTime();
getDate();

setInterval(getDate,1000*60*60*24); //@todo
setInterval(getTime,1000);