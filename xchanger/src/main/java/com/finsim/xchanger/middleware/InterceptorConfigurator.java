package com.finsim.xchanger.middleware;

import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.config.annotation.InterceptorRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

@Configuration
public class InterceptorConfigurator implements WebMvcConfigurer {
    @Override
    public void addInterceptors(InterceptorRegistry registry) {
        // Logger
        registry.addInterceptor(new LoggerInterceptor()).addPathPatterns("/*");

        // API token authorizer
        registry.addInterceptor(new AuthorizerInterceptor()).addPathPatterns("/dossiers/*");
    }
}
