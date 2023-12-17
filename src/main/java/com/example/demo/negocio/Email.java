package com.example.demo.negocio;

import javax.persistence.Entity;
import javax.persistence.Id;

@Entity
public class Email {
    
    @Id
    private String email;

    public Email(String email){
        this.email = email;
    }

    public String getEmail(){
        return email;
    }

    public Character charAt(int pos){
        return email.charAt(pos);
    }

    public int length(){
        return email.length();
    }
}
