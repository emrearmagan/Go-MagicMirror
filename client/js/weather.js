function CurrentWeather() {
    let req;
    req = new XMLHttpRequest();
    req.onreadystatechange = async function () {
        if (req.readyState == 4 && req.status == 200) {
            let weatherToday = document.getElementsByClassName(".weather_today");
            //Convert json in an object and fill the div in a list
            let obj = (function (raw) {
                try {
                    return JSON.parse(raw);
                } catch (err) {
                    return false;
                }
            })(req.responseText);
            if (!obj) {
                error(weatherToday, req.responseText);
                return false;
            }

            //set information
            document.getElementById("city").innerText = obj.name;
            document.getElementById("weather_icon").src = "/client/img/weather_icon/" + obj.weather[0].icon + ".png";
            document.getElementById("weather_desc").innerText = obj.weather[0].description;
            document.getElementById("degree").innerText = Math.floor(obj.main.temp) + "°";
        }
    };
    req.open("GET", "http://localhost:8080/weather");
    req.send();
}
