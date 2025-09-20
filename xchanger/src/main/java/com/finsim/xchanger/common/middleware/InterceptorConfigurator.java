package com.finsim.xchanger.common.middleware;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.config.annotation.InterceptorRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

@Configuration
public class InterceptorConfigurator implements WebMvcConfigurer {
    @Autowired
    private LoggerInterceptor loggerInterceptor;

    @Autowired
    private AdminAuthorizerInterceptor adminAuthorizerInterceptor;

    @Autowired
    private AuthorizerInterceptor authorizerInterceptor;

    @Override
    public void addInterceptors(InterceptorRegistry registry) {
        // Logger
        registry.addInterceptor(loggerInterceptor).addPathPatterns("/**");

        // API token authorizer
        registry.addInterceptor(adminAuthorizerInterceptor).addPathPatterns(
            "/banks", "/banks/**");

        registry.addInterceptor(authorizerInterceptor).addPathPatterns(
            "/dossiers", "/dossiers/**",
                        "/stocks", "/stocks/**",
                        "/orders", "/orders/**");
    }
}
