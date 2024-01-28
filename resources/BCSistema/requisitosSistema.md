
glossário:
    atores:= personas que interagem com o sistema de alguma maneira

    descricao:=o que aquela funcionalidade deve permitir em linguagem natural

    continuacao:=personas responsáveis por progredir com o fluxo da funcionalidade

    output:saída do sistema

    limitacoes:o que é válido para qualquer fluxo daquela funcionalidade sempre
input não foi mencionado visto que é parte da UI e UX, pode ser abstraído no momento

-Sistema/genérico
Trocar senha se esquecer da atual

    atores: aluno
    descricao: fornece pagina html
    continuacao:
    output: confirmacao ou rejeicao da nova senha
    limitacoes: nenhuma

autenticação via provider

    atores: aluno e provider
    descricao: fornece pagina html
    continuacao: provider
    output: login do aluno
    limitacoes: sem suponte para autenticação de 2 fatores

colocar uma foto de perfil

    atores: aluno
    descricao: fornece página html
    continuacao:
    output: atualizacao da foto de perfil
    limitacoes: fotos que estejam salvas localmente no dispositivo apenas

alterar dados cadastrais/pessoais/bancarios

    atores: aluno
    descricao: fornece pagina html
    continuacao:
    output: atualizar dados cadastrais/pessoais/bancarios
    limitacoes: nenhuma

logar e deslogar do sistema

    atores: aluno
    descricao: loga/desloga aluno
    continuacao:
    output: loga/desloga aluno
    limitacoes: não há suporte para sessoes

tema preto/branco da interface

    atores: aluno
    descricao: alterna entre tema preto ou branco
    continuaca;o:
    output: troca tema da interface
    limitacoes: apenas preto e branco, sem suporte para grayscale ou outras modificacoes

pesquisar por alguma feature

    atores: aluno
    descricao: fornece barrad de pesquisa
    continuacao:
    output: acessa pagina da feature
    limitacoes: usadas apenas na interfaces que não envolvem formulário

acessar notificacoes de outros subsistemas

    atores: aluno
    descricao: fornece notificacoes
    continuacao:
    output: notificacao é exibida
    limitacaoes: nenhuma

prover feedback de alguma feature ou bug

    atores: aluno e Equipe de desenvolvimento
    descricao: fornece formulário em cada página html da interface, exceto nos formulários
    continuacao: Equipe de desenvolvimento
    output: equiped de dev é notificada e abre status de pedido
    limitacoes: podem ter vários feedback para cada página html por aluno
