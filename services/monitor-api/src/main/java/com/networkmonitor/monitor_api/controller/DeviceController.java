package com.networkmonitor.monitor_api.controller;

import com.networkmonitor.monitor_api.dto.DeviceRequest;
import com.networkmonitor.monitor_api.dto.DeviceResponse;
import com.networkmonitor.monitor_api.service.DeviceService;

import java.util.List;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestBody;

@RestController
@RequestMapping("/monitor-api/devices")
public class DeviceController {
    
    private final DeviceService service;

    public DeviceController(DeviceService service) {
        this.service = service;
    }

    @PostMapping
    public ResponseEntity<Void> receiveDevices( @RequestBody List<DeviceRequest> requests) {

        System.out.printf("Received %d devices%n", requests.size());

        service.saveAll(requests);

        return ResponseEntity.ok().build();
    }

    @GetMapping
    public List<DeviceResponse> getAll() {
        return service.getAll();
    }

}
