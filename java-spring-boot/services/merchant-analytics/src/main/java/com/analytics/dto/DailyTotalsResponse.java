package com.analytics.dto;

import java.util.Map;

public record DailyTotalsResponse(
        double total,
        int count,
        Map<String, Integer> byStatus
) {
}
