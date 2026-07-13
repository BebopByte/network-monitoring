
function DeviceRow({ device }) {
    return(
        <tr>
            <td>{device.ip}</td>
            <td>{device.hostname}</td>
            <td>{device.online ? "Online" : "Offline"}</td>
            <td>{device.responseTime} ms</td>
            <td>{device.scannedAt}</td>
        </tr>
    )
}

export default DeviceRow;