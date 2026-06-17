package com.networkmonitor.monitor_api.dto;

public record DeviceRequest(
    String ip,
    String hostname,
    boolean online,
    long responseTimeMs,
    String scannedAt
) {}