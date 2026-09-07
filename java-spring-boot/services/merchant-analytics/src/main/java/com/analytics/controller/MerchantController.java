package com.analytics.controller;

import com.analytics.dto.DailyTotalsResponse;
import com.analytics.service.AnalyticsService;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.text.SimpleDateFormat;
import java.util.Calendar;
import java.util.Date;
import java.util.TimeZone;

@RestController
public class MerchantController {
    private final AnalyticsService service;

    public MerchantController(AnalyticsService service) {
        this.service = service;
    }

    @GetMapping("/merchants/{id}/daily-totals")
    public ResponseEntity<DailyTotalsResponse> daily(
            @PathVariable String id,
            @RequestParam String date
    ) throws Exception {
        SimpleDateFormat formatter = new SimpleDateFormat("yyyy-MM-dd");
        formatter.setTimeZone(TimeZone.getTimeZone("UTC"));

        Date start = formatter.parse(date);
        Calendar calendar = Calendar.getInstance(TimeZone.getTimeZone("UTC"));
        calendar.setTime(start);
        calendar.add(Calendar.DATE, 1);

        return ResponseEntity.ok(service.daily(id, start, calendar.getTime()));
    }
}
