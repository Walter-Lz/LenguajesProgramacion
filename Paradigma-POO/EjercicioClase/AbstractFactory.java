interface AbstractFactory {
    Contacto crearContacto(Persona persona, String telefono);
    Evento crearEvento(String titulo, String fecha, String lugar);
}

class ContactoFamiliarFactory implements AbstractFactory {
    private String relacion;

    public ContactoFamiliarFactory(String relacion) {
        this.relacion = relacion;
    }

    @Override
    public Contacto crearContacto(Persona persona, String telefono) {
        return new ContactoFamiliar(persona, telefono, relacion);
    }

    @Override
    public Evento crearEvento(String titulo, String fecha, String lugar) {
        throw new UnsupportedOperationException("No está permitido para crear eventos.");
    }
}


class ContactoEmpresarialFactory implements AbstractFactory {
    private String empresa;
    private String puesto;

    public ContactoEmpresarialFactory(String empresa, String puesto) {
        this.empresa = empresa;
        this.puesto = puesto;
    }

    @Override
    public Contacto crearContacto(Persona persona, String telefono) {
        return new ContactoEmpresarial(persona, telefono, empresa, puesto);
    }
    @Override
    public Evento crearEvento(String titulo, String fecha, String lugar) {
        throw new UnsupportedOperationException("No está permitido para crear eventos.");
    }
}
class CelebracionFactory implements AbstractFactory {
    private String tipo;

    public CelebracionFactory(String tipo) {
        this.tipo = tipo;
    }

    @Override
    public Contacto crearContacto(Persona persona, String telefono) {
        throw new UnsupportedOperationException("No está permitido para crear contactos.");
    }

    @Override
    public Evento crearEvento(String titulo, String fecha, String lugar) {
        return new Celebracion(titulo, fecha, lugar, tipo);
    }
}
