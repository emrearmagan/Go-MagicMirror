function Forecast() {
    let req;
    req = new XMLHttpRequest();
    req.onreadystatechange = async function () {
        if (req.readyState === 4 && req.status === 200) {
            //@todo
            let loading = document.querySelector(".loading");
            loading.parentNode.removeChild(loading);

            let forecastList = document.querySelector(".weather_forecast_list");
            let elements = forecastList.children;
            for (let i = 0; i < elements.length; i++) {
                let s = elements[i];
                removeElements(s)
            }
            let obj = (function (raw) {
                try {
                    return JSON.parse(raw);
                } catch (err) {
                    return false;
                }
            })(req.responseText);
            if (!obj) {
                error(forecastList, req.responseText);
                return false;
            }
            let forecastDate = document.querySelector(".forecast_date.forecast_list");
            let forecastIcon = document.querySelector(".forecast_icon.forecast_list");
            let forecastRain = document.querySelector(".rain.forecast_list");
            let forecastTopDegree = document.querySelector(".topDegree.forecast_list");
            let forecastLowDegree = document.querySelector(".lowDegree.forecast_list");
            let forecastHumidity = document.querySelector(".humidity.forecast_list");
            for (let i = 0; i < obj.list.length; i++) {
                //Create the elements
                let date = document.createElement("LI");
                let icon = document.createElement("img");
                let rain = document.createElement("LI");
                let topDegree = document.createElement("LI");
                let lowDegree = document.createElement("LI");
                let humidity = document.createElement("LI");

                icon.setAttribute("class", "forecast_icon_img");
                icon.src = "/client/img/weather_icon/" + obj.list[i].weather[0].icon + ".png";
                Unix_timestamp(1412743274)

                let dateText = document.createTextNode(Unix_timestamp(obj.list[i].dt, i));
                //@todo
                let rainText = document.createTextNode("10%");
                var topDegreeText = document.createTextNode(Math.floor(obj.list[i].main.temp_max) + "°");
                var lowDegreeText = document.createTextNode(Math.floor(obj.list[i].main.temp_min) + "°");
                var humidityText = document.createTextNode(Math.floor(obj.list[i].main.humidity) + "%");

                date.appendChild(dateText);
                rain.appendChild(rainText);
                topDegree.appendChild(topDegreeText);
                lowDegree.appendChild(lowDegreeText);
                humidity.appendChild(humidityText);

                forecastDate.appendChild(date);
                forecastIcon.appendChild(icon);
                forecastRain.appendChild(rain);
                forecastTopDegree.appendChild(topDegree);
                forecastLowDegree.appendChild(lowDegree);
                forecastHumidity.appendChild(humidity);

                //Return only the first 3 times and afterwards the days
                // @todo server should only send needed data
                if (i >= 3) {
                    i = i + 8
                }
                if (i > 20) {
                    return
                }

                await sleep(200);

            }
        } else {
            return;
        }
    }
    ;
    req.open("GET", "http://localhost:8080/forecast");
    req.send();
}

function removeElements(element) {
    let node = element.lastElementChild;
    while (node) {
        element.removeChild(node);
        node = element.lastElementChild;
    }
}

/**
 * @return {string}
 */
function Unix_timestamp(t, i) {
    if (i >= 3) {
        var day = new Map([[0 , "Mo"], [1 ,"Di"] ,[2, "Mi"],[3, "Do"],[4, "Fr"],[5, "Sa"],[6, "So"]]);
        let dt = new Date(t * 1000);
        return day.get(dt.getDay());
    }
    let dt = new Date(t * 1000);
    let hr = dt.getHours();
    let m = "0" + dt.getMinutes();
    return hr + ':' + m.substr(-2);
}