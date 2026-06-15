package com.networkmonitor.monitor_api.controller;

import com.networkmonitor.monitor_api.dto.DeviceRequest;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestBody;

@RestController
@RequestMapping("/monitor-api/devices")
public class DeviceController {
    
    @PostMapping
    public ResponseEntity<Void> receiveDevice( @RequestBody DeviceRequest request ) {

        System.out.println("Recieved device:");
        System.out.println(request);

        return ResponseEntity.ok().build();

    }

}
