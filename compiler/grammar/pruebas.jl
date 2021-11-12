struct Persona
    nombre::String;
    edad::Int64;
    numeroFamiliares::Int64;
end;

   
function AgregarFamiliar(persona::Persona)
    persona.numeroFamiliares = persona.numeroFamiliares + 1;
end;

function ImprimirDatosPersona(persona::Persona)
    println("Nombre: ", persona.nombre);
    println("Edad: ", persona.edad);
    println("Numero de familiares: ", persona.numeroFamiliares);
end;

manuel = Persona("Manuel", 22, 4);
ImprimirDatosPersona(manuel);
AgregarFamiliar(manuel);
ImprimirDatosPersona(manuel);