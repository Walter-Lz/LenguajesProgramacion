sub_conjunto([],L2):- is_list(L2). %Validar que sea tipo lista.
sub_conjunto([H|T],L):-
    miembro(H,L),
    sub_conjunto(T,L).

miembro(E,[E|_]):- !.
miembro(E,[_|Tail]) :-
        miembro(E,Tail).


%        Ejecución del código
%    ?- sub_conjunto([],[1,2,3,4,5]).
%    true.
%
%    ?- sub_conjunto([1,2,3],[1,2,3,4,5]).
%    true.
%
%    ?- sub_conjunto([1,2,6],[1,2,3,4,5]).
%    false.
%
%    ?-
%
