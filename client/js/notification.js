// function test() {
//     var xmlhttp;
//     xmlhttp = new XMLHttpRequest();
//     xmlhttp.onreadystatechange = function () {
//         if (xmlhttp.readyState == 4 && xmlhttp.status == 200) {
//             document.getElementById("numberOne").innerHTML = "THIS: " + this.responseText
//         }
//     }
//     xmlhttp.open("GET", "http://localhost:8080/notification");
//     xmlhttp.send();
// }

function notifier() {
    let source = new EventSource("http://localhost:8080/notification");
    source.onmessage = function (event) {
        console.log("OnMessage called:");
        console.dir(event);

        //Get the Node and clear it
        let notfications = document.querySelector(".notifications-list");
        let child = notfications.lastElementChild;
        while (child) {
            notfications.removeChild(child);
            child = notfications.lastElementChild;
        }

        //@todo add sound on new notifcation
        //Convert json in an object and fill the div in a list
        let obj = (function (data) {
            try {
                return JSON.parse(data);
            } catch (err) {
                return false;
            }
        })(event.data);


        if (!obj) {
            error(notfications, event.responseText);
            return false;
        }

        //If no notifications show message
        if (obj.Message.length < 1) {
            let node = document.createElement("div");
            node.setAttribute('class', 'no-notifications');
            let text = document.createTextNode("No new notifications");
            node.appendChild(text);
            notfications.appendChild(node);
            return;
        }
        for (let i = 0; i < obj.Message.length; i++) {
            let node = document.createElement("LI");
            let text = document.createTextNode(obj.Message[i]);
            node.appendChild(text);
            notfications.appendChild(node);
        }

    }
}
