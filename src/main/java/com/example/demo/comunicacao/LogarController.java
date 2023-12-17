package com.example.demo.comunicacao;

import com.example.demo.negocio.Email;
import com.example.demo.negocio.Credencial;
import com.example.demo.negocio.Fachada;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;



@Controller
public class LogarController {

    //private static final AtomicLong idCounter = new AtomicLong();

    @Autowired
    private Fachada fachada;


    @GetMapping("/logar")
    public String realizarLogin(Model model){
        //model.addAttribute(null, model);
        return "logar";
    }

    @GetMapping("/logar/email")
    public String realizarLoginEmail(){


        return "logar_email";
    }

    @GetMapping("/logar/autenticar")
    public String autenticarUsuario(@RequestParam(name="username") String username,
     @RequestParam(name="password") String password,
     @RequestParam(name="formato") String formato){
        
        Credencial credencial = null;

        if (formato.equals("user")) {
            credencial = new Credencial(username,password);
        }else if (formato.equals("email")) {
            
            Email email = new Email(username);
            credencial = new Credencial();
            credencial = fachada.adaptarCredencial(email, password);


        }

        Boolean existe = fachada.autenticarCredencial(credencial);

        System.out.println(formato);


        if (existe) {
            return "redirect:/mainMenu";
        }else{
            return "redirect:/logar";
        }
    }



    //@GetMapping("/logar/oauth2/code/google")


    /*@GetMapping("/cliente/inserir/")
    public String novoCliente(@RequestParam(name = "nome") String nome){
        Conta conta = new Conta(idCounter.get(), String.valueOf(idCounter.get()));
        fachada.inserirConta(conta);
        Cliente cliente = new Cliente(idCounter.getAndIncrement(), nome, conta);
        fachada.inserirCliente(cliente);
        return "redirect:/cliente";
    }*/

}