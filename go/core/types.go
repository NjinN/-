package core

const (
	NIL			= iota
	NONE		
	ERR 		
	LIT_WORD	
	GET_WORD 	
	DATATYPE 	
	LOGIC 		
	INTEGER 	
	DECIMAL 	
	CHAR	  	
	STRING 		
	PAREN 		
	BLOCK
	OBJECT
	PROP
	WORD
	SET_WORD
	PUT_WORD 		
	PATH 			
	OP 			
	NATIVE 		
	FUNC

	UNDEFINED 		
)

func TypeToStr(n int) string{
	switch n {
	case NIL:
		return "nil!"
	case NONE:
		return "none!"
	case ERR:
		return "error!"
	case LIT_WORD:
		return "lit-word!"
	case GET_WORD:
		return "get-word!"
	case DATATYPE:
		return "datatype!"
	case LOGIC:
		return "logic!"
	case INTEGER:
		return "integer!"
	case DECIMAL:
		return "decimal!"
	case CHAR:
		return "char!"
	case STRING:
		return "string!"
	case PAREN:
		return "paren!"
	case BLOCK:
		return "block!"
	case OBJECT:
		return "object!"
	case PROP:
		return "prop!"
	case PATH:
		return "path!"
	case WORD:
		return "word!"
	case SET_WORD:
		return "set-word!"
	case PUT_WORD:
		return "put-word!"
	case OP:
		return "op!"
	case NATIVE:
		return "native!"
	case FUNC:
		return "function!"
		
	default:
		return "UNDEFINED!!!"
	}
}

func StrToType(s string) int{
	switch s {
	case "nil!":
		return NIL
	case "none!":
		return NONE
	case "err!":
		return ERR
	case "lit-word!":
		return LIT_WORD
	case "get-word!":
		return GET_WORD
	case "datatype!":
		return DATATYPE
	case "logic!":
		return LOGIC
	case "integer!":
		return INTEGER
	case "decimal":
		return DECIMAL
	case "char!":
		return CHAR
	case "string!":
		return STRING
	case "paren!":
		return PAREN
	case "block!":
		return BLOCK
	case "object!":
		return OBJECT
	case "prop!":
		return PROP
	case "word!":
		return WORD
	case "set-word!":
		return SET_WORD
	case "put-word!":
		return PUT_WORD
	case "path!":
		return PATH
	case "op!":
		return OP
	case "native!":
		return NATIVE
	case "func!":
		return FUNC

	default:
		return UNDEFINED

	}


}
