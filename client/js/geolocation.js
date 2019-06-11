// Not working with Safari
navigator.geolocation.getCurrentPosition(function (location) {
    console.log(location.coords.latitude);
    console.log(location.coords.longitude);
    console.log(location.coords.accuracy);
});