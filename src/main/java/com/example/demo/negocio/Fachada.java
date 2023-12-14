package com.example.demo.negocio;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class Fachada {

    @Autowired
    private ControladorCliente controladorCliente;
    @Autowired
    private ControladorConta controladorConta;

    public void inserirCliente(Cliente cliente) {
        controladorCliente.inserir(cliente);
    }

    public Iterable<Cliente> getAllClientes() {
        return controladorCliente.getAll();
    }

    public void inserirConta(Conta conta) {
        controladorConta.inserir(conta);
    }
    
}
