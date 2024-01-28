
glossário:
        atores:= personas que interagem com o sistema de alguma maneira

        descricao:=o que aquela funcionalidade deve permitir em linguagem natural

        continuacao:=personas responsáveis por progredir com o fluxo da funcionalidade

        output:saída do sistema

        limitacoes:o que é válido para qualquer fluxo daquela funcionalidade sempre

    input não foi mencionado visto que é parte da UI e UX, pode ser abstraído no momento

-Auxílio
Quais infomações mais importantes para esse BC?
Solicitar bolsa auxílio fora do período de solicitação/ pedido emergencial

    atores: aluno e PROAES
    descricao: formulário e notificar PROAES
    continuacao: aluno e PROAES
    output: proaes é notificada e status de pedido aberto
    limitacoes: o aluno deve ser de baixa renda comprovadamente

Filtrar bolsas de auxílio disponíveis por algum critério

    atores: aluno
    descricao: apresenta ementa de bolsas para consulta e baixar
    continuacao: aluno
    output:disponibilza pagin html
    limitacoes: apenas de bolsas vigentes

Aplicar para bolsas de auxilio

    atores: aluno e PROAES
    descricao: formulário e notifica PROAES
    continuacao: aluno e PROAES
    output: notifica PROAES e abre status de pedido
    limitacoes: apenas de bolsas vigentes e aluno é de baixa renda

Acompanhar status do pedido da aplicação

    atores: aluno
    descricao: mostra status do pedido
    continuacao: PROAES
    output: notifica o aluno a cada atualização da proaes
    limitacoes: o aluno deve sastifazer os critérios prévios e a proaes precisa atualizar o status

Gerenciar atuais bolsa de auxílio

    atores: aluno
    descricao: consulta de bolsas de auxílio
    continuacao: aluno
    output:atualiza pagina html
    limitacoes: o aluno deve estar vinculado a alguma bolsa vigente

Solicitar extensão para entrega de algum documento necessário para confirmação da bolsa

    atores: aluno e PROAES
    descricao: formulário e notificar PROAES
    continuacao: aluno e PROAES
    output: notifica proaes e abre stasus de pedido
    limitacoes: o aluno precisa ter aplicado para uma bolsa de auxilio e satisfazer os critérios para tal

Solicitar revisão de indeferimento do pedido

    atores: aluno e PROAES
    descricao: formulário e notifica PROAES
    continuação: PROAES
    output: notifica proaes e abre status de pedido
    limitacoes: o aluno precisa ter tido pedido prévido indeferido

Enviar documentos para verificação de elegibilidade da bolsa

    atores: aluno e PROAES
    descricao: formulário e notifica PROAES
    continuacao: PROAES
    output: notifica proaes e abre status de pedido
    limitacoes:
