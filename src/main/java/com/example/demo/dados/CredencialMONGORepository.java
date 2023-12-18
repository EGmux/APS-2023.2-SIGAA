package com.example.demo.dados;

import com.example.demo.negocio.Credencial;

import org.springframework.data.mongodb.repository.MongoRepository;


public interface CredencialMONGORepository extends  MongoRepository <Credencial, Long>{
    
    
}
