class Evento {
    private String titulo;
    private String fecha;
    private String lugar;

    public Evento(String titulo, String fecha, String lugar) {
        this.titulo = titulo;
        this.fecha = fecha;
        this.lugar = lugar;
    }

    public String getTitulo() {
        return titulo;
    }

    public String getFecha() {
        return fecha;
    }

    public String getLugar() {
        return lugar;
    }

    @Override
    public String toString() {
        return "Evento{" +
                "titulo='" + titulo + '\'' +
                ", fecha='" + fecha + '\'' +
                ", lugar='" + lugar + '\'' +
                '}';
    }
}

class Reunion extends Evento {
    private String agenda;

    public Reunion(String titulo, String fecha, String lugar, String agenda) {
        super(titulo, fecha, lugar);
        this.agenda = agenda;
    }

    public String getAgenda() {
        return agenda;
    }

    @Override
    public String toString() {
        return "Reunion{" +
                "titulo='" + getTitulo() + '\'' +
                ", fecha='" + getFecha() + '\'' +
                ", lugar='" + getLugar() + '\'' +
                ", agenda='" + agenda + '\'' +
                '}';
    }
}

class Celebracion extends Evento {
    private String tipo;

    public Celebracion(String titulo, String fecha, String lugar, String tipo) {
        super(titulo, fecha, lugar);
        this.tipo = tipo;
    }

    public String getTipo() {
        return tipo;
    }

    @Override
    public String toString() {
        return "Celebracion{" +
                "titulo='" + getTitulo() + '\'' +
                ", fecha='" + getFecha() + '\'' +
                ", lugar='" + getLugar() + '\'' +
                ", tipo='" + tipo + '\'' +
                '}';
    }
}
