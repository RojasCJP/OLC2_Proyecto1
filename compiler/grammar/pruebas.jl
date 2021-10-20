function factorial(numero::Int64)::Int64
    if numero != 0
        numero2 = factorial(numero-1);
        return numero*numero2;
    end;
    return 1;
end;
println(factorial(5));