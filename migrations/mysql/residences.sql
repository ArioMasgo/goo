CREATE TABLE residencias
(
    id          CHAR(36)                                                                   NOT NULL PRIMARY KEY,
    nombre      VARCHAR(255)                                                               NOT NULL UNIQUE,
    genero      ENUM ('femenino', 'masculino')                                             NOT NULL,
    description TEXT,
    direccion   VARCHAR(255),
    estado      ENUM ('mantenimiento', 'deshabilitado', 'habilitado') DEFAULT 'habilitado' NOT NULL,
    created_by  int,
    created_at  TIMESTAMP                                             DEFAULT CURRENT_TIMESTAMP,
    updated_by  int,
    updated_at  TIMESTAMP                                             DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE cuartos
(
    id            CHAR(36)                                                                   NOT NULL PRIMARY KEY,
    numero        int                                                                        NOT NULL,
    capacidad     INT                                                                        NOT NULL,
    estado        ENUM ('mantenimiento', 'deshabilitado', 'habilitado') DEFAULT 'habilitado' NOT NULL,
    piso          INT                                                                        NOT NULL,
    residencia_id CHAR(36)                                                                   NOT NULL,
    created_by    int,
    created_at    TIMESTAMP                                             DEFAULT CURRENT_TIMESTAMP,
    updated_by    int,
    updated_at    TIMESTAMP                                             DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (residencia_id) REFERENCES residencias (id) ON DELETE CASCADE
);

CREATE TABLE configuracion_residencias
(
    id                     CHAR(36) NOT NULL PRIMARY KEY,
    porcentaje_fcea        FLOAT    NOT NULL,
    porcentaje_ingenieria  FLOAT    NOT NULL,
    nota_minima_fcea       FLOAT    NOT NULL,
    nota_minima_ingenieria FLOAT    NOT NULL,
    residencia_id          CHAR(36) NOT NULL,
    es_cachimbo            BOOLEAN  NOT NULL DEFAULT FALSE,
    created_by             int,
    created_at             TIMESTAMP         DEFAULT CURRENT_TIMESTAMP,
    updated_by             int,
    updated_at             TIMESTAMP         DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (residencia_id) REFERENCES residencias (id) ON DELETE CASCADE
);

CREATE TABLE `asignacion_cuartos`
(
    `id`               varchar(36)                                              NOT NULL,
    `alumno_id`        bigint(20) unsigned                                      NOT NULL,
    `cuarto_id`        varchar(36)                                              NOT NULL,
    `convocatoria_id`  bigint(20) unsigned                                      NOT NULL,
    `fecha_asignacion` datetime                                                 NOT NULL,
    `estado`           ENUM ('activo', 'desocupado', 'suspendido', 'cancelado') NOT NULL DEFAULT 'activo',
    `observaciones`    text,
    `created_at`       timestamp                                                NULL     DEFAULT NULL,
    `updated_at`       timestamp                                                NULL     DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_asignacion_alumno` (`alumno_id`),
    KEY `idx_asignacion_cuarto` (`cuarto_id`),
    KEY `idx_asignacion_convocatoria` (`convocatoria_id`),
    UNIQUE KEY `uk_alumno_convocatoria` (`alumno_id`, `convocatoria_id`),
    CONSTRAINT `fk_asignacion_alumno` FOREIGN KEY (`alumno_id`) REFERENCES `alumnos` (`id`),
    CONSTRAINT `fk_asignacion_cuarto` FOREIGN KEY (`cuarto_id`) REFERENCES `cuartos` (`id`),
    CONSTRAINT `fk_asignacion_convocatoria` FOREIGN KEY (`convocatoria_id`) REFERENCES `convocatorias` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
  ROW_FORMAT = DYNAMIC;

-- u103412080_obu_db.asignacion_cuartos definition

CREATE TABLE `asignacion_cuartos` (
                                      `id` varchar(36) NOT NULL,
                                      `alumno_id` bigint(20) unsigned NOT NULL,
                                      `cuarto_id` varchar(36) NOT NULL,
                                      `convocatoria_id` bigint(20) unsigned NOT NULL,
                                      `fecha_asignacion` datetime NOT NULL,
                                      `estado` varchar(50) NOT NULL DEFAULT 'activo',
                                      `observaciones` text DEFAULT NULL,
                                      `created_at` timestamp NULL DEFAULT NULL,
                                      `updated_at` timestamp NULL DEFAULT NULL,
                                      PRIMARY KEY (`id`),
                                      UNIQUE KEY `uk_alumno_convocatoria` (`alumno_id`,`convocatoria_id`),
                                      KEY `idx_asignacion_alumno` (`alumno_id`),
                                      KEY `idx_asignacion_cuarto` (`cuarto_id`),
                                      KEY `idx_asignacion_convocatoria` (`convocatoria_id`),
                                      CONSTRAINT `fk_asignacion_alumno` FOREIGN KEY (`alumno_id`) REFERENCES `alumnos` (`id`),
                                      CONSTRAINT `fk_asignacion_convocatoria` FOREIGN KEY (`convocatoria_id`) REFERENCES `convocatorias` (`id`),
                                      CONSTRAINT `fk_asignacion_cuarto` FOREIGN KEY (`cuarto_id`) REFERENCES `cuartos` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;