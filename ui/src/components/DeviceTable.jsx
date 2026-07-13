import DeviceRow from "./DeviceRow";

function DeviceTable({ devices }) {
    return (
        <table>
            <thead>
                <th>IP</th>
                <th>Hostname</th>
                <th>Status</th>
                <th>Response Time</th>
                <th>Last Seen</th>
            </thead>

            <tbody>
                {devices.map(device => (
                    <DeviceRow
                        key={device.ip}
                        device={device}
                    />
                ))}
            </tbody>      
        </table>
    )
}

export default DeviceTable;