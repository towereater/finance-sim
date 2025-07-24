package com.finsim.xchanger.middleware;

import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Component;
import org.springframework.web.servlet.HandlerInterceptor;

import com.finsim.xchanger.model.Bank;
import com.finsim.xchanger.service.BankService;

import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;

@Component
public class AuthorizerInterceptor implements HandlerInterceptor {
    @Autowired
    private BankService bankService;

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {
        String apiToken = request.getHeader("Authorization");

        if (apiToken == null || apiToken.isEmpty()) {
            System.out.printf("Found no api token%n");
            response.setStatus(HttpStatus.UNAUTHORIZED.value());
            return false;
        }

        Optional<Bank> bank = bankService.findBankByApiToken(apiToken);

        if (bank.isPresent()) {
            return true;
        } else {
            response.setStatus(HttpStatus.UNAUTHORIZED.value());
            System.out.printf("Found no bank with api token %s%n", apiToken);
            return false;
        }
    }
}
