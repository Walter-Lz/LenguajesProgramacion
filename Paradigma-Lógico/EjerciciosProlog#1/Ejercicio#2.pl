aplanar([],[]).
aplanar([H|T],X):-
    (is_list(H) ->
      aplanar(H,HPlano),
      aplanar(T,TPlano),
      append(HPlano, TPlano, X);
        X = [H|TPlano],
        aplanar(T,TPlano)
    ).

%         Ejecución del código
%
%    ?- aplanar([1,2,[3,[4,5],[6,7]]],X).
%    X = [1, 2, 3, 4, 5, 6, 7].
%    ?-
