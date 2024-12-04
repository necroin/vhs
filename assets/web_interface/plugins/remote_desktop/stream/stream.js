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

function isRotated() {
    let screenWidthConst = document.documentElement.clientWidth
    let screenHeightConst = document.documentElement.clientHeight
    return screenWidthConst < screenHeightConst && (screen.orientation.type == "landscape-primary" || screen.orientation.type == "landscape-secondary")
}

function GetScreenWidth() {
    let screenWidthConst = document.documentElement.clientWidth
    let screenHeightConst = document.documentElement.clientHeight

    if (isRotated()) {
        return screenHeightConst
    }
    return screenWidthConst
}

function GetScreenHeight() {
    let screenWidthConst = document.documentElement.clientWidth
    let screenHeightConst = document.documentElement.clientHeight

    if (isRotated()) {
        return screenWidthConst
    }
    return screenHeightConst
}

function scaleAxis(value, scale) {
    return Math.floor(value / scale)
}

function GetMouseCoords(event, canvas) {
    return {
        x: scaleAxis(event.offsetX, canvas.__custom__.scale),
        y: scaleAxis(event.offsetY, canvas.__custom__.scale)
    }
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
                coords: GetMouseCoords(event, canvas)
            })
        )
    }

    canvas.onmouseup = (event) => {
        request(
            "POST",
            mouseInputUrl,
            JSON.stringify({
                type: "leftUp",
                coords: GetMouseCoords(event, canvas)
            })
        )
    }

    canvas.onmousemove = (event) => {
        request(
            "POST",
            mouseInputUrl,
            JSON.stringify({
                type: "move",
                coords: GetMouseCoords(event, canvas)
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

        if (event.touches.length == 1) {
            let offsetX = scaleAxis(event.touches[0].pageX, canvas.__custom__.scale)
            let offsetY = scaleAxis(event.touches[0].pageY, canvas.__custom__.scale)
            canvas.__custom__.touchContext = { x: offsetX, y: offsetY }
        }
    }

    canvas.ontouchmove = (event) => {
        if (event.touches.length == 1 && canvas.__custom__.touchContext != null) {

            let offsetX = scaleAxis(event.touches[0].pageX, canvas.__custom__.scale)
            let offsetY = scaleAxis(event.touches[0].pageY, canvas.__custom__.scale)

            let deltaX = Math.floor((canvas.__custom__.touchContext.x - offsetX) / 1.5)
            let deltaY = Math.floor((canvas.__custom__.touchContext.y - offsetY) / 1.5)

            request(
                "POST",
                mouseInputUrl,
                JSON.stringify({
                    type: "scroll",
                    coords: {
                        x: canvas.__custom__.touchContext.x,
                        y: canvas.__custom__.touchContext.y
                    },
                    scroll_delta: {
                        x: deltaX,
                        y: deltaY
                    }
                })
            )

            event.preventDefault()
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
        // let screenWidth = GetScreenWidth()
        // let screenHeight = GetScreenHeight()
        // // let scale = Math.min(screenWidth / newImage.width, screenHeight / newImage.height);
        let scale = 1

        document.querySelector('meta[name="viewport"]').setAttribute('content', `width=${newImage.width}`)

        canvas.__custom__.scale = scale

        canvas.width = newImage.width * scale
        canvas.height = newImage.height * scale

        // canvas.imageSmoothingEnabled = false
        // ctx.imageSmoothingEnabled = false

        ctx.drawImage(newImage, 0, 0, newImage.width * scale, newImage.height * scale);
        setTimeout(window.LaunchStream, 16, url)
    }
    newImage.src = proto + url + "/remote_desktop/image" + "?time=" + new Date().getTime();
}
