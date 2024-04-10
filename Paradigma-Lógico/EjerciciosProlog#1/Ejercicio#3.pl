
distanciaH(Str1,Str2,N):-
    % se pasan a listas para recorrerlas.
    string_chars(Str1,Chars1),
    string_chars(Str2,Chars2),
    distanciaHaming(Chars1,Chars2,N),!.


distanciaHaming([],[],0). % caso base que sean vacias
distanciaHaming([],[_|_],0).% caso base de que la primera palabra llegue a ser vacía
distanciaHaming([_|_],[],0).% caso base de que la segunda palabra llegue a ser vacía
distanciaHaming([Hchar|T1],[Hchar2|T2],N):-
    Hchar \= Hchar2,
    distanciaHaming(T1,T2,N2),
    N is N2 +1.
distanciaHaming([_|T1],[_|T2],N):- distanciaHaming(T1,T2,N).%caso base de que el primer caracter en ambas palabras coincidan.

%       Ejecución del código.
%    ?- distanciaH("romano","comino",X).
%     X = 2.
%    ?- distanciaH("romano","camino",X).
%     X = 3.
%    ?- distanciaH("roma","comino",X).
%     X = 2.
%    ?- distanciaH("romano","ron",X).
%     X = 1.
%    ?- distanciaH("romano","cama",X).
%     X = 2.
%    ?-
