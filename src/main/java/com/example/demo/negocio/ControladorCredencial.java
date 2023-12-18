package com.example.demo.negocio;


import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;



@Component
public class ControladorCredencial {

    @Autowired
    private  ControladorCadastroCredencial cadastroCredencial;
    @Autowired
    private AdaptardorCredencial adaptadorCredencial;
    


    private Iterable<Credencial> credenciais;

 
    public boolean verificarCredencial(Credencial credencial, String banco_dados){

        credenciais = cadastroCredencial.getAll(banco_dados);


        for (Credencial elm : credenciais) {
            

            if (elm.getUser().equals(credencial.getUser())
            && elm.getSenha().equals(credencial.getSenha())) {
                
                return true;
            }
        }
        return false;
        
    }

    public void inserir(Credencial credencial, String banco_dados){

        cadastroCredencial.inserir(credencial, banco_dados);
    }

    public Iterable<Credencial> getAllCredentials(String banco_dados){
        return cadastroCredencial.getAll(banco_dados);
    }

    public Credencial adaptarCredencial (Email username, String password){
        return adaptadorCredencial.adaptarCredencial(username, password);
    }

    //public Credencial adapterEmailToCredencial (String email){
    //    return 
    //}
    
}
