(define (contiene-subcadena? cadena subcadena)
  (not (equal? #f (regexp-match? (regexp (string-append ".*" subcadena ".*")) cadena))))

(define (filtrar_subcadena ListaPalabras subcadena)
  (filter (lambda (cadena) (contiene-subcadena? cadena subcadena)) ListaPalabras))


#|
    Ejecución del Código
   Welcome to DrRacket, version 8.12 [cs].
   Language: Pretty Big; memory limit: 128 MB.
    > (filtrar_subcadena '("la casa" "La mora" "el perro" "pintando la cerca") "la")
     ("la casa" "pintando la cerca")
    > 
 
|#