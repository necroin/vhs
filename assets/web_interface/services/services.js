const proto = "http://"

function async_request(method, url, data, callback) {
    url = proto + url
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
    GetServices(window.location.host)
}

function GetDevices(url) {
    callback = (response) => {
        let devices = JSON.parse(response)

        let devicesList = document.getElementById("hosts")
        devicesList.replaceChildren()

        for (let deviceIndex in devices) {
            let device = devices[deviceIndex]

            let deviceElement = document.createElement("div")
            deviceElement.className = "list-item device vertical-layout"

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

            deviceElement.onclick = () => window.open(proto + device.url)

            devicesList.appendChild(deviceElement)
        }
    }
    async_request("GET", url + "/devices", null, callback)
}

function GetServices(url) {
    callback = (response) => {
        let services = JSON.parse(response)

        let servicesList = document.getElementById("services")
        servicesList.replaceChildren()

        for (let serviceName in services) {
            let serviceUrl = services[serviceName]

            let serviceElement = document.createElement("div")
            serviceElement.className = "list-item"

            serviceElement.innerText = serviceName
            serviceElement.onclick = () => window.open(serviceUrl)

            servicesList.appendChild(serviceElement)
        }
    }
    async_request("GET", url + "/services", null, callback)
}

