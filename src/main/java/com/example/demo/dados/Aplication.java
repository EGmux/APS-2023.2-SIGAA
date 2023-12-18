package com.example.demo.dados;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import com.example.demo.negocio.Credencial;


@Component
public class Aplication {
    

    @Autowired
    private ServiceCredencialMONGO serviceCredencialMONGO;

    @Autowired
    private RepositorioCredencialSpring repositorioCredencialSpring;


    public void inserir(Credencial credencial, String banco_dados){

        if (banco_dados.equals("SQL")) {
            System.out.println(banco_dados);
            repositorioCredencialSpring.inserir(credencial);
        }else if (banco_dados.equals("NOSQL")) {
            System.out.println(banco_dados);
            serviceCredencialMONGO.inserir(credencial);
        }

    }


    public Iterable<Credencial> getAll(String banco_dados){

        Iterable<Credencial> variable = null;

        if (banco_dados.equals("SQL")) {
            System.out.println(banco_dados);
            variable = repositorioCredencialSpring.getAll();

        }else if(banco_dados.equals("NOSQL")){
            System.out.println(banco_dados);
            variable = serviceCredencialMONGO.getAll();
        }
        
        return variable;

    }




}
