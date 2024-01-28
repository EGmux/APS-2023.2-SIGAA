
glossário:
    atores:= personas que interagem com o sistema de alguma maneira

    descricao:=o que aquela funcionalidade deve permitir em linguagem natural

    continuacao:=personas responsáveis por progredir com o fluxo da funcionalidade

    output:saída do sistema

    limitacoes:o que é válido para qualquer fluxo daquela funcionalidade sempre
input não foi mencionado visto que é parte da UI e UX, pode ser abstraído no momento

- Extensão
Quais informações mais importantes para esse BC?
Filtrar projetos de extensão disponíveis baseados em certos critérios

    atores: aluno
    descricao: filtrar projetos
    continuacao: aluno
    output: atualiza pagina html
    limitacoes: projetos que já terminaram ou estão em execução podem ser mostrados

Aplicar para projetos de extensão

    atores: aluno líder de projeto
    descricao: formulário de inscrição e notificação para líder de projeto
    continuacao: aluno
    output: líder de projeto é notificado e status do pedido é aberto
    limitacoes: apenas de projetos vigentes

ver projetos de extensão que está associado

    atores: aluno
    descricao: consultar projetos de extensao que ja participou ou participa
    continuacao: aluno
    output: atualiza pagina html
    limitacoes: apenas de projetos que ja foi/e membro

Ver a descrição de um projeto de extensão

    atores: aluno
    descricao: exibir mais detalhes de um projeto de extensão
    continuacao: aluno
    output: atualiza página html
    limitacoes: todos os projetos são candidatos a esse evento

Sair de um projeto de extensão

    atores: aluno e líder de projeto
    descricao: prover um formulário que o aluno preenche e é enviado para o líder de projeto
    continuacao: aluno
    output: formulário e notificação para líder de projeto que deve abrir um novo status de pedido
    limitações: apenas de projetos que o aluno é membro vigente

Acessar lates de outros membros do projeto de extensão

    atores: aluno
    descricao: fornecer link para perfil acadêmico de membros de um projeto que o aluno esteja ou não inserido
    continuacao: aluno
    output: acessa  pagina html
    limitacoes: somente links da plataforma lates

Fazer uma análise da experiencia ao ter passado pelo projeto de extensão, um review público

    atores: aluno e líder de projeto
    descricao: formulário de avaliação e notificar membros de projeto
    continuacao: aluno
    output: formulário e notificação para membros  de projeto
    limitações: somente de projetos que o aluno já foi membro ou é atualmente membro, o aluno só pode atualizar seu review enquanto estiver no status de membro do projeto vigente

Acessar homepage/perfil público de outros membros que particiam de um projeto de extensão seja o aluno associado ou não

    atores: aluno
    descricao: links para diversas homepages que não são lattes
    continuacao: aluno
    output: acessa pagina html
    limitações: o aluno pode consultar links de projetos vigentes ou arquivados

Acessar homepage do projeto de extensão seja o aluno associado ou não

    atores: aluno
    descricao: link para pagina do projeto de extensão
    continuacao: aluno
    output: acessa pagina html
    limitações: mesma da feature acima

Solicitar bolsas para algum projeto de extensão que esteja inserido, caso já não tenha

    atores: aluno e líder de projeto
    descricao: formulário e notificar líder do projeto
    continuacao: aluno
    output: formulário e notificação ao líder de projeto
    limitações: apenas de projeto que o aluno é membro vigente e não é bolsista

Acompanhar status do pdido de aplicação

    atores: aluno e líder de projeto
    descricao: atualiza pagina html e notificar aluno
    continuacao: líder de projeto
    output:o líder de projeto atualiza o status do pedido de ser membro do projeto
    limitações: apenas alunos que aplicaram para ser membros podem ter pedido atualizado
