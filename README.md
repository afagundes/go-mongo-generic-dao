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

Outros exemplos se encontram no arquivo `main.go`

## TODO
Remover a string de conexão com o MongoBD do arquivo `dao.go`
