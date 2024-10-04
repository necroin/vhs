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
    GetServices(window.location.host)
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