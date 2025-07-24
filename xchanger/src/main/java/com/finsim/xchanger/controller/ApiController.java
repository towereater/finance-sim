package com.finsim.xchanger.controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/hello")
public class ApiController {
    @GetMapping
    public String sayHello() {
        return "Hello, World!";
    }
}
