 
; definir una lista
(define ListaProductos '(("arroz" 8 6000)
                        ("frijoles" 5 5000)
                        ("azucar" 6 1100)
                        ("cafe" 2 2800)
                        ("leche" 9 6000)))

(define (agregarProducto Lista nombre cantidad precio )
  (cond ((null? Lista) (list(list nombre cantidad precio)))
        ((equal? nombre (caar Lista))
         (cons (list (caar Lista) (+ (cadar Lista) cantidad) precio ) (cdr Lista)))
        (else
         (cons (car Lista) (agregarProducto (cdr Lista)
                                              nombre
                                              cantidad
                                              precio)))) 
  )

(define (venderProducto Lista nombre cantidad )
  (cond  ((null? Lista) (displayln "La lista se encuentra vacía." ) '())
         ((equal? nombre (caar Lista))
          (cons (list (caar Lista) (- (list-ref (car Lista) 1) cantidad) (list-ref (car Lista) 2)) (cdr Lista) ))
         (else
         (cons (car Lista) (venderProducto (cdr Lista) nombre cantidad ) ))
        ))


(define(existenciasMinimas Lista cantidad)
  (filter (lambda(x)  (<= (cadr x) cantidad)) Lista))

#|
  Ejercicio a repositorio

|#
(define (comprarProducto inventario Lista nombre cantidad)
  (let ((producto (buscarProducto Lista nombre)))
    (cond ((null? producto)
           (displayln "Producto no encontrado en el inventario.")
           inventario)
          ((>= (cadr producto) cantidad)
           (let ((nuevoInventario (actualizarInventario Lista nombre (- (cadr producto) cantidad))))
             (displayln (format "Compra realizada: ~a unidades de ~a" cantidad nombre))
             nuevoInventario))
          (else
           (displayln(format "No hay suficientes existencias para realizar la compra de ~a unidades de ~a." cantidad nombre))
           inventario))))

(define (buscarProducto Lista nombre)
  (cond ((null? Lista) '())
        ((equal? nombre (caar Lista)) (car Lista))
        (else (buscarProducto (cdr Lista) nombre))))

(define (actualizarInventario Lista nombre nuevaCantidad)
  (cond ((null? Lista) '())
        ((equal? nombre (caar Lista))
         (cons (list nombre nuevaCantidad (caddr (car Lista))) (cdr Lista)))
        (else (cons (car Lista) (actualizarInventario (cdr Lista) nombre nuevaCantidad)))))

(define (realizarCompras Lista compras)
  (define (realizarCompra compra inventario)
    (let ((nombre (car compra))
          (cantidad (cadr compra)))
     (let ((nuevoInventario (comprarProducto inventario Lista nombre cantidad)))
        (if (equal? nuevoInventario inventario)
            inventario ; No se realizó la compra, devuelve el inventario sin cambios
            (let ((producto (buscarProducto Lista nombre)))
              (cons (list nombre cantidad (caddr producto)) inventario))))))

  ; acumula los productos comprados
  (let ((productosComprados (foldl realizarCompra '() compras)))
    productosComprados))


(define montoImpuesto 0.13) ; se define el impuesto a aplicar.
; Se pide que el rangoImpuesto sea argumento de la funcion facturaVenta y no un valor ya definido.
(define (calcularImpuesto Lista rangoImp)
  (cond ((null? Lista) '())
        (else (cons  (if (>= (list-ref(car Lista) 2)rangoImp) (*(list-ref(car Lista) 2) montoImpuesto) 0)
                (calcularImpuesto(cdr Lista) rangoImp)))))

(define (sumarImpuesto Lista)
  (cond ((null? Lista) 0  )
        (else(+ (car Lista) (sumarImpuesto (cdr Lista))))))


(define (facturaVenta Lista rangoImpuesto)
  (cond  ((null? Lista) (displayln "La lista se encuentra vacia.") '())
         (else (begin (display(format"Factura de venta \nLista de compras-->~a" Lista) )
             (display "\nEl total de impuesto es de: ") (let ((impuesto (sumarImpuesto(calcularImpuesto Lista rangoImpuesto)))) (display impuesto)
             (display "\nEl total de la factura sin impuestos: ") (let((total-factura(foldl (lambda (x sum) (+ (list-ref x 2) sum))0 Lista)))
                                                                  (display total-factura)
             (display "\nTotal de la Factura: ") (display(+ impuesto total-factura) )))
         ))))
#|
       Ejecución del Código
      Welcome to DrRacket, version 8.12 [cs].
      Language: Pretty Big; memory limit: 128 MB.
      > (facturaVenta(realizarCompras(agregarProducto(agregarProducto ListaProductos "fideos" 10 3200) "queso" 25 2600) '(("frijoles" 3) ("leche" 5) ("arroz" 2))) 5000)
       Compra realizada: 3 unidades de frijoles
       Compra realizada: 5 unidades de leche
       Compra realizada: 2 unidades de arroz
       Factura de venta 
       Lista de compras-->((arroz 2 6000) (leche 5 6000) (frijoles 3 5000))
       El total de impuesto es de: 2210.0
       El total de la factura sin impuestos: 17000
       Total de la Factura: 19210.0
       >

   Caso de que la compra de un producto sea mayor a la que hay en ListaProductos no se permite.
   > (facturaVenta(realizarCompras(agregarProducto(agregarProducto ListaProductos "fideos" 10 3200) "queso" 25 2600) '(("frijoles" 6) ("leche" 5) ("arroz" 2) ("cafe" 1)("queso" 2))) 5000)
    No hay suficientes existencias para realizar la compra de 6 unidades de frijoles.
     Compra realizada: 5 unidades de leche
     Compra realizada: 2 unidades de arroz
     Compra realizada: 1 unidades de cafe
     Compra realizada: 2 unidades de queso
     Factura de venta 
     Lista de compras-->((queso 2 2600) (cafe 1 2800) (arroz 2 6000) (leche 5 6000))
     El total de impuesto es de: 1560.0
     El total de la factura sin impuestos: 17400
     Total de la Factura: 18960.0
     > 
|#
 
