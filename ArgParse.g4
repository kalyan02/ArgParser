grammar ArgParse;

//--args-what --key=val -arg-val=foo -arg -test -t4 --what "hello" 

args 
    : arg ( SPACE arg )* 
    ;

arg 
    : dash_opt_kv_eq 
    | dash_opt_kv 
    | dash_opt_k_bool
    ;

dash_opt_k_bool
    : DASH key
    ;

dash_opt_kv 
    : DASH key SPACE val 
    ;

dash_opt_kv_eq
    : DASH key '=' val 
    ;

key 
    : TEXT 
    ;

val 
    : TEXT 
    | STRING 
    ;
DASH 
    : '--' 
    | '-' 
    ;

SPACE
    : ' '+
    ;

NEWLINE
    : [\r\n]+
    -> skip
    ;


TEXT : [A-Za-z0-9_]+ ;

STRING
   : '"' (ESC | ~ ["\\])* '"'
   ;


fragment ESC
   : '\\' (["\\/bfnrt] | UNICODE)
   ;


fragment UNICODE
   : 'u' HEX HEX HEX HEX
   ;


fragment HEX
   : [0-9a-fA-F]
   ;


