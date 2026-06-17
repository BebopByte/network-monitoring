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
    public ResponseEntity<Void> receiveDevice( @RequestBody DeviceRequest request ) {

        System.out.printf("Received Device: %s (%s)%n", request.hostname(), request.ip());

        service.save(request);

        return ResponseEntity.ok().build();

    }

    @GetMapping
    public List<DeviceResponse> getAll() {
        return service.getAll();
    }

}
