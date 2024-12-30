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

---

# Thread e Goroutine

Goroutine: É uma abstração leve sobre threads, gerenciadas pelo runtime do Go. Ela é muito mais
leve e eficiente em termos de recursos do que as threads tradicionais do sistema operacional. O Go
pode executar milhares de goroutines em um único programa sem grandes impactos de memória ou
processamento.

Thread: Uma thread é uma unidade de execução do sistema operacional. Criar e gerenciar threads
pode ser custoso, pois o sistema operacional precisa reservar mais recursos (como memória e tempo
de CPU) para cada uma.

## Principais diferenças

### Leveza

As goroutines são muito mais leves que as threads tradicionais. Cada goroutine começa com um
stack de apenas 2 KB, enquanto uma thread do sistema operacional normalmente inicia com um
stack de 1 MB ou mais. Isso significa que você pode criar milhares de goroutines sem problemas de
consumo de memória, algo impraticável com threads devido ao alto custo de criação e consumo de memória.

### Escalonamento

No Go, o runtime gerencia as goroutines e as distribui automaticamente para um número limitado
de threads do sistema operacional, usando um escalonador cooperativo. Isso significa que o runtime
decide quando cada goroutine deve ser pausada ou retomada. Em contraste, as threads são
escalonadas diretamente pelo sistema operacional, o que pode ser menos eficiente, especialmente
em programas com muitas threads.

### Múltiplas goroutines por thread

Em um thread do sistema operacional, apenas uma tarefa pode ser executada por vez. O runtime
do Go permite que várias goroutines sejam executadas dentro de uma única thread e, quando
necessário, ajusta dinamicamente o uso das threads. Isso economiza recursos e reduz a sobrecarga
em comparação com as threads, onde cada tarefa exige uma thread exclusiva e um contexto
separado.

### Gerenciamento pelo Go runtime

As goroutines são interamente gerenciadas pelo runtime de Go, enquanto as threads são
gerenciadas pelo sistema operacional. Esse gerenciamento direto pelo Go oferece mais controle e
eficiência para lidar com a concorrência. Como resultado, o uso de goroutines é mais leves e
simplificado, permitindo que o runtime ajuste as goroutines conforme as necessidades do programa
sem depender diretamente das capacidades do sistema operacional.

> Portanto, goroutines não são threads diretamente, mas são executadas sobre threads,
oferecendo uma forma mais eficiente e escalável e lidar com a concorrência em comparação
com o uso direto de threads.

---
