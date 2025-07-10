INSERT INTO solicitudes (alumno_id, convocatoria_id, created_at, updated_at)
SELECT
    id as alumno_id,
    98 as convocatoria_id,
    NOW(),
    NOW()
FROM alumnos, (SELECT @row_number:=0) AS t
WHERE id <= 300;

INSERT INTO detalle_solicitudes (solicitud_id, requisito_id, respuesta_formulario, url_documento, opcion_seleccion, created_at, updated_at, `order`)
SELECT
    s.id AS solicitud_id,
    r.id AS requisito_id,
    CASE
        WHEN r.type_input = 'text' THEN CONCAT('Respuesta ', FLOOR(RAND() * 100))
        WHEN r.type_input = 'number' THEN FLOOR(RAND() * 1000)
        WHEN r.type_input = 'date' THEN DATE_FORMAT(NOW() - INTERVAL FLOOR(RAND() * 365) DAY, '%Y-%m-%d')
        ELSE NULL
        END AS respuesta_formulario,
    CASE
        WHEN r.tipo_requisito_id = 1 THEN CONCAT('url_doc_', FLOOR(RAND() * 1000), '.pdf')
        WHEN r.tipo_requisito_id = 2 THEN CONCAT('url_image_', FLOOR(RAND() * 1000), '.jpg')
        ELSE NULL
        END AS url_documento,
    CASE
        WHEN r.type_input = 'text' AND r.opciones != '' THEN
            SUBSTRING_INDEX(SUBSTRING_INDEX(r.opciones, '|', FLOOR(1 + RAND() * (LENGTH(r.opciones) - LENGTH(REPLACE(r.opciones, '|', '')) + 1))), '|', -1)
        ELSE NULL
        END AS opcion_seleccion,
    NOW() AS created_at,
    NOW() AS updated_at,
    1 AS `order`
FROM solicitudes s
         CROSS JOIN (
    SELECT * FROM requisitos
    WHERE id BETWEEN 572 AND 580  -- Tomamos solo algunos requisitos para el ejemplo
) r
WHERE s.convocatoria_id = 98;

INSERT INTO servicio_solicitado (estado, servicio_id, solicitud_id, fecha_revision, created_at, updated_at)
SELECT
    CASE WHEN RAND() < 0.7 THEN 'aprobado' ELSE 'desaprobado' END as estado,
    1 as servicio_id,  -- Asumiendo que el servicio_id 1 es para residencia
    s.id as solicitud_id,
    CURDATE() as fecha_revision,
    NOW() as created_at,
    NOW() as updated_at
FROM solicitudes s
WHERE s.convocatoria_id = 98;