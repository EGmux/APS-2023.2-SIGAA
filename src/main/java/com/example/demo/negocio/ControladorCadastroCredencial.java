package com.example.demo.negocio;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import com.example.demo.dados.Aplication;



@Component
public class ControladorCadastroCredencial {
    
    

    @Autowired
    private Aplication aplication;


    public Iterable<Credencial> getAll(String banco_dados){

        return aplication.getAll(banco_dados);
    }

    public void inserir(Credencial credencial, String banco_dados){
        aplication.inserir(credencial, banco_dados);
    }

}
