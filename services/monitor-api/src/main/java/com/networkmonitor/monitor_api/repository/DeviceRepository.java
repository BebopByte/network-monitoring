package com.networkmonitor.monitor_api.repository;

import java.util.List;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

import com.networkmonitor.monitor_api.domain.DeviceEntity;

@Repository
public interface DeviceRepository extends JpaRepository<DeviceEntity, Long>{
    
    @Query(value = """
            SELECT DISTINCT ON (ip) *
            FROM devices
            ORDER BY ip, scanned_at DESC
            """, nativeQuery = true)
    List<DeviceEntity> findLatestDevices();

}
