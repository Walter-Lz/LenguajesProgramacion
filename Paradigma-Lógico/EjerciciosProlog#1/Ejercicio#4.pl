restaurante(guacamole,200,entrada).
restaurante(ensalada,150,entrada).
restaurante(consome,300,entrada).
restaurante(tostadas_caprese,250,entrada).

restaurante(filete_cerdo,400,carne).
restaurante(pollo_al_horno,280,carne).
restaurante(carne_en_salsa,320,carne).

restaurante(tilapia,160,pescado).
restaurante(salmon,300,pescado).
restaurante(trucha,225,pescado).

restaurante(flan,200,postre).
restaurante(nueces_con_miel,500,postre).
restaurante(naranja_confitada,450,postre).
restaurante(flan_de_coco,375,postre).

combinaciones(Calorias, Combinaciones) :-
    retractall(elegido(_)),
    generar_combinaciones(Calorias, 5, [], Combinaciones).

generar_combinaciones(_, 0, Combinaciones, Combinaciones). % caso base de que se hayan hecho las 5 combinaciones
generar_combinaciones(Calorias, N, CombinacionesParciales, Combinaciones) :-
    N > 0,
    combinacion(Calorias, Combinacion),
    append(CombinacionesParciales, [Combinacion], NuevaLista),
    N1 is N - 1,
    generar_combinaciones(Calorias, N1, NuevaLista, Combinaciones).

combinacion(Calorias, Combinacion) :-
    findall(Plato, restaurante(Plato, _, _), TodosPlatos),  % se buscan todos los platos
    generar_combinacion_sin_repetir(TodosPlatos, Calorias, [], Combinacion). % se hacen la busqueda de combinaciones

generar_combinacion_sin_repetir([], _, Combinacion, Combinacion). % caso de que ya no queden platos
generar_combinacion_sin_repetir(_, 0, Combinacion, Combinacion). % caso de que ya se haya alcanzado el rango de caloria
generar_combinacion_sin_repetir([Plato|Resto], Calorias, CombinacionParcial, Combinacion) :-
    restaurante(Plato, CaloriasPlato, _),
    Calorias >= CaloriasPlato,
    \+ elegido(Plato),
    append(CombinacionParcial, [Plato], NuevaCombinacion),
    NuevasCalorias is Calorias - CaloriasPlato,
    asserta(elegido(Plato)),
    generar_combinacion_sin_repetir(Resto, NuevasCalorias, NuevaCombinacion, Combinacion),
    !. % Cortar para evitar que intente generar más combinaciones


generar_combinacion_sin_repetir([_|Resto], Calorias, CombinacionParcial, Combinacion) :- %caso de que el plato no cumpla para ser seleccionado.
    generar_combinacion_sin_repetir(Resto, Calorias, CombinacionParcial, Combinacion).


%       Ejecución del código
%     ?- combinaciones(1200,X).
%     X = [[guacamole, ensalada, consome, tostadas_caprese,
%          pollo_al_horno], [filete_cerdo, carne_en_salsa, tilapia,
%          salmon], [trucha, flan, nueces_con_miel], [naranja_confitada,
%          flan_de_coco], []] .
%     ?-
%
%     como es de manera recursiva se queda esperando como si existieran
%     más resultados, pero como se puede apreciar la siguiente solicitud
%     ya es el final de la consulta.
%
%     ?- combinaciones(1200,X).
%     X = [[guacamole, ensalada, consome, tostadas_caprese,
%       pollo_al_horno], [filete_cerdo, carne_en_salsa, tilapia,
%       salmon], [trucha, flan, nueces_con_miel], [naranja_confitada,
%       flan_de_coco], []] ;
%     false.
%     ?-
