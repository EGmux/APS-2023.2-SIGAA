package com.example.demo.negocio;

import org.springframework.stereotype.Component;

@Component
public class AdaptardorCredencial {
    

    private String usr = "";

    public Credencial adaptarCredencial(Email email, String password){

        for (int i = 0; i < email.length(); i++) {
            
            char c = email.charAt(i);

            if (c != '@') {
                usr += c;
            }else{
                break;
            }

        }

        Credencial credencial = new Credencial(usr, password);

        return credencial;
    }




}
