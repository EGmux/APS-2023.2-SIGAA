package com.example.demo.dados;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class FabricaRepositoriosNOSQL implements FabricaAbstrata {
    
    @Autowired
    private ServiceCredencialMONGO serviceCredencialMONGO;

    @Override
    public ServiceCredencialMONGO CreateRepository(){

        System.out.println("Creating MongoDB");
        return serviceCredencialMONGO;
    }

    
}
