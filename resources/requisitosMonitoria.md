
glossário:
    atores:= personas que interagem com o sistema de alguma maneira
    descricao:=o que aquela funcionalidade deve permitir em linguagem natural
    continuacao:=personas responsáveis por progredir com o fluxo da funcionalidade
    output:saída do sistema
    limitacoes:o que é válido para qualquer fluxo daquela funcionalidade sempre
input não foi mencionado visto que é parte da UI e UX, pode ser abstraído no momento

-Monitoria
Quais informações mais importantes para esse BC?
Filtrar projetos de monitoria disponíveis baseado em certos critérios

atores: aluno
descricao:  filtrar por bolsas de monitorias vigentes
continuacao: aluno
output: atualiza pagina hmtl
limitações apenas de projetos vigentes e que o aluno tem os pré-requistios para ser monitor

Aplicar para projetos de monitoria

atores: aluno e professor
descricao: formulário e notificar professores associados
continuacao: aluno
output: notificação aos professores responsáveis
limitações: o aluno não pode ter curso trancado.

Ver projetos de monitoria que está associado

atores: aluno
descricao: consultar projetos de monitoria que está associado
continuacao: aluno
output: atualiza pagina html
limitações: apenas de projetos vigentes e que está associado

Ver a descrição de projetos de monitoria

atores: aluno
descricao: exibir pagina html para aluno
continuacao: aluno
output: atualiza pagina html
limitações: apenas de projetos vigentes

Sair de um projeto de monitoria

atores: aluno e professores
descricao: formulário e notificar professores
continuacao: aluno e professor
output: professores são notificados
limitações: apenas de projetos que o aluno participa atualmente

Fazer relatórios de monitoria

atores: aluno e  membros de projeto
descricao: forncere formulário para aluno e notificar monitores
continuacao:aluno e líder de projeto
output: o líder de projeto deverá validar relatório e monitores são notificados
limitacoes: o período que essa feature fica disponibilizada é de acordo com o cronograma do período

solicitar bolsas para o atual projeto de monitoria, caso já não tenha

atores: aluno e professores
descricao: fornece formulário e notifica professores
continuacao: aluno e professores
output: os professores são notificados e o status do pedido é criado
limitacoes: o aluno só pode solicitar se for membro de um único projeto de monitoria naquele período

Dar uma nota informal dos projetos de monitoria nos quais participou

atores: aluno e professores
descricao: fornece formulário e notifica professores
continuacao: aluno e professores
output: os professores são notificados e devem validar o review
limitacoes: o aluno precisa ser membro até o final do período antes de fazer um review

Acesso a página da disciplina/cronograma/da turma que é monitor

atores: aluno
descricao: fornece link para pagina html
continuacao: aluno
output: acessa pagina html
limitacoes: apenas de projetos de monitoria vigentes

Acompanhar status de pedidos de aplicação

atores: aluno
descricao: pagina htmml
continuacao: aluno
output:atualiza pagina html
limitacoes: apenas de pedidos vigentes
