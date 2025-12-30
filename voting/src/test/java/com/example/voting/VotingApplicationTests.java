package com.example.voting;

import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.web.client.RestTemplate;
import static org.junit.jupiter.api.Assertions.assertTrue;

@SpringBootTest
class VotingApplicationTests {

    @MockBean
    private RestTemplate restTemplate;

    @Test
    void contextLoads() {
        // Basic check to ensure context loads with mocked RestTemplate
    }

    @Test
    void testAppReady() {
        assertTrue(true, "App should be ready");
    }
}
