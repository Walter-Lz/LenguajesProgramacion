;Ejercicio #1


(define (calcular-capital capital interes año)
  (cond ((equal? año 0)(displayln(format "Después de ~a años el monto sería de: ~a " año  capital )))
    (else(displayln(format "Después de ~a años el monto sería de: ~a  " año(round (* capital(expt (+ 1 interes) año))))) ; se calcula los interes del capital y se imprime en consola.
       (calcular-capital capital interes (- año 1))) ; se realiza recursividad para continuar el informe de cada año.
   ) 
)

#|
         Ejecución del código
      Welcome to DrRacket, version 8.12 [cs].
      Language: Pretty Big; memory limit: 128 MB.
      > (calcular-capital 2000 0.10 3)
       Después de 3 años el monto sería de: 2662.0  
       Después de 2 años el monto sería de: 2420.0  
       Después de 1 años el monto sería de: 2200.0  
       Después de 0 años el monto sería de: 2000 
      > 
|#