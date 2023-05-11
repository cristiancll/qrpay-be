package errors

const (
	CONNECTION_ERROR        string = "Erro de conexão"
	UUID_REQUIRED           string = "UUID é obrigatório"
	PHONE_REQUIRED          string = "Telefone é obrigatório"
	PASSWORD_REQUIRED       string = "Senha é obrigatória"
	NAME_REQUIRED           string = "Nome é obrigatório"
	AUTH_ERROR              string = "Erro de autenticação"
	INTERNAL_ERROR          string = "Erro interno"
	INVALID_PASSWORD        string = "Senha inválida"
	INVALID_CREDENTIALS     string = "Usuario ou senha incorretos"
	DISABLED_USER           string = "Usuário desabilitado"
	VERIFICATION_ERROR      string = "Usuário não verificado"
	USER_ALREADY_EXISTS     string = "Usuário já existe"
	WHATSAPP_ALREADY_EXISTS string = "Whatsapp já existe"
	DATABASE_ERROR          string = "Erro de banco de dados"

	USER_NOT_FOUND string = "Usuário não encontrado"
	NO_USERS_FOUND string = "Nenhum usuário encontrado"

	WHATSAPP_NOT_FOUND string = "Whatsapp não encontrado"
	NO_WHATSAPP_FOUND  string = "Nenhum whatsapp encontrado"

	UNAUTHORIZED string = "Não autorizado"
)