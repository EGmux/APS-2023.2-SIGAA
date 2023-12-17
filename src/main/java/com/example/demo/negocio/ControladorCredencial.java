package com.example.demo.negocio;


import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;


@Component
public class ControladorCredencial {

    @Autowired
    private  CadastroCredencial cadastroCredencial;
    @Autowired
    private AdaptardorCredencial adaptadorCredencial;
    

    private Iterable<Credencial> credenciais;

 
    public boolean verificarCredencial(Credencial credencial){

        credenciais = cadastroCredencial.getAll();


        for (Credencial elm : credenciais) {
            

            if (elm.getUser().equals(credencial.getUser())
            && elm.getSenha().equals(credencial.getSenha())) {
                
                return true;
            }
        }
        return false;
        
    }

    public void inserir(Credencial credencial){

        cadastroCredencial.inserir(credencial);
    }

    public Iterable<Credencial> getAllCredentials(){
        return cadastroCredencial.getAll();
    }

    public Credencial adaptarCredencial (Email username, String password){
        return adaptadorCredencial.adaptarCredencial(username, password);
    }

    //public Credencial adapterEmailToCredencial (String email){
    //    return 
    //}
    
}
