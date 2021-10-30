struct Prueba 
    nombre::String;
    edad::Int64;
end;

prueba = Prueba("hola buenas",412);
println(prueba.edad);
println(prueba.nombre);