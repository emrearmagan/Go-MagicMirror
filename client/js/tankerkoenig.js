function FuelPrice() {
    let req;
    req = new XMLHttpRequest();
    req.onreadystatechange = async function () {
        if (req.readyState == 4 && req.status == 200) {
            //Get the Node and clear it
            let fuelTable = document.querySelector(".fuel-table");
            let child = fuelTable.lastElementChild;
            while (child) {
                fuelTable.removeChild(child);
                child = fuelTable.lastElementChild;
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
                error(fuelTable, req.responseText);
                return false;
            }

            //Listing all elements
            for (let i = 0; i < obj.stations.length; i++) {
                //Thats where each event and date is gonna be stored
                let fuelList = document.createElement("ul");
                fuelList.setAttribute("class", "fuel list");

                //Create the elements
                let firm = document.createElement("LI");
                let dieselPrice = document.createElement("LI");
                let benzinPrice = document.createElement("LI");
                let location = document.createElement("LI");

                firm.setAttribute("class", "f_firm");
                dieselPrice.setAttribute("class", "f_diesel-price");
                benzinPrice.setAttribute("class", "f_benzin-price");
                location.setAttribute("class", "f_location");

                let textFirm = document.createTextNode(obj.stations[i].brand);
                let textDiesel = document.createTextNode(obj.stations[i].diesel);
                let textBenzin = document.createTextNode(obj.stations[i].e10);
                let textLocation = document.createTextNode(obj.stations[i].street + obj.stations[i].houseNumber);

                firm.appendChild(textFirm);
                dieselPrice.appendChild(textDiesel);
                benzinPrice.appendChild(textBenzin);
                location.appendChild(textLocation);


                fuelList.appendChild(firm);
                fuelList.appendChild(dieselPrice);
                fuelList.appendChild(benzinPrice);
                fuelList.appendChild(location);

                fuelTable.appendChild(fuelList);
                await sleep(200);
            }

        }
    };
    req.open("GET", "http://localhost:8080/fuel");
    req.send();
}