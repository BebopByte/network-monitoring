const API_URL = "http://localhost:8080/monitor-api/devices"

export async function getLatestDevices() {
    const response = await fetch(API_URL + "/latest")

    if (!response.ok) {
        throw new Error("Failed to fetch devices");
    }

    return response.json();
}