package com.networkmonitor.monitor_api.service;

import java.time.OffsetDateTime;
import java.util.List;

import org.springframework.stereotype.Service;

import com.networkmonitor.monitor_api.domain.DeviceEntity;
import com.networkmonitor.monitor_api.dto.DeviceRequest;
import com.networkmonitor.monitor_api.dto.DeviceResponse;
import com.networkmonitor.monitor_api.repository.DeviceRepository;

@Service
public class DeviceService {
    
    private final DeviceRepository repository;

    public DeviceService(DeviceRepository repository) {
        this.repository = repository;
    }

    public void save(DeviceRequest request) {

        DeviceEntity device = new DeviceEntity();

        device.setIp(request.ip());
        device.setHostname(request.hostname());
        device.setOnline(request.online());
        device.setResponseTimeMs(request.responseTimeMs());
        device.setScannedAt(OffsetDateTime.parse(request.scannedAt()).toLocalDateTime());

        repository.save(device);
    }

    public List<DeviceResponse> getAll(){
        return repository.findAll()
        .stream()
        .map(device -> toResponse(device))
        .toList();
    }

    private DeviceResponse toResponse(DeviceEntity device) {
        return new DeviceResponse(
            device.getId(),
            device.getIp(),
            device.getHostname(),
            device.getOnline(),
            device.getResponeTimeMs(),
            device.getScannedAt()
        );
    }

}
