package com.finsim.xchanger.middleware;

import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.web.servlet.HandlerInterceptor;

import com.finsim.xchanger.entity.Bank;
import com.finsim.xchanger.service.BankService;

import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;

@Component
public class AuthorizerInterceptor implements HandlerInterceptor {
    @Autowired
    private BankService bankService;

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {
        String apiToken = request.getHeader("authorization");

        Optional<Bank> bank = bankService.findBankByApiToken(apiToken);

        if (!bank.isPresent()) {
            return false;
        }

        return true;
    }
}
