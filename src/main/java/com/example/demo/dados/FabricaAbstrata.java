package com.example.demo.dados;

import org.springframework.stereotype.Component;

@Component
public interface FabricaAbstrata {
    
    IRepositorioCredencial CreateRepository();
}
