# Concorrência

Concorrência é o termo da ciência da computação para dividir um único processo em componentes
independentes e especificar como esses componentes compartilham dados com segurança. A
maioria das linguagens fornece essa simultaneidade por meio de uma biblioteca usando threads de
nível de sistema operacional que compartilham dados tentando adquirir bloqueios.

Mas em Go é diferente. Seu principal modelo de simultaneidade, sem dúvida o recurso mais famoso
de Go, é baseado em Processos Sequenciais de Comunicação (CSP). Esse estilo de simultaneidade foi
descrito em 1978 em um artigo de [Tony Hoare (artigo em inglês)](https://dl.acm.org/doi/pdf/10.1145/359576.359585), o homem que inventou o
algoritmo [Quicksort](https://pt.wikipedia.org/wiki/Quicksort). Os padrões implementados com CSP são tão poderosos quanto os padrões, mas
são muito mais fáceis de entender.

Vou criar uma analogia deste artigo, para melhor entendimento. Imagine que você está organizando
uma festa de aniversário e precisa decidir:

- **Variáveis:** Quem convidar (cada pessoa é uma variável), o que servir (cada comida é uma
variável), que música tocar (cada música é uma variável).
- **Domínios:** Para cada variável, você tem um conjunto de opções possíveis (por exemplo, para
"quem convidar", você tem uma lista de amigos, para "o que servir", você tem uma lista de
comidas e para "que música tocar", você tem uma lista de músicas).
- **Restrições:** Existem regras que limitam suas escolhas. Por exemplo, você só pode convidar
pessoas que você sabe que se dão bem (restrição entre a variável "quem convidar" e outras
variáveis). Você também pode ter restrições quanto ao preço da comida, por exemplo.

O CSP tenta encontrar um conjunto de valores para todas as variáveis que satisfaçam todas as
restrições. Para isso, ele utiliza algoritmos de busca que exploram as possíveis combinações de
valores para as variáveis.

> Em resumo, o CSP é uma técnica que ajuda a resolver problemas complexos com muitas
variáveis e restrições, encontrando a melhor combinação de soluções que satisfazem todas as
condições.

O [artigo de Tony](https://pt.wikipedia.org/wiki/Charles_Antony_Richard_Hoare) destaca a importância do trabalho de Hoare no desenvolvimento de algoritmos de
ordenação e pesquisa, como o Quicksort. Embora o Quicksort não seja diretamente aplicado na
resolução de CSPs, seus princípios de dividir e conquistar e recorrência podem ser considerados
relevantes no contexto da Busca em Profundidade com Backtracking, onde o espaço de busca é
dividido recursivamente em subproblemas menores que são resolvidos individualmente.

Muitas pessoas acreditam que adicionar concorrência automaticamente torna os programas mais
rápidos, mas isso nem sempre é verdade. A concorrência é uma ferramenta que ajuda a estruturar
problemas complexos, mas não necessiariamente leva à execução paralela, que depende do
hardware e das condições do algoritmo. É importante distinguir a concorrência de paralelismo:
enquanto a primeira organiza o fluxo de trabalho, a segunda refere-se à execução real de múltiplas
tarefas ao mesmo tempo.

Em termos gerais, todos os programas seguem três etapas principais: capturam dados, processam
essas informações e, finalmente, geram o resultado. A decisão  de usar concorrência em um
programa depende de como os dados se movem entre essas etapas. Em algumas situações, duas
etapas podem ser executadas em paralelom pois uma não depende dos dados gerados pela
anterior. A concorrência é útil quando é preciso combinar dados de várias operações que podem ser
realizadas de forma independente.
