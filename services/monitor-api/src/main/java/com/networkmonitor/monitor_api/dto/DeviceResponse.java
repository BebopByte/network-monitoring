package com.networkmonitor.monitor_api.dto;

import java.time.Instant;

public record DeviceResponse(
    Long id,
    String ip,
    String hostname,
    boolean online,
    long responseTimeMs,
    Instant scannedAt
) {
}
