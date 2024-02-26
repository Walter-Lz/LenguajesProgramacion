; Ejercicio #9


(define (eliminar_elemento elem lista)
  (apply append (map (lambda(x)(cond ((equal? elem x)'())   ;Para cada elemento igual se coloca una lista vacia 
                               (else (list x))))            ; distinto se conserva, y se aplica un apply para generar una nueva lista
                           lista))
)
#|
            Ejecución del código

        Welcome to DrRacket, version 8.12 [cs].
        Language: Pretty Big; memory limit: 128 MB.
        > (eliminar_elemento 3 '(1 2 3 4 5))
         (1 2 4 5)
        > (eliminar_elemento 0 '(1 2 3 4 5))
         (1 2 3 4 5)
        > (eliminar_elemento 2 '(1 2 3 2 4 2 5 6))
          (1 3 4 5 6)
        > 
|#