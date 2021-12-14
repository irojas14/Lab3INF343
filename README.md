# Sistemas Distribuídos - Tarea 3

## Integrantes:
* Cristóbal Abarca, 201573060-1
* Maximiliano Bombin, 201104308-1
* Ignacio Rojas, 201403027-4

---------------------------------------------------
# Instrucciones

## Iniciar Procesos:
1. En cada máquina iniciando sesión como el usuario 'alumno', deben ingresar a la carpeta llamada 'Lab3INF343' de la siguiente forma:

    cd Lab3INF343


2. Luego, en esta carpeta se pueden ejecutar los comandos correspondientes a cada máquina. Iniciar los procesos en su máquina correspondiente:
    * Máquina dis149 - Fulcrum1
        
            make fulcrum1

    * Máquina dis149 - Leia
        
            make leia
        
   * Máquina dis150 - MosEisley
        
            make moseisley

    * Máquina dis151 - Fulcrum2
        
            make fulcrum2

    * Máquina dis151 - Informantes
        
            make informante

    * Máquina dis152 - Fulcrum3
        
            make fulcrum3
        

3. Con make clean se limpian los .txt en caso de querer realizar otra ejecución limpiando todo en la máquina donde se ejecuta el clean.
---------------------------------------------------
## Jugar:

1. Inicializar MosEisley, Fulcrum 1, Fulcrum 2 y Fulcrum 3 (no importa el orden de inicialización).
2. Inicializar informante y Leia (sin importar el orden).
3. Desde la consola de Informante se puede hacer lo siguiente:
    3.1. A + enter = Agregar una ciudad -> Ingresar nombre de planeta, nombre de ciudad y finalmente ingresar el número de rebeldes.
    3.2. N + enter = Actualizar nombre de ciudad -> Ingresar nombre de planeta, nombre de ciudad, ingresar nuevo nombre de ciudad.
    3.3. F + enter = Actualizar número de rebeldes en ciudad -> Ingresar nombre de planeta, nombre de ciudad, ingresar nuevo número.
    3.4. D + enter = Borrar ciudad -> Ingresar nombre del planeta, ingresar nombre de ciudad.
    3.5. P + enter = Preguntar por número de rebeldes para confirmar consistencia read-your-writes -> Ingresar nombre de planeta, ingresar nombre de ciudad.
    3.6. E + enter = Salir.
4. Desde la consola de Leia se puede hacer lo siguiente:
    4.1. A + enter = Obtener número de rebeldes (se ocupa monotonic-reads) -> Ingresar nombre de planeta, ingresar nombre de ciudad.
    4.2. E + enter = Salir.
5. Si se quiere detener a Fulcrum y MosEisley se debe hacer de forma manual.