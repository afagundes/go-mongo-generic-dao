# go-mongo-generic-dao
DAO genérico para MongoBD escrito em GO

## Utilização

Basta inicializar a estrutura DAO passando o nome do banco de dados e da coleção.

```go
usuarioDAO := dao.DAO{ Database: DATABASE, Collection: COLLECTION }
```

Após isso já é possível interagir com o MongoBD através dos métodos do DAO.

## Exemplos

Vamos tomar como exemplo a estrutura Usuario:

```go
type Usuario struct {
  ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
  Nome string `bson:"nome" json:"nome,omitempty"`
  Idade int `bson:"idade" json:"idade,omitempty"`
  Bio string `bson:"bio" json:"bio,omitempty"`
  Foto string `bson:"foto" json:"foto,omitempty"`
}
```

Para recuperar todos os usuários cadastrados podemos fazer:

```go
var usuarios []model.Usuario
usuarioDAO.GetAll(&usuarios)

for _, usuario := range usuarios {
  fmt.Printf("%+v\n", usuario)
}
```

Para inserir um novo usuário:

```go
usuario = model.Usuario{
  Nome: "Ari", 
  Idade: 34, 
  Bio: "Testando inserção de usuário", 
  Foto: "https://pbs.twimg.com/profile_images/767051707389992963/w0x4-LIs_400x400.jpg"}

objectID := usuarioDAO.Insert(usuario)

id = objectID.(primitive.ObjectID).Hex()
fmt.Printf("Inseriu novo usuário %s\n", id)
```

## Configuração

Para configurar o acesso ao MongoDB é necessário alterar o arquivo `config\constants.go` e adicionar a string de conexão.

## API para testes

Incluí uma API que disponibiliza um CRUD de usuários na porta 8080. Dessa forma é possível ver os métodos do DAO em ação.

Os seguintes endpoints foram adicionados:

<table>
<tr>
<th>Endpoint</th>
<th>Método</th>
<th>Descrição</th>
</tr>
<tr>
<td>/usuario</td>
<td>POST</td>
<td>Adiciona um usuário</td>
</tr>
<tr>
<td>/usuarios</td>
<td>GET</td>
<td>Retorna a lista de usuários cadastrados</td>
</tr>
<tr>
<td>/usuario/{id}</td>
<td>GET</td>
<td>Retorna um usuário pelo seu ID</td>
</tr>
<tr>
<td>/usuario/{id}</td>
<td>PUT</td>
<td>Atualiza um usuário (exemplo de JSON abaixo)</td>
</tr>
<tr>
<td>/usuario/{id}</td>
<td>DELETE</td>
<td>Remove um usuário pelo seu ID</td>
</tr>
</table>

O JSON utilizado nos endpoints POST e PUT é o seguinte:

```json
{
  "nome": "Maria",
  "idade": 30,
  "bio": "Testando a API com Mongo Generic DAO",
  "foto": "https://image.flaticon.com/icons/png/512/25/25231.png"
}
```
