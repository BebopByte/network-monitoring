package com.networkmonitor.monitor_api.service;

import java.time.Instant;
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
        device.setScannedAt(Instant.parse(request.scannedAt()));

        repository.save(device);
    }

    public void saveAll(List<DeviceRequest> requests) {

        List<DeviceEntity> devices = requests.stream()
            .map(device -> toEntity(device))
            .toList();

        repository.saveAll(devices);
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

    private DeviceEntity toEntity(DeviceRequest request) {
        
        DeviceEntity device = new DeviceEntity();

        device.setIp(request.ip());
        device.setHostname(request.hostname());
        device.setOnline(request.online());
        device.setResponseTimeMs(request.responseTimeMs());
        device.setScannedAt(Instant.parse(request.scannedAt())); 

        return device;
    }

    public List<DeviceResponse> getLatestDevices() {

        return repository.findLatestDevices()
                .stream()
                .map(device -> toResponse(device))
                .toList();
    }

}
