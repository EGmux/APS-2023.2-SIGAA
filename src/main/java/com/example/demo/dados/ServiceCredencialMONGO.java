package com.example.demo.dados;

import com.example.demo.negocio.Credencial;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;



@Service
public class ServiceCredencialMONGO implements IRepositorioCredencial{

    @Autowired
    CredencialMONGORepository credencialMONGO;


    public void inserir(Credencial credencial){

        this.credencialMONGO.save(credencial);
        
    }
    

    public Iterable<Credencial> getAll(){

        return this.credencialMONGO.findAll();
    }

}
