glossário:
    atores:= personas que interagem com o sistema de alguma maneira

    descricao:=o que aquela funcionalidade deve permitir em linguagem natural

    continuacao:=personas responsáveis por progredir com o fluxo da funcionalidade

    output:saída do sistema

    limitacoes:o que é válido para qualquer fluxo daquela funcionalidade sempre
input não foi mencionado visto que é parte da UI e UX, pode ser abstraído no momento

- Ensino
Quais informações mais importantes para esse BC?
- Agora vamos associar a cada feature atores, o comportamento de sistemas e requisitos de input e output
A nota de uma cadeira.

atores: aluno
descricao: deve permitir download e consulta das notas
continuacao: aluno
output: disponibiliza pagina html
limitacoes:o sistema só deve mostrar as notas de cadeiras que o aluno já cursou ou está cursando

A matrícula em um conjunto de cadeiras.

atores: aluno
descricao: deve permitir filtro de cadeiras, confirmação de matricula, status de acompnhamento da matricula.
continuacao:aluno
output:atualiza página html e status do pedido é aberto
limitacoes: o aluno pode realizar a matricula várias veze durante o período de matrícula, se o aluno estiver com o curso trancado não podera realizar matricula naquele período

Acesso ao histórico escolar.

atores: aluno
descricao:deve permitir download e consulta
continuacao: aluno
output: disponibilza pagina html
limitacoes: o historico escolar apresenta cadeiras de todo periodo que o aluno ja cursou ou esta cursando

Quando terá uma avaliação.

atores: aluno
descricao:deve permitir consulta e notificar usuário
continuacao: aluno
output: atualiza pagina html
limitacoes: o sistema só deve notificar o aluno de cadeiras que esta atualmente cursando

Quantas faltas ainda são possíveis numa cadeira e evitar reprovação.

atores: aluno
descricao: deve permitir consulta por cadeira e notificar usuário
continuacao: aluno
output: atualiza pagina html
limitcaoes: o sistema só deve filtrar por cadeiras que o aluno está cursando atualmente

Cronograma das cadeiras.

atores: aluno
descricao: consultar e fazer download do documento
continuacao: aluno
output: atualiza pagina html
limitacoes: apenas do período atual

Sala de aula e horário das cadeiras.

atores: aluno
descricao: consultar por cadeiras
continuacao: aluno
output: atualiza pagina html
limitacoes: apenas de cadeiras que o aluno está cursando atualmente

Poder trancar o curso.

atores: aluno
descricao: trancar o curso é uma ação irreversível durante um período
continuacao: aluno
output: atualiza página html e status do pedido é aberto
limitacoes: o aluno só pode realizar essa acao duas vezes durante a graduação, na vez seguinte o aluno perde o vinculo com a instituição de ensino, enquanto não for validada ele pode cancelar o pedido, demora 24h a validação

Emitir declaração de vínculo

atores: aluno
descricao: consultar e baixar documento
continuacao: aluno
output: disponibilza pagina html
limitacoes: nenhuma óbvia

Consultar calendário acadêmico

atores: aluno
descricao: consultar e baixar documento
continuacao: aluno
output: disponibiliza pagina html
limitacoes: deve ser apresentado apenas o cronograma do ano atual

Checar a estrutura curricular de cursos, quais cadeiras obrigatórias, eletivas, ch e relacação de depedência entre as mesmas
Solicitar transferência interna/externa
