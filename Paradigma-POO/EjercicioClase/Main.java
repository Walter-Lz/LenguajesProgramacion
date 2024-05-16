import java.util.List;
import java.util.stream.Collectors;

public class Main {
    public static void main(String[] args) {
        
        Persona p1 = new Persona("12345678", "Juan", "Perez", "juan@ejemplo.com");
        Persona p2 = new Persona("87654321", "Ana", "Gomez", "ana@ejemplo.com");

        // Uso de las fábricas
        AbstractFactory familiarFactory = new ContactoFamiliarFactory("Hermana");
        AbstractFactory empresarialFactory = new ContactoEmpresarialFactory("TechCorp", "Ingeniero");
        AbstractFactory celebracionFactory = new CelebracionFactory("Cumpleaños de Juan");

        Contacto c1 = new Contacto(p1, "88445566");
        Contacto c2 = familiarFactory.crearContacto(p2, "987654321");
        Contacto c3 = empresarialFactory.crearContacto(p1, "123456789");
        //Contacto c4 = new Contacto(p3, "111222333"); // Este contacto no se creará debido a la cédula duplicada

        Evento e1 = new Evento("Reunion de Proyecto", "2024-05-20", "Oficina");
        Evento e2 = new Reunion("Reunion Ejecutiva", "2024-05-21", "Sala de Juntas", "Progreso del Q2");
        Evento e3 = celebracionFactory.crearEvento("Fiesta de Cumpleaños", "2024-05-22", "Casa");

        // Uso del Singleton para la Agenda
        Agenda agenda = Agenda.getInstancia();
        agenda.agregarElemento(c1);
        agenda.agregarElemento(e1);
        agenda.agregarElemento(c2);
        agenda.agregarElemento(e2);
        agenda.agregarElemento(c3);
        agenda.agregarElemento(e3);

        System.out.println("Agenda Completa:");
        System.out.println(agenda);
        System.out.println("------------------------------");
        List<Object> contactos = agenda.getElementos().stream()
                .filter(elemento -> elemento instanceof Contacto)
                .collect(Collectors.toList());
        List<Object> eventos = agenda.getElementos().stream()
                .filter(elemento -> elemento instanceof Evento)
                .collect(Collectors.toList());
                
        System.out.println("Contactos:");
        contactos.forEach(System.out::println);
        System.out.println("------------------------------");
        System.out.println("Eventos:");
        eventos.forEach(System.out::println);
        System.out.println("------------------------------");
    }
}
/* Uso de Singleton y Abstract Factory

La clase Agenda utiliza el patrón Singleton para asegurar que solo
haya una instancia de la agenda en toda la aplicación. El método getInstancia se encarga de 
retornar la única instancia de Agenda.

La interfaz AbstractFactory define métodos para crear contactos y eventos.
Las clases ContactoFamiliarFactory, ContactoEmpresarialFactory, y CelebracionFactory implementan la interfaz AbstractFactory,
permitiendo la creación de diferentes tipos de contactos y eventos sin especificar las clases concretas en el código cliente.

Definicion:
- Eager Singleton: La instancia se crea incluso si no se utiliza inmediatamente, lo que puede consumir recursos innecesarios.
- lazy Singleton: la instancia única de la clase no se crea hasta que se solicite explícitamente mediante un método de obtención.

# para este problema es mucho mas recomendable utilizar Lazy singleton ya que la instancia de la clase Agenda se va a necesitar una vez
# que existan eventos y personas a las cuales se requiera insertar en dicha agenda, por lo que implementar un Eager Singleton desde el principio 
# lo único que se logra con eso es consumir recursos innecesarios. 
*/ 

/*
        Ejecución del Código

        Agenda Completa:
        Agenda{elementos=[Contacto{persona=Persona{cedula='12345678', nombre='Juan', apellido='Perez', email='juan@ejemplo.com'}, telefono='88445566'},
        Evento{titulo='Reunion de Proyecto', fecha='2024-05-20', lugar='Oficina'},
        ContactoFamiliar{persona=Persona{cedula='87654321', nombre='Ana', apellido='Gomez', email='ana@ejemplo.com'}, telefono='987654321', relacion='Hermana'}, 
        Reunion{titulo='Reunion Ejecutiva', fecha='2024-05-21', lugar='Sala de Juntas', agenda='Progreso del Q2'}, 
        ContactoEmpresarial{persona=Persona{cedula='12345678', nombre='Juan', apellido='Perez', email='juan@ejemplo.com'}, 
        telefono='123456789', empresa='TechCorp', puesto='Ingeniero'}, Celebracion{titulo='Fiesta de Cumpleaños', fecha='2024-05-22', lugar='Casa', tipo='Cumpleaños de Juan'}]}
        ------------------------------
        Contactos:
        Contacto{persona=Persona{cedula='12345678', nombre='Juan', apellido='Perez', email='juan@ejemplo.com'}, telefono='88445566'}
        ContactoFamiliar{persona=Persona{cedula='87654321', nombre='Ana', apellido='Gomez', email='ana@ejemplo.com'}, telefono='987654321', relacion='Hermana'}
        ContactoEmpresarial{persona=Persona{cedula='12345678', nombre='Juan', apellido='Perez', email='juan@ejemplo.com'}, telefono='123456789', empresa='TechCorp', puesto='Ingeniero'}
        ------------------------------
        Eventos:
        Evento{titulo='Reunion de Proyecto', fecha='2024-05-20', lugar='Oficina'}
        Reunion{titulo='Reunion Ejecutiva', fecha='2024-05-21', lugar='Sala de Juntas', agenda='Progreso del Q2'}
        Celebracion{titulo='Fiesta de Cumpleaños', fecha='2024-05-22', lugar='Casa', tipo='Cumpleaños de Juan'}
        ------------------------------
 */
