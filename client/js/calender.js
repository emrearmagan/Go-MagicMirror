function Calender() {
    let req;
    req = new XMLHttpRequest();
    req.onreadystatechange = async function () {
        if (req.readyState == 4 && req.status == 200) {
            //Get the Node and clear it
            let calender = document.querySelector(".calender");
            let child = calender.lastElementChild;
            while (child) {
                calender.removeChild(child);
                child = calender.lastElementChild;
            }
            //Convert json in an object and fill the div in a list
            let obj = (function (raw) {
                try {
                    return JSON.parse(raw);
                } catch (err) {
                    return false;
                }
            })(req.responseText);

            console.log(req.responseText)
            if (!obj) {
                error(calender, req.responseText);
                return false;
            }

            if (obj.items.length < 1) {
                let noEvents = document.createElement("div");
                noEvents.setAttribute('class', 'no events');
                let text = document.createTextNode("No Events");
                noEvents.appendChild(text);
                calender.appendChild(noEvents);
                return false;
            }

            //Listing all elements
            for (let i = 0; i < obj.items.length; i++) {
                //Thats where each event and date is gonna be stored
                let eventDiv = document.createElement("ul");
                eventDiv.setAttribute("class", "event");

                //Create the elements
                let img = document.createElement("img");
                let event = document.createElement("LI");
                let date = document.createElement("LI");

                img.setAttribute('class', 'c icon');
                event.setAttribute("class", "c event");
                date.setAttribute("class", "c date");

                let textEvent = document.createTextNode(obj.items[i].summary);
                let textDate = document.createTextNode(obj.items[i].start.date);
                img.src = "/client/img/icons/calender_icon.png";

                event.appendChild(textEvent);
                date.appendChild(textDate);

                eventDiv.appendChild(img);
                eventDiv.appendChild(event);
                eventDiv.appendChild(date);

                calender.appendChild(eventDiv);
                await sleep(100);


            }

        }
    };
    req.open("GET", "http://localhost:8080/calender");
    req.send();
}