function async_request(method, url, data, callback) {
    url = "http://" + url
    console.log(url)

    var req = new XMLHttpRequest();
    req.onload = () => {
        if (req.readyState === XMLHttpRequest.DONE) {
            callback(req.responseText)
        }
    }
    req.open(method, url, true);
    req.send(data);
}

function Init() {
    GetDevices(window.location.host)
}

function GetDevices(url) {
    callback = (response) => {
        let devices = JSON.parse(response)

        let devicesList = document.getElementById("devices")
        devicesList.replaceChildren()

        for (let deviceIndex in devices) {
            let device = devices[deviceIndex]
            console.log(device)
            let streamUrl = "http://" + device.url + "/remote_desktop/stream"

            let deviceElement = document.createElement("div")
            deviceElement.className = "device"
            deviceElement.innerText = device.hostname
            deviceElement.onclick = () => window.open(streamUrl)
            devicesList.appendChild(deviceElement)
        }
    }
    async_request("GET", url + "/devices", null, callback)
}