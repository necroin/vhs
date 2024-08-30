window.__context__ = {}

function remove_dropdown(event, event_id, dropdown_id) {
    if (!event.target.matches('#' + event_id)) {
        document.getElementById(dropdown_id).classList.remove('dropdown-show');
    }
}

window.onclick = function (event) {
    remove_dropdown(event, "tool-bar-button-create", "tool-bar-create-options")
}

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

function request(method, url, data) {
    url = "http://" + url

    var req = new XMLHttpRequest();
    req.open(method, url, false);
    req.send(data);
    return req.responseText
}

function Init() {
    let savedPath = window.localStorage.getItem("explorer-path")
    if (savedPath == null) {
        savedPath = "/"
    }

    GetDevices(window.location.host)
    GetFilesystem(window.location.host, savedPath)
}

function GetCurrentPath() {
    return document.getElementById("explorer-address-line").value
}

function SetCurrentPath(path) {
    document.getElementById("explorer-address-line").value = path
    window.localStorage.setItem("explorer-path", path)
}

function SetFocusItem(element, item) {
    if (element == item) {
        return
    }
    if (element.focusItem != null) {
        element.focusItem.style.backgroundColor = "var(--backgroud-color)"
    }
    element.focusItem = item
    element.focusItem.style.backgroundColor = "var(--focus-bg-color)"
}

function OpenDialog(dialog, overlay) {
    document.getElementById(dialog).style.display = "flex";
    document.getElementById(overlay).style.display = "block";
}

function CloseDialog(dialog, overlay) {
    document.getElementById(dialog).style.display = "none";
    document.getElementById(overlay).style.display = "none";
}

function UpdateStatusBar(data) {
    document.getElementById("status-bar-progress").innerHTML = data.status
    document.getElementById("status-bar-text").innerText = data.text
}

function GetDevices(url) {
    callback = (response) => {
        let devices = JSON.parse(response)

        let devicesList = document.getElementById("devices")
        devicesList.replaceChildren()

        let createOptions = document.getElementById("create-storage-select")
        createOptions.replaceChildren()

        let allDevicesElement = document.createElement("span")
        allDevicesElement.className = "device"
        allDevicesElement.innerText = "All"
        allDevicesElement.tabIndex = 0
        allDevicesElement.onclick = () => {
            window.__context__.storage = null
            GetFilesystem(window.location.host, GetCurrentPath())
        }
        devicesList.appendChild(allDevicesElement)

        let devicesCount = 1
        for (let deviceIndex in devices) {
            let device = devices[deviceIndex]

            let deviceElement = document.createElement("span")
            deviceElement.className = "device"
            deviceElement.innerText = device.hostname
            deviceElement.tabIndex = devicesCount
            deviceElement.onclick = () => {
                window.__context__.storage = device
                GetFilesystem(window.location.host, GetCurrentPath())
            }
            devicesList.appendChild(deviceElement)

            let createOptionDeviceElement = document.createElement("option")
            createOptionDeviceElement.innerText = device.hostname
            createOptionDeviceElement.__custom__ = {
                storage: device,
            }
            createOptions.appendChild(createOptionDeviceElement)

            devicesCount = devicesCount + 1
        }
    }
    async_request("GET", url + "/filesystem/devices", null, callback)
}

function GetFilesystem(url, path) {
    SetCurrentPath(path)

    callback = (response) => {
        let rowsCount = 0

        let filesystem = JSON.parse(response)

        let filesystemTable = document.getElementById("explorer-content-body")
        filesystemTable.replaceChildren()
        filesystemTable.focusItem = null

        let directories = filesystem.directories
        for (let directory in directories) {
            let info = directories[directory]

            let tableRow = document.createElement("tr")
            tableRow.tabIndex = String(rowsCount)
            tableRow.__custom__ = {
                "name": directory
            }

            let nameElement = document.createElement("td")
            let storageElement = document.createElement("td")
            let dateElement = document.createElement("td")
            let typeElement = document.createElement("td")
            let sizeElement = document.createElement("td")

            nameElement.innerText = "📁 " + directory
            dateElement.innerText = info["mod_time"]
            typeElement.innerText = "Directory"
            sizeElement.innerText = ""

            tableRow.appendChild(nameElement)
            tableRow.appendChild(storageElement)
            tableRow.appendChild(dateElement)
            tableRow.appendChild(typeElement)
            tableRow.appendChild(sizeElement)

            let openPath = [path, directory].join("/")
            if (path == "/") {
                openPath = path + directory
            }

            tableRow.ondblclick = () => GetFilesystem(url, openPath)
            tableRow.ontouchend = () => GetFilesystem(url, openPath)

            filesystemTable.appendChild(tableRow)

            rowsCount = rowsCount + 1
        }

        let files = filesystem.files
        for (let file in files) {
            let infos = files[file]
            for (infoIndex in infos) {
                let info = infos[infoIndex]

                let tableRow = document.createElement("tr")
                tableRow.tabIndex = String(rowsCount)
                tableRow.__custom__ = {
                    "name": file,
                    "url": info["url"],
                    "platform": info["platform"],
                    "hostname": info["hostname"],
                }

                let nameElement = document.createElement("td")
                let storageElement = document.createElement("td")
                let dateElement = document.createElement("td")
                let typeElement = document.createElement("td")
                let sizeElement = document.createElement("td")

                nameElement.innerText = file
                storageElement.innerText = info["hostname"]
                dateElement.innerText = info["mod_time"]
                typeElement.innerText = "File"
                sizeElement.innerText = info["size"] + " KB"

                tableRow.appendChild(nameElement)
                tableRow.appendChild(storageElement)
                tableRow.appendChild(dateElement)
                tableRow.appendChild(typeElement)
                tableRow.appendChild(sizeElement)

                filesystemTable.appendChild(tableRow)

                rowsCount = rowsCount + 1
            }
        }
    }

    let endpoint = "/filesystem/all"

    if (window.__context__.storage != null) {
        url = window.__context__.storage.url
        endpoint = "/filesystem/self"
    }

    async_request("POST", url + endpoint, path, callback)
}

function Back() {
    let pathItems = GetCurrentPath().split("/");
    let path = pathItems.slice(0, pathItems.length - 1).join("/");
    if (path == "") {
        path = "/"
    }
    GetFilesystem(window.location.host, path)
}

function Create(type) {
    let createStorageSelect = document.getElementById("create-storage-select")
    let url = createStorageSelect.options[createStorageSelect.selectedIndex].__custom__.storage.url

    let currenPath = GetCurrentPath()
    let name = document.getElementById("create-dialog-name").value

    let createPath = [currenPath, name].join("/")
    if (currenPath == "/") {
        createPath = currenPath + name
    }

    let data = JSON.stringify(
        {
            "type": type,
            "path": createPath,
        }
    )

    let response = request("POST", url + "/filesystem/create", data);
    console.log(response)
    CloseDialog('create-dialog', 'create-dialog-overlay')

    if (response == "") {
        UpdateStatusBar({
            status: "✓",
            text: "successfully created"
        })
    } else {
        UpdateStatusBar({
            status: "✕",
            text: response
        })
    }

}