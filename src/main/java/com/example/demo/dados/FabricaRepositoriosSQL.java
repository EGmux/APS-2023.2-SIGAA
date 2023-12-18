package com.example.demo.dados;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;



@Component
public class FabricaRepositoriosSQL implements FabricaAbstrata{
    

    @Autowired
    private RepositorioCredencialSpring repositorioCredencialSpring;

    @Override
    public RepositorioCredencialSpring CreateRepository(){ 
        
        System.out.println("Creating SQLDB");

        repositorioCredencialSpring = new RepositorioCredencialSpring();

        System.out.println(repositorioCredencialSpring);

        return repositorioCredencialSpring;
    }


}
