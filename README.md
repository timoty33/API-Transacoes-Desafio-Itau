# Desafio de vaga BackEnd junior do Itau resolvido em Golang

Desafio de vaga BackEnd Júnior da Instituição Itau. O desafio original, tinha como objetivo ser feito em Java ou Kotlin, porém nesta solução iremos usar Golang com Gin.

Desafio resolvido e testado com ThunderClient.

Modificações:

- Ajuste de Segurança no Endpoint DELETE /transacao, agora ele requer senha, se tornando DELETE /transacao/:senha
- Desafio feito em GO com Gin ao invés de JAVA com SpringBoot

EXTRAS:

- Health Check na rota GET /health
