import java.util.ArrayList;
import java.util.List;

class Agenda {
    private static Agenda instancia;
    private List<Object> elementos;

    private Agenda() {
        this.elementos = new ArrayList<>();
    }

    public static Agenda getInstancia() {
        if (instancia == null) {
            instancia = new Agenda();
        }
        return instancia;
    }

    public void agregarElemento(Object elemento) {
        elementos.add(elemento);
    }

    public void eliminarElemento(Object elemento) {
        elementos.remove(elemento);
    }

    public void modificarElemento(int index, Object nuevoElemento) {
        if (index >= 0 && index < elementos.size()) {
            elementos.set(index, nuevoElemento);
        }
    }

    @Override
    public String toString() {
        return "Agenda{" +
                "elementos=" + elementos +
                '}';
    }

    public List<Object> getElementos() {
        return elementos;
    }
}
