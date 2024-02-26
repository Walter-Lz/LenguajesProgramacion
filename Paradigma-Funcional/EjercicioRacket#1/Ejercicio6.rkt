; Ejercicio #6 Ordenamiento de listas

(define (merge lista1 lista2 )
  (cond ((null? lista1) (sort lista2 < ))
        ((null? lista2) (sort lista1 < ))
        (else (sort  (append lista1 lista2)<))
    )
 )

#|
         Ejecución del código
      
      Welcome to DrRacket, version 8.12 [cs].
      Language: Pretty Big; memory limit: 128 MB.
      > (merge '(1 2 3 4) '(5 6 7 8))
       (1 2 3 4 5 6 7 8)
      > (merge '(1 2 3) '(1 2 3 4))
       (1 1 2 2 3 3 4)
      > 

|#