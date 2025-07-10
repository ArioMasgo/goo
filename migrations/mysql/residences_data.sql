-- Primero insertamos las residencias para mujeres y guardamos sus IDs en variables
SET @pomarrosas_id = UUID();
SET @tulumayo_id = UUID();
SET @mariaangola_id = UUID();
SET @damas_id = UUID();

INSERT INTO residencias (id, nombre, genero, description, direccion, estado, created_by, updated_by)
VALUES (@pomarrosas_id, 'POMARROSAS', 'femenino',
        'Residencia universitaria femenina con amplias áreas comunes, sala de estudio y lavandería. Ubicada en zona tranquila y segura.',
        'Av. Los Girasoles 123, Huancayo',
        'habilitado', 1, 1),
       (@tulumayo_id, 'TULUMAYO', 'femenino',
        'Moderna residencia para estudiantes con internet de alta velocidad, áreas verdes y sistema de seguridad 24/7.',
        'Jr. Tulumayo 456, Huancayo',
        'habilitado', 1, 1),
       (@mariaangola_id, 'MARIA ANGOLA', 'femenino',
        'Residencia estudiantil con espacios de recreación, comedor común y salas de estudio grupales. Excelente ubicación cerca al campus.',
        'Calle María Angola 789, Huancayo',
        'habilitado', 1, 1),
       (@damas_id, 'DAMAS', 'femenino',
        'Residencia con ambiente familiar, áreas comunes acogedoras y servicios básicos incluidos. Perfecta para estudiantes de primer año.',
        'Jr. Las Damas 234, Huancayo',
        'habilitado', 1, 1);

-- Insertamos las residencias para varones
SET @sheraton_id = UUID();
SET @bambu_id = UUID();
SET @britanico_id = UUID();
SET @castano_id = UUID();

INSERT INTO residencias (id, nombre, genero, description, direccion, estado, created_by, updated_by)
VALUES (@sheraton_id, 'SHERATON', 'masculino',
        'Residencia moderna con gimnasio, sala de juegos y espacios de estudio. Ubicación estratégica cerca de la biblioteca central.',
        'Av. San Carlos 567, Huancayo',
        'habilitado', 1, 1),
       (@bambu_id, 'BAMBU', 'masculino',
        'La residencia más grande del campus, cuenta con áreas deportivas, salas de estudio y terraza común. Ambiente ideal para el desarrollo académico.',
        'Jr. Los Bambúes 890, Huancayo',
        'habilitado', 1, 1),
       (@britanico_id, 'BRITANICO', 'masculino',
        'Residencia de estilo clásico con servicios modernos, sala de cómputo y áreas de descanso. Perfecta para estudiantes internacionales.',
        'Calle Inglaterra 123, Huancayo',
        'habilitado', 1, 1),
       (@castano_id, 'CASTAÑO', 'masculino',
        'Residencia estudiantil con ambiente acogedor, áreas verdes y espacios de estudio tranquilos. Excelente para la concentración académica.',
        'Jr. Los Castaños 345, Huancayo',
        'habilitado', 1, 1);

-- Configuración para residencias femeninas
INSERT INTO configuracion_residencias (id,
                                       porcentaje_fcea,
                                       porcentaje_ingenieria,
                                       nota_minima_fcea,
                                       nota_minima_ingenieria,
                                       residencia_id,
                                       created_by,
                                       updated_by)
VALUES
-- POMARROSAS
(UUID(), 60, 40, 14, 13, @pomarrosas_id, 1, 1),

-- TULUMAYO
(UUID(), 50, 50, 14, 13, @tulumayo_id, 1, 1),

-- MARIA ANGOLA
(UUID(), 45, 55, 13.5, 13, @mariaangola_id, 1, 1),

-- DAMAS
(UUID(), 55, 45, 14, 13.5, @damas_id, 1, 1),

-- SHERATON
(UUID(), 40, 60, 13.5, 13, @sheraton_id, 1, 1),

-- BAMBU
(UUID(), 35, 65, 13.5, 13, @bambu_id, 1, 1),

-- BRITANICO
(UUID(), 45, 55, 14, 13.5, @britanico_id, 1, 1),

-- CASTAÑO
(UUID(), 50, 50, 14, 13, @castano_id, 1, 1);

-- POMARROSAS: 20 cuartos, 2 estudiantes por cuarto
INSERT INTO cuartos (id, numero, capacidad, estado, piso, residencia_id, created_by, updated_by)
SELECT UUID(),
       numero,
       2,
       'habilitado',
       CEILING(numero / 10),
       @pomarrosas_id,
       1,
       1
FROM (SELECT @row := @row + 1 AS numero
      FROM (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4) t1,
           (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4) t2,
           (SELECT @row := 0) r
      LIMIT 20) numbers;

-- TULUMAYO: 20 cuartos, 3 estudiantes por cuarto
INSERT INTO cuartos (id, numero, capacidad, estado, piso, residencia_id, created_by, updated_by)
SELECT UUID(),
       numero,
       3,
       'habilitado',
       CEILING(numero / 10),
       @tulumayo_id,
       1,
       1
FROM (SELECT @row := @row + 1 AS numero
      FROM (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4) t1,
           (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4) t2,
           (SELECT @row := 0) r
      LIMIT 20) numbers;

-- MARIA ANGOLA: 24 cuartos, 3 estudiantes por cuarto
INSERT INTO cuartos (id, numero, capacidad, estado, piso, residencia_id, created_by, updated_by)
SELECT UUID(),
       numero,
       3,
       'habilitado',
       CEILING(numero / 10),
       @mariaangola_id,
       1,
       1
FROM (SELECT @row := @row + 1 AS numero
      FROM (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4 UNION ALL SELECT 5) t1,
           (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4) t2,
           (SELECT @row := 0) r
      LIMIT 24) numbers;

-- DAMAS: 20 cuartos, 2 estudiantes por cuarto
INSERT INTO cuartos (id, numero, capacidad, estado, piso, residencia_id, created_by, updated_by)
SELECT UUID(),
       numero,
       2,
       'habilitado',
       CEILING(numero / 10),
       @damas_id,
       1,
       1
FROM (SELECT @row := @row + 1 AS numero
      FROM (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4) t1,
           (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4) t2,
           (SELECT @row := 0) r
      LIMIT 20) numbers;

-- SHERATON: 30 cuartos, 2 estudiantes por cuarto
INSERT INTO cuartos (id, numero, capacidad, estado, piso, residencia_id, created_by, updated_by)
SELECT UUID(),
       numero,
       2,
       'habilitado',
       CEILING(numero / 10),
       @sheraton_id,
       1,
       1
FROM (SELECT @row := @row + 1 AS numero
      FROM (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4) t1,
           (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4) t2,
           (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2) t3,
           (SELECT @row := 0) r
      LIMIT 30) numbers;

-- BAMBU: 54 cuartos, 3 estudiantes por cuarto
INSERT INTO cuartos (id, numero, capacidad, estado, piso, residencia_id, created_by, updated_by)
SELECT UUID(),
       numero,
       3,
       'habilitado',
       CEILING(numero / 10),
       @bambu_id,
       1,
       1
FROM (SELECT @row := @row + 1 AS numero
      FROM (SELECT 0
            UNION ALL
            SELECT 1
            UNION ALL
            SELECT 2
            UNION ALL
            SELECT 3
            UNION ALL
            SELECT 4
            UNION ALL
            SELECT 5
            UNION ALL
            SELECT 6
            UNION ALL
            SELECT 7
            UNION ALL
            SELECT 8) t1,
           (SELECT 0
            UNION ALL
            SELECT 1
            UNION ALL
            SELECT 2
            UNION ALL
            SELECT 3
            UNION ALL
            SELECT 4
            UNION ALL
            SELECT 5
            UNION ALL
            SELECT 6) t2,
           (SELECT @row := 0) r
      LIMIT 54) numbers;

-- BRITANICO: 20 cuartos, 2 estudiantes por cuarto
INSERT INTO cuartos (id, numero, capacidad, estado, piso, residencia_id, created_by, updated_by)
SELECT UUID(),
       numero,
       2,
       'habilitado',
       CEILING(numero / 10),
       @britanico_id,
       1,
       1
FROM (SELECT @row := @row + 1 AS numero
      FROM (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4) t1,
           (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4) t2,
           (SELECT @row := 0) r
      LIMIT 20) numbers;

-- CASTAÑO: 22 cuartos, 3 estudiantes por cuarto
INSERT INTO cuartos (id, numero, capacidad, estado, piso, residencia_id, created_by, updated_by)
SELECT UUID(),
       numero,
       3,
       'habilitado',
       CEILING(numero / 10),
       @castano_id,
       1,
       1
FROM (SELECT @row := @row + 1 AS numero
      FROM (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4) t1,
           (SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4 UNION ALL SELECT 5) t2,
           (SELECT @row := 0) r
      LIMIT 22) numbers;