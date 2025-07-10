CREATE TABLE response_messages
(
    id           INT PRIMARY KEY AUTO_INCREMENT,
    code         INT                                  NOT NULL UNIQUE,
    message      VARCHAR(255)                         NOT NULL,
    message_type ENUM ('ERROR', 'WARNING', 'SUCCESS') NOT NULL,
    http_status  INT                                  NOT NULL,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Insertar mensajes de error del sistema (1-50)
INSERT INTO response_messages (code, message, message_type, http_status)
VALUES (1, 'Error al procesar el cuerpo de la solicitud', 'ERROR', 400),
       (2, 'Error de validación de datos', 'ERROR', 400),
       (3, 'Error en la base de datos', 'ERROR', 500),
       (4, 'Recurso no encontrado', 'ERROR', 404),
       (5, 'Error interno del servidor', 'ERROR', 500),
       (6, 'Método no permitido', 'ERROR', 405),
       (7, 'Cabeceras incompletas', 'ERROR', 400),
       (8, 'Formato de datos inválido', 'ERROR', 400),
       (9, 'Usuario no autenticado', 'ERROR', 401),
       (10, 'Permisos insuficientes', 'ERROR', 403),
       (11, 'No se pudo cargar el mensaje', 'ERROR', 500),
       (12, 'Error al crear el registro', 'ERROR', 500),
       (13, 'Error al actualizar el registro', 'ERROR', 500),
       (14, 'Error al eliminar el registro', 'ERROR', 500),
       (15, 'Error al obtener el registro', 'ERROR', 500),
       (16, 'Error al listar los registros', 'ERROR', 500);

INSERT INTO response_messages (code, message, message_type, http_status)
VALUES
    (17, 'Error al crear el archivo Excel', 'ERROR', 500),
    (18, 'Error al guardar el archivo Excel', 'ERROR', 500),
    (19, 'Error al cerrar el archivo Excel', 'ERROR', 500),
    (20, 'Error al eliminar la hoja por defecto', 'ERROR', 500),
    (21, 'Nombre de archivo Excel inválido', 'ERROR', 400),
    (22, 'Ruta de archivo Excel inválida', 'ERROR', 400),
    (23, 'El archivo Excel debe tener al menos una página', 'ERROR', 400),
    (24, 'Nombre de hoja Excel inválido', 'ERROR', 400),
    (25, 'Error al procesar las filas del Excel', 'ERROR', 500),
    (26, 'Error al procesar las columnas del Excel', 'ERROR', 500),
    (27, 'Error al establecer el valor de la celda', 'ERROR', 500),
    (28, 'Referencia de celda Excel inválida', 'ERROR', 400),
    (29, 'Error al aplicar estilos al Excel', 'ERROR', 500),
    (30, 'Error al aplicar formato a las celdas', 'ERROR', 500),
    (31, 'Error al ajustar el ancho de las columnas', 'ERROR', 500),
    (32, 'Error al establecer el zoom de la hoja', 'ERROR', 500),
    (33, 'Error al crear el directorio para el archivo Excel', 'ERROR', 500),
    (34, 'Error de permisos al escribir el archivo Excel', 'ERROR', 500),
    (35, 'Error de espacio en disco al guardar el Excel', 'ERROR', 500),
    (36, 'La ruta del archivo Excel no existe', 'ERROR', 404),
    (37, 'Memoria insuficiente para procesar el Excel', 'ERROR', 500),
    (38, 'Límite de tamaño de archivo Excel excedido', 'ERROR', 400),
    (39, 'Demasiadas hojas en el archivo Excel', 'ERROR', 400),
    (40, 'Demasiadas filas en la hoja Excel', 'ERROR', 400);

-- Insertar errores específicos de autenticación (51-70)
INSERT INTO response_messages (code, message, message_type, http_status)
VALUES (51, 'Credenciales inválidas', 'ERROR', 401),
       (52, 'Token expirado', 'ERROR', 401),
       (53, 'Token inválido', 'ERROR', 401),
       (54, 'Error al generar token', 'ERROR', 500),
       (55, 'Usuario bloqueado', 'ERROR', 401),
       (56, 'Sesión inválida', 'ERROR', 401),
       (57, 'IP no autorizada', 'ERROR', 403),
       (58, 'Error en refresh token', 'ERROR', 400),
       (91, 'Error al ejecutar la asignación automática de cuartos', 'ERROR', 500),
       (92, 'No se generó ninguna asignación de cuartos', 'ERROR', 500),
       (93, 'Habitación sin capacidad disponible', 'ERROR', 400),
       (94, 'El estudiante ya está asignado a esta habitación', 'ERROR', 400),
       (95, 'El estudiante ya tiene una asignación en otra habitación', 'ERROR', 400),
       (96, 'El estudiante no está aceptado en esta convocatoria', 'ERROR', 400);

-- Mensajes de éxito generales (210-219)
INSERT INTO response_messages (code, message, message_type, http_status)
VALUES (206, 'Asignación de cuartos realizada exitosamente', 'SUCCESS', 200),
       (210, 'Operación realizada exitosamente', 'SUCCESS', 200),
       (211, 'Registro creado exitosamente', 'SUCCESS', 201),
       (212, 'Registro actualizado exitosamente', 'SUCCESS', 200),
       (213, 'Registro eliminado exitosamente', 'SUCCESS', 200),
       (214, 'Registro obtenido exitosamente', 'SUCCESS', 200),
       (215, 'Registros listados exitosamente', 'SUCCESS', 200);

-- Mensajes de éxito específicos de autenticación (220-229)
INSERT INTO response_messages (code, message, message_type, http_status)
VALUES (220, 'Sesión creada exitosamente', 'SUCCESS', 200),
       (221, 'Token refrescado exitosamente', 'SUCCESS', 200),
       (222, 'Cierre de sesión exitoso', 'SUCCESS', 200),
       (223, 'Archivo generado exitosamente', 'SUCCESS', 200);