struct Actor
    nombre;
    edad;
end;

struct Pelicula
    nombre;
    posicion;
end;

struct Contrato
    actor;
    pelicula;
end;

actores = ["Elizabeth Olsen", "Adam Sandler", "Christian Bale", "Jennifer Aniston"];
peliculas = ["Avengers: Age of Ultron", "Mr. Deeds", "Batman: The Dark Knight", "Marley & Me"];

function contratar(actor, pelicula)
    contrato:: Contrato;
    contrato.actor=actor;
    contrato.pelicula=pelicula;
    return contrato;
end;

function crearActor(nombre, edad)
    actor::Actor;
    actor.nombre = nombre;
    actor.edad = edad;
    return actor;
end;

function crearPelicula(nombre, posicion)
    pelicula::Pelicula;
    pelicula.nombre = nombre;
    pelicula.posicion = posicion;
    return pelicula;
end;

function imprimir(contrato)
    println(contrato);
end;

function contratos()
    for i in 1:(1*1+2)
        contrato::Contrato;
        if(i < 4)
            actor = crearActor(actores[i], i+38);
            pelicula = crearPelicula(peliculas[i], i);
            contrato = contratar(actor, pelicula);
        end;
        imprimir(contrato);
    end;
end;

contratos();
