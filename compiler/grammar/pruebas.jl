function fibonacci(numero::Int64)::Int64
    if (numero > 1)
        f1 = fibonacci(numero-1);
        f2 = fibonacci(numero-2);
        return f1+f2;
    elseif (numero == 1)
        return 1;
    elseif (numero == 0)
        return 0;
    end;
end;

function impares(numero :: Int64)::Int64
    if  (numero == 1)
        return 1;
    elseif (numero % 2 == 1)
        aux = impares(numero -2);
        return numero *aux;
    end;
    aux2 = impares(numero-1);
    return aux2;
end;

println(impares(7));