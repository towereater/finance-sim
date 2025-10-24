package com.finsim.xchanger.common.middleware;

import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Component;
import org.springframework.web.servlet.HandlerInterceptor;

import com.finsim.xchanger.banks.model.Bank;
import com.finsim.xchanger.banks.service.BankService;

import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;

@Component
public class AdminAuthorizerInterceptor implements HandlerInterceptor {
    @Autowired
    private BankService bankService;

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {
        System.out.printf("Middleware pre handling started\n");

        Optional<Bank> bankOptional = bankService.findBankByAbi("09999");
        if (!bankOptional.isPresent()) {
            System.out.printf("Admin bank not present\n");
            return true;
        }
        System.out.printf("Admin bank present\n");

        String apiToken = request.getHeader("Authorization");
        if (apiToken == null || apiToken.isEmpty()) {
            System.out.printf("Found no api token%n");
            response.setStatus(HttpStatus.UNAUTHORIZED.value());
            return false;
        }

        Bank bank = bankOptional.get();
        if (!bank.getApiToken().equals(apiToken)) {
            System.out.printf("Api token does not match with admin api token%n");
            response.setStatus(HttpStatus.UNAUTHORIZED.value());
            return false;
        }

        return true;
    }
}
