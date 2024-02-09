

POST /clientes/:id/transacoes
{
    "valor": 1000, //valor em centavos
    "tipo": "c",
    "descricao": "descricao"
}
Todos os campos são obrigatórios.

Status Code 200
{
    "limite" : 100000,
    "saldo" : -9098
}
