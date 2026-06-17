package com.networkmonitor.monitor_api.dto;

import java.time.LocalDateTime;

public record DeviceResponse(
    Long id,
    String ip,
    String hostname,
    boolean online,
    long responseTimeMs,
    LocalDateTime scannedAt
) {
}
