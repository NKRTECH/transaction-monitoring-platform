package com.gtrs.transaction;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

/**
 * Transaction Service - Independent Spring Boot Microservice
 * 
 * This service handles all transaction-related operations including:
 * - Transaction CRUD operations
 * - Transaction validation coordination
 * - Business rule enforcement
 * - Audit logging
 */
@SpringBootApplication
public class TransactionServiceApplication {

    public static void main(String[] args) {
        SpringApplication.run(TransactionServiceApplication.class, args);
    }
}