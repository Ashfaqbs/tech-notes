package com.ashfaq.sdk;


import org.springframework.http.ResponseEntity;
import org.springframework.web.client.RestTemplate;

public class SdkClient {
    private final RestTemplate restTemplate;

    public SdkClient() {
        this.restTemplate = new RestTemplate();
    }

    public ApiResponse getApiResponse(String url) {
        ResponseEntity<ApiResponse> response = restTemplate.getForEntity(url, ApiResponse.class);
        return response.getBody();
    }
}
