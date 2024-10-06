const proto = "http://"

function request(methood, url, data) {
    url = proto + url

    var req = new XMLHttpRequest();
    req.open(methood, url, false);
    req.send(data);
    return req.responseText
}

function Init() {
    let canvas = document.getElementById("canvas")
    canvas.__custom__ = {}

    setTimeout(window.LaunchStream, 0, window.location.host)
}

function LaunchStream(url) {
    let canvas = document.getElementById("canvas")
    let ctx = canvas.getContext('2d')
    newImage = new Image();
    newImage.onload = function () {
        canvas.width = newImage.width
        canvas.height = newImage.height
        // let scale = window.screen.width / newImage.width
        // canvas.__custom__.scale = scale
        // ctx.scale(scale, scale)
        ctx.drawImage(newImage, 0, 0, newImage.width, newImage.height);
        setTimeout(window.LaunchStream, 16, url)
    }
    newImage.src = proto + url + "/remote_desktop/image" + "?time=" + new Date().getTime();
}