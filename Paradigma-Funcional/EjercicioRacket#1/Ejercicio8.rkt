; Ejercicio #8

(define (pertenece-elemento? elemento conjunto) 
  (cond ( (null? conjunto ) #f)                 
        ((equal? (car conjunto) elemento)#t)        ; se valida si el primer elemento del conjunto equivale al elemento.
        (else (pertenece-elemento? elemento (cdr conjunto) ))  ; se recorre el conjunto para validar.
        
     )
  )

(define (sub-conjunto? conjuntoA conjuntoB)
  (cond ( (null? conjuntoA) #t)  ; caso de que el conjunto sea vacio
        ( (pertenece-elemento?  (car conjuntoA) conjuntoB ) (sub-conjunto? (cdr conjuntoA) conjuntoB)) ; pasa a comprobar el elemento del conjunto A del B
        (else #f))
  )
#|
            Ejecución del código
       
       Welcome to DrRacket, version 8.12 [cs].
       Language: Pretty Big; memory limit: 128 MB.
       > (sub-conjunto? '() '(a b c d e f))
        #t
       > (sub-conjunto? '(a b c) '(a b c d e f))
        #t
       > (sub-conjunto? '(a b x) '(a b c d e f))
        #f
       > 
|#