
Abaixo um guia de como documentar decisões de arquitetura
[[Documenting Software Architecture – @hgraca]]

A.Segue abaixo notas que podem melhorar a arquitetura, identificar módulos e coesão da aplicação Note que o input é abstráido para algum tipo de interface que não importa no momento

Consultar status de pedido de algo muit específico para cada subdomínio pode ser tratado como um subsistema
EX: pedido de sair da monitoria, pedido de bolsa sem ser bolsista, pedido emergencial etc - "Consultar status de pedido"

2.feature de notificar usuários do sistema pode ser um subsistema tambem
Ex: notificar aluno, notificar professores, notificar membros de projeto - "Sistema de notificar aluno"

3.Ver descrição é uma feature que pode ser generalizada, vários subdomínios a utiliza, portanto é forte candidato para subsistema
Ex: ver descrição de cadeiras, de bolsa de alguma coisa, de página de projeto etc - "Descrição de algum conceito"

4.Nota-se que tem muita repetição de regras de negócio como checar que o aluno é atualmente vigente em algum papel ou que está vinculado ao curso etc
pode ser um subsitema tambem - "Sistema de checar regras de negócio"

5.assim como 3 filtrar é forte candidato para ser um subsistema - "Sistema de filtragem"

6.Consultar links/atalhos pode ser um subsistema tambem. - "Sitema de links externos"

7.Lidar com formulários aparenta ser algo bem comum, transformemos num subsistema tmb - "Sistema de formulários"

8.Baixar e acessar algum documento virtual parece ser forte candidato para subsistema tmb - "Sistema de download"

9.Não tão forte, mas um notificador visual dos eventos que estão para acontecer é um bom candidato para subsistema tmb, algo como
avaliações que irão aconter em breve, oportunidades de bolsa que irão aparecer etc - "Sistema de dashboard/news"

B. Notar que colocar o status de um pedido é criado ajuda a perceber o que pode ou não ser pedido

C. Seria interessante colocar uma lista de regras de propriedades que deverão ser codificadas
BC- Ensino
1a. Aluno está no período de modificação
1b. Aluno está vinculado ao curso
1c. Aluno não estourou tempo máximo de integralização
1d. Aluno tem prérequisito das cadeiras
1e. Aluno é blocado
1f. Aluno tem prioridade de disciplina
1g. Aluno tem algum status de pedido relacionado ao BC ensino
1h. Aluno está cursando cadeira
1i. Aluno está reprovado por falta
1j. Aluno pode realizar matrícula

BC - Extensão
2a. Aluno está vinculado a projeto de extensão
2b. Aluno já foi membro de projeto de extensão
2c. Aluno já fez review de algum projeto do qual faz parte
2d. Período de review está aberto?
2e. Aluno é bolsista de algum projeto de extensão?

BC - Monitoria
+BC-Extensao

BC -Auxilio

D. Fazer uma aplicação SPA parece ser uma boa para UX do usuário

E. Cada conjunto de funcionalidade compartilhada pode ser seu próprio subsitema 🤔

Onde encaixar poder trancar o curso?
