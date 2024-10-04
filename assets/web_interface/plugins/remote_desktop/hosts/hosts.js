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
            for (let i = 0; i < 20; i = i + 1) {
                let device = devices[deviceIndex]

                let deviceElement = document.createElement("div")
                deviceElement.className = "vertical-layout device"

                let deviceNameElement = document.createElement("span")
                deviceNameElement.innerText = device.hostname

                let deviceAddressElement = document.createElement("span")
                deviceAddressElement.innerText = "Address: " + device.url

                let devicePlatformElement = document.createElement("span")
                devicePlatformElement.innerText = "Platform: " + device.platform

                deviceElement.appendChild(deviceNameElement)
                deviceElement.appendChild(Object.assign(document.createElement("div"), { className: "splitter" }))
                deviceElement.appendChild(deviceAddressElement)
                deviceElement.appendChild(devicePlatformElement)

                let streamUrl = "http://" + device.url + "/remote_desktop/stream"
                deviceElement.onclick = () => window.open(streamUrl)

                devicesList.appendChild(deviceElement)
            }
        }
    }
    async_request("GET", url + "/devices", null, callback)
}