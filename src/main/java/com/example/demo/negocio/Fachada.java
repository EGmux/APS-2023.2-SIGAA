package com.example.demo.negocio;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;


@Component
public class Fachada {

    private String banco_dados = "SQL";

    
    @Autowired
    private ControladorCredencial controladorCredencial;


    public boolean autenticarCredencial(Credencial credencial){
        return controladorCredencial.verificarCredencial(credencial, banco_dados);
    }

    public void inserirCredencial(Credencial credencial){
        controladorCredencial.inserir(credencial, banco_dados);
    }

    public Iterable<Credencial> getAllCredentials(){
        return controladorCredencial.getAllCredentials(banco_dados);
    }

    public Credencial adaptarCredencial(Email username, String password){
        return controladorCredencial.adaptarCredencial(username, password);
    }

}
