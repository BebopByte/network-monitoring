package com.networkmonitor.monitor_api.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import com.networkmonitor.monitor_api.domain.DeviceEntity;

public interface DeviceRepository extends JpaRepository<DeviceEntity, Long>{
    
}
