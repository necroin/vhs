const proto = "http://"
const servicePath = "/remote_desktop"
const mouseInputPath = "/input/mouse"
const keyboardInputPath = "/input/keyboard"

const mouseInputUrl = window.location.host + servicePath + mouseInputPath
const keyboardInputUrl = window.location.host + servicePath + keyboardInputPath

function request(methood, url, data) {
    url = proto + url

    var req = new XMLHttpRequest();
    req.open(methood, url, false);
    req.send(data);
    return req.responseText
}

function Init() {
    let canvas = document.getElementById("canvas")
    canvas.focus = () => { }

    canvas.__custom__ = {
        doubleClicked: false
    }

    canvas.onmousedown = (event) => {
        request(
            "POST",
            mouseInputUrl,
            JSON.stringify({
                type: "leftDown",
                coords: { x: event.offsetX, y: event.offsetY }
            })
        )
    }

    canvas.onmouseup = (event) => {
        request(
            "POST",
            mouseInputUrl,
            JSON.stringify({
                type: "leftUp",
                coords: { x: event.offsetX, y: event.offsetY }
            })
        )
    }

    canvas.onmousemove = (event) => {
        request(
            "POST",
            mouseInputUrl,
            JSON.stringify({
                type: "move",
                coords: { x: event.offsetX, y: event.offsetY }
            })
        )
    }

    canvas.onwheel = (event) => {
        request(
            "POST",
            mouseInputUrl,
            JSON.stringify({
                type: "scroll",
                coords: { x: event.offsetX, y: event.offsetY },
                scroll_delta: { x: Math.floor(-event.deltaX), y: Math.floor(-event.deltaY) }
            })
        )
    }

    canvas.onclick = (event) => { }
    canvas.ondblclick = (event) => { }

    canvas.ontouchstart = (event) => {
        canvas.__custom__.touchContext = null
        canvas.contentEditable = false

        if (event.touches.length == 1) {
            if (canvas.__custom__.doubleClicked) {
                canvas.contentEditable = true
            }
            canvas.__custom__.doubleClicked = true
            setTimeout(() => {
                canvas.__custom__.doubleClicked = false
            }, 500);
        }

        if (event.touches.length == 3) {
            let offsetX = event.touches[0].pageX
            let offsetY = event.touches[0].pageY
            canvas.__custom__.touchContext = { x: offsetX, y: offsetY }
        }
    }
    canvas.ontouchmove = (event) => {
        if (event.touches.length == 3 && canvas.__custom__.touchContext != null) {

            let offsetX = event.touches[0].pageX
            let offsetY = event.touches[0].pageY

            let deltaX = (canvas.__custom__.touchContext.x - offsetX) / 2
            let deltaY = (canvas.__custom__.touchContext.y - offsetY) / 2

            request(
                "POST",
                mouseInputUrl,
                JSON.stringify({
                    type: "scroll",
                    coords: { x: offsetX, y: offsetY },
                    scroll_delta: { x: deltaX, y: deltaY }
                })
            )
        }
    }
    canvas.ontouchend = (event) => { }

    canvas.onkeydown = (event) => {
        event.preventDefault()
        let keycode = event.keyCode
        request("POST", keyboardInputUrl, JSON.stringify({ type: "keydown", keycode: keycode }))
    }
    canvas.onkeyup = (event) => {
        event.preventDefault()
        let keycode = event.keyCode
        request("POST", keyboardInputUrl, JSON.stringify({ type: "keyup", keycode: keycode }))
    }

    setTimeout(window.LaunchStream, 0, window.location.host)
}

function LaunchStream(url) {
    let canvas = document.getElementById("canvas")
    let ctx = canvas.getContext('2d')
    newImage = new Image();
    newImage.onload = function () {
        canvas.width = newImage.width
        canvas.height = newImage.height

        canvas.style.width = newImage.width
        canvas.style.height = newImage.height
        document.body.style.width = newImage.width
        document.body.style.height = newImage.height
        // let scale = window.screen.width / newImage.width
        // canvas.__custom__.scale = scale
        // ctx.scale(scale, scale)
        ctx.drawImage(newImage, 0, 0, newImage.width, newImage.height);
        setTimeout(window.LaunchStream, 16, url)
    }
    newImage.src = proto + url + "/remote_desktop/image" + "?time=" + new Date().getTime();
}
