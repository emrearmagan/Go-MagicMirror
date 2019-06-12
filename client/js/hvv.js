function HVV() {
    let req;
    req = new XMLHttpRequest();
    req.onreadystatechange = async function () {
        if (req.readyState === 4 && req.status === 200) {
            //Get the Node and clear it
            let hvvList = document.querySelector(".hvv-list");
            let child = hvvList.lastElementChild;
            while (child) {
                hvvList.removeChild(child);
                child = hvvList.lastElementChild;
            }
            //Convert json in an object and fill the div in a list
            let obj = (function (raw) {
                try {
                    return JSON.parse(raw);
                } catch (err) {
                    return false;
                }
            })(req.responseText);

            if (!obj) {
                error(hvvList, req.responseText);
                return false;
            }

            //Listing all elements
            for (let i = 0; i < obj.realtimeSchedules.length; i++) {
                //Thats where each row is gonna be store
                let hvvRow = document.createElement("ul");
                hvvRow.setAttribute("class", "hvv row");

                //create the elements inside of ul hvv row


                let t = obj.realtimeSchedules[i].time;
                let totalTime = document.createElement("LI");
                totalTime.setAttribute("class", "hvv TotalTime");
                let textTotalTime = document.createTextNode(formatTime(t));
                totalTime.appendChild(textTotalTime);


                let time = document.createElement("LI");


                // List the Icons and transit numbers
                let scheduleElements = document.createElement("div");
                scheduleElements.setAttribute("class", "hvv ScheduleElements");
                for (let z = 0; z < obj.realtimeSchedules[i].scheduleElements.length; z++) {
                    let elements = obj.realtimeSchedules[i].scheduleElements[z];
                    if (z == 0) {
                        time.setAttribute("class", "hvv Time");
                        let textTime = document.createTextNode(timeLine(elements.from.depTime.time, t, "hh:mm"));
                        time.appendChild(textTime);
                    }

                    //Create the elements inside hvv ScheduleElements
                    let hvvLine = document.createElement("LI");
                    hvvLine.setAttribute("class", "hvv line");
                    let hvvIcon = document.createElement("img");
                    hvvIcon.setAttribute("class", "hvv icon");
                    if (elements.line.type.shortInfo != "") {
                        hvvIcon.src = "/client/img/hvv/" + elements.line.type.shortInfo + ".png";
                    } else {
                        hvvIcon.src = "/client/img/hvv/" + elements.line.type.simpleType + ".png";
                    }

                    let hvvArrow = document.createElement("img");
                    //only if next element exist add arrow
                    if (obj.realtimeSchedules[i].scheduleElements[z + 1] != null) {
                        hvvArrow.setAttribute("class", "hvv arrow");
                        hvvArrow.src = "/client/img/hvv/arrow.png";
                    }

                    let textHvvLine = document.createTextNode(elements.line.name);
                    hvvLine.appendChild(textHvvLine);

                    scheduleElements.appendChild(hvvLine);
                    scheduleElements.appendChild(hvvIcon);
                    scheduleElements.appendChild(hvvArrow);

                }

                hvvRow.appendChild(time);
                hvvRow.appendChild(scheduleElements);
                hvvRow.appendChild(totalTime);

                hvvList.appendChild(hvvRow);
                await sleep(200);
            }
        }
    };
    req.open("GET", "http://localhost:8080/hvv");
    req.send();
}

function formatTime(time) {
    let hours = Math.floor((time / 60));
    let minutes = (time % 60);

    if (minutes < 10) {
        minutes = "0" + minutes;
    }
    return hours + ":" + minutes + "h"
}

//return the time from start to destinaion in format like hh:mm - hh:mm
function timeLine(startTime, time, format) {
    var start = new Date();
    if (format == "hh:mm") {
        start.setHours(startTime.substr(0, startTime.indexOf(":")));
        start.setMinutes(startTime.substr(startTime.indexOf(":") + 1));
        let startHours = start.getHours();
        let startMinutes = start.getMinutes();
        if (startMinutes < 10) {
            startMinutes = "0" + startMinutes
        }

        let arrival = addMinutes(start, time);
        let arrivalHours = arrival.getHours();
        let arrivalMinutes = arrival.getMinutes();
        if (arrivalMinutes < 10) {
            arrivalMinutes = "0" + arrivalMinutes
        }

        return startHours + ":" + startMinutes + "-" + arrivalHours + ":" + arrivalMinutes;
    } else
        return "Invalid Format";
}

function addMinutes(date, minutes) {
    return new Date(date.getTime() + minutes * 60000);
}