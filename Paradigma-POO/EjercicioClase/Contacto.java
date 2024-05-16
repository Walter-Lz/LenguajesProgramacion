class Contacto {
    protected Persona persona;
    protected String telefono;

    public Contacto(Persona persona, String telefono) {
        this.persona = persona;
        this.telefono = telefono;
    }

    public Persona getPersona() {
        return persona;
    }

    public String getTelefono() {
        return telefono;
    }

    @Override
    public String toString() {
        return "Contacto{" +
                "persona=" + persona +
                ", telefono='" + telefono + '\'' +
                '}';
    }
}

class ContactoFamiliar extends Contacto {
    private String relacion;

    public ContactoFamiliar(Persona persona, String telefono, String relacion) {
        super(persona, telefono);
        this.relacion = relacion;
    }

    public String getRelacion() {
        return relacion;
    }

    @Override
    public String toString() {
        return "ContactoFamiliar{" +
                "persona=" + persona +
                ", telefono='" + telefono + '\'' +
                ", relacion='" + relacion + '\'' +
                '}';
    }
}
class ContactoEmpresarial extends Contacto {
    private String empresa;
    private String puesto;

    public ContactoEmpresarial(Persona persona, String telefono, String empresa, String puesto) {
        super(persona, telefono);
        this.empresa = empresa;
        this.puesto = puesto;
    }

    public String getEmpresa() {
        return empresa;
    }

    public String getPuesto() {
        return puesto;
    }

    @Override
    public String toString() {
        return "ContactoEmpresarial{" +
                "persona=" + persona +
                ", telefono='" + telefono + '\'' +
                ", empresa='" + empresa + '\'' +
                ", puesto='" + puesto + '\'' +
                '}';
    }
}