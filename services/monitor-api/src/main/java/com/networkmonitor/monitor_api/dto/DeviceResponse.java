package com.networkmonitor.monitor_api.dto;

public record DeviceResponse(
    Long id,
    String ip,
    String hostname,
    boolean online,
    long responseTimeMs
) {
}
