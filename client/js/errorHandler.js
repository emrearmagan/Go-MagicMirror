
//if our server returns an error, we catch it here.
function error(container, error) {
    console.log(error)
    let errorEvent = document.createElement("div");
    errorEvent.setAttribute('class', 'errorHandler');
    let text = document.createTextNode(error);
    errorEvent.appendChild(text);
    container.appendChild(errorEvent);
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}