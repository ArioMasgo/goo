-- u103412080_obu_db.pacientes definition

CREATE TABLE `pacientes` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `is_deleted` tinyint(1) DEFAULT 0,
  `user_deleted` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_creator` varchar(255) DEFAULT NULL,
  `tipo_persona` varchar(255) DEFAULT NULL,
  `codigo_sga` varchar(255) DEFAULT NULL,
  `dni` varchar(20) NOT NULL,
  `nombres` varchar(255) NOT NULL,
  `apellidos` varchar(255) NOT NULL,
  `sexo` varchar(10) DEFAULT NULL,
  `edad` varchar(3) DEFAULT NULL,
  `estado_civil` varchar(20) DEFAULT NULL,
  `grupo_sanguineo` varchar(5) DEFAULT NULL,
  `fecha_nacimiento` varchar(30) DEFAULT NULL,
  `lugar_nacimiento` varchar(255) DEFAULT NULL,
  `procedencia` varchar(255) DEFAULT NULL,
  `escuela_profesional` varchar(255) DEFAULT NULL,
  `ocupacion` varchar(255) DEFAULT NULL,
  `correo_electronico` varchar(255) DEFAULT NULL,
  `numero_celular` varchar(20) NOT NULL,
  `direccion` varchar(100) DEFAULT NULL,
  `factor_rh` varchar(255) DEFAULT NULL,
  `alergias` varchar(100) DEFAULT NULL,
  `ram` tinyint(1) DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- u103412080_obu_db.paciente_antecedentes definition

CREATE TABLE `paciente_antecedentes` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `is_deleted` tinyint(1) DEFAULT 0,
  `user_deleted` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_creator` varchar(255) DEFAULT NULL,
  `paciente_id` char(36) NOT NULL,
  `nombre_antecedente` varchar(100) NOT NULL,
  `estado_antecedente` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_paciente_antecedentes` (`paciente_id`),
  CONSTRAINT `fk_paciente_antecedentes` FOREIGN KEY (`paciente_id`) REFERENCES `pacientes` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- u103412080_obu_db.consultas_areas_medicas definition

CREATE TABLE `consultas_areas_medicas` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `is_deleted` tinyint(1) DEFAULT 0,
  `user_deleted` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_creator` varchar(255) DEFAULT NULL,
  `paciente_id` char(36) NOT NULL,
  `fecha_consulta` varchar(100) NOT NULL,
  `area_medica` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_consultas_areas_medicas` (`paciente_id`),
  CONSTRAINT `fk_consultas_areas_medicas` FOREIGN KEY (`paciente_id`) REFERENCES `pacientes` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- u103412080_obu_db.area_medica_asignada definition

CREATE TABLE `area_medica_asignada` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `is_deleted` tinyint(1) DEFAULT 0,
  `user_deleted` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_creator` varchar(255) DEFAULT NULL,
  `consulta_id` char(36) NOT NULL,
  `area_asignada` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_enfermeria_area_asignada` (`consulta_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- u103412080_obu_db.firmas_convocatoria_area_medica definition

CREATE TABLE `firmas_convocatoria_area_medica` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `is_deleted` tinyint(1) DEFAULT 0,
  `user_deleted` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_creator` varchar(255) DEFAULT NULL,
  `convocatoria_id` varchar(10) NOT NULL,
  `paciente_id` char(36) NOT NULL,
  `firma_enfermeria` varchar(255) DEFAULT NULL,
  `firma_medicina` varchar(255) DEFAULT NULL,
  `firma_odontologia` varchar(255) DEFAULT NULL,
  `firma_psicologia` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- u103412080_obu_db.enfermeria_datos_acompanante definition

CREATE TABLE `enfermeria_datos_acompanante` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `is_deleted` tinyint(1) DEFAULT 0,
  `user_deleted` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_creator` varchar(255) DEFAULT NULL,
  `consulta_enfermeria_id` char(36) NOT NULL,
  `dni` varchar(255) DEFAULT NULL,
  `nombres_apellidos` varchar(255) DEFAULT NULL,
  `edad` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_enfermeria_datos_acompanante` (`consulta_enfermeria_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- u103412080_obu_db.enfermeria_examen_fisico definition

CREATE TABLE `enfermeria_examen_fisico` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `is_deleted` tinyint(1) DEFAULT 0,
  `user_deleted` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_creator` varchar(255) DEFAULT NULL,
  `consulta_enfermeria_id` char(36) NOT NULL,
  `talla_peso` varchar(255) DEFAULT NULL,
  `perimetro_cintura` varchar(255) DEFAULT NULL,
  `indice_masa_corporal_img` varchar(255) DEFAULT NULL,
  `presion_arterial` varchar(255) DEFAULT NULL,
  `comentarios` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_enfermeria_examen_fisico` (`consulta_enfermeria_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- u103412080_obu_db.enfermeria_examen_laboratorio definition

CREATE TABLE `enfermeria_examen_laboratorio` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `is_deleted` tinyint(1) DEFAULT 0,
  `user_deleted` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_creator` varchar(255) DEFAULT NULL,
  `consulta_enfermeria_id` char(36) NOT NULL,
  `serologia` varchar(255) DEFAULT NULL,
  `bk` varchar(255) DEFAULT NULL,
  `hemograma` varchar(255) DEFAULT NULL,
  `examen_orina` varchar(255) DEFAULT NULL,
  `colesterol` varchar(255) DEFAULT NULL,
  `glucosa` varchar(255) DEFAULT NULL,
  `comentarios` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_enfermeria_examen_laboratorio` (`consulta_enfermeria_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- u103412080_obu_db.enfermeria_examen_preferencial definition

CREATE TABLE `enfermeria_examen_preferencial` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `is_deleted` tinyint(1) DEFAULT 0,
  `user_deleted` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_creator` varchar(255) DEFAULT NULL,
  `consulta_enfermeria_id` char(36) NOT NULL,
  `aparato_respiratorio` varchar(255) DEFAULT NULL,
  `aparato_cardiovascular` varchar(255) DEFAULT NULL,
  `aparato_digestivo` varchar(255) DEFAULT NULL,
  `aparato_genitourinario` varchar(255) DEFAULT NULL,
  `papanicolau` varchar(255) DEFAULT NULL,
  `examen_mama` varchar(255) DEFAULT NULL,
  `comentarios` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_enfermeria_examen_preferencial` (`consulta_enfermeria_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- u103412080_obu_db.enfermeria_examen_sexualidad definition

CREATE TABLE `enfermeria_examen_sexualidad` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `is_deleted` tinyint(1) DEFAULT 0,
  `user_deleted` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_creator` varchar(255) DEFAULT NULL,
  `consulta_enfermeria_id` char(36) NOT NULL,
  `actividad_sexual` varchar(255) DEFAULT NULL,
  `planificacion_familiar` varchar(255) DEFAULT NULL,
  `comentarios` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_enfermeria_examen_sexualidad` (`consulta_enfermeria_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- u103412080_obu_db.enfermeria_examen_visual definition

CREATE TABLE `enfermeria_examen_visual` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `is_deleted` tinyint(1) DEFAULT 0,
  `user_deleted` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_creator` varchar(255) DEFAULT NULL,
  `consulta_enfermeria_id` char(36) NOT NULL,
  `ojo_derecho` varchar(255) DEFAULT NULL,
  `ojo_izquierdo` varchar(255) DEFAULT NULL,
  `presion_ocular` varchar(255) DEFAULT NULL,
  `comentarios` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_enfermeria_examen_visual` (`consulta_enfermeria_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- u103412080_obu_db.enfermeria_procedimientos_realizados definition

CREATE TABLE `enfermeria_procedimientos_realizados` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `is_deleted` tinyint(1) DEFAULT 0,
  `user_deleted` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_creator` varchar(255) DEFAULT NULL,
  `consulta_enfermeria_id` char(36) NOT NULL,
  `procedimiento` varchar(255) DEFAULT NULL,
  `numero_recibo` varchar(255) DEFAULT NULL,
  `comentarios` varchar(255) DEFAULT NULL,
  `costo` varchar(50) DEFAULT NULL,
  `fecha_pago` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_enfermeria_procedimientos_realizados` (`consulta_enfermeria_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- u103412080_obu_db.enfermeria_revision_rutina definition

CREATE TABLE `enfermeria_revision_rutina` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `is_deleted` tinyint(1) DEFAULT 0,
  `user_deleted` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_creator` varchar(255) DEFAULT NULL,
  `consulta_enfermeria_id` char(36) NOT NULL,
  `fiebre_ultimo_quince_dias` varchar(255) DEFAULT NULL,
  `tos_mas_quince_dias` varchar(255) DEFAULT NULL,
  `secrecion_lesion_genitales` varchar(255) DEFAULT NULL,
  `fecha_ultima_regla` varchar(50) DEFAULT NULL,
  `comentarios` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_enfermeria_revision_rutina` (`consulta_enfermeria_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- u103412080_obu_db.enfermeria_tratamiento_medicamentoso definition

CREATE TABLE `enfermeria_tratamiento_medicamentoso` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `is_deleted` tinyint(1) DEFAULT 0,
  `user_deleted` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_creator` varchar(255) DEFAULT NULL,
  `consulta_enfermeria_id` char(36) NOT NULL,
  `nombre_generico_medicamento` varchar(255) DEFAULT NULL,
  `via_administracion` varchar(255) DEFAULT NULL,
  `hora_aplicacion` varchar(255) DEFAULT NULL,
  `responsable_atencion` varchar(255) DEFAULT NULL,
  `observaciones` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_enfermeria_tratamiento_medicamentoso` (`consulta_enfermeria_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- u103412080_obu_db.enfermeria_vacuna definition

CREATE TABLE `enfermeria_vacuna` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `is_deleted` tinyint(1) DEFAULT 0,
  `user_deleted` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_creator` varchar(255) DEFAULT NULL,
  `consulta_enfermeria_id` char(36) NOT NULL,
  `tipo_vacuna` varchar(255) DEFAULT NULL,
  `fecha_dosis` varchar(255) DEFAULT NULL,
  `comentarios` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_enfermeria_vacuna` (`consulta_enfermeria_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- u103412080_obu_db.medicina_atencion_integral definition

CREATE TABLE `medicina_atencion_integral` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `is_deleted` tinyint(1) DEFAULT 0,
  `user_deleted` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_creator` varchar(255) DEFAULT NULL,
  `consulta_id` char(36) NOT NULL,
  `fecha` varchar(50) NOT NULL,
  `hora` varchar(20) DEFAULT NULL,
  `edad` varchar(5) DEFAULT NULL,
  `motivo_consulta` varchar(50) DEFAULT NULL,
  `tiempo_enfermedad` varchar(50) DEFAULT NULL,
  `apetito` varchar(50) DEFAULT NULL,
  `sed` varchar(50) DEFAULT NULL,
  `suenio` varchar(50) NOT NULL,
  `estado_animo` varchar(50) NOT NULL,
  `orina` varchar(50) DEFAULT NULL,
  `deposiciones` varchar(50) DEFAULT NULL,
  `temperatura` varchar(30) DEFAULT NULL,
  `presion_arterial` varchar(30) DEFAULT NULL,
  `frecuencia_cardiaca` varchar(30) DEFAULT NULL,
  `frecuencia_respiratoria` varchar(30) DEFAULT NULL,
  `peso` varchar(20) DEFAULT NULL,
  `talla` varchar(20) DEFAULT NULL,
  `indice_masa_corporal` varchar(30) DEFAULT NULL,
  `diagnostico` varchar(255) DEFAULT NULL,
  `tratamiento` varchar(255) DEFAULT NULL,
  `examenes_axuliares` varchar(255) DEFAULT NULL,
  `referencia` varchar(255) DEFAULT NULL,
  `observacion` varchar(255) DEFAULT NULL,
  `numero_recibo` varchar(50) DEFAULT NULL,
  `costo` varchar(500) DEFAULT NULL,
  `fecha_pago` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- u103412080_obu_db.medicina_consulta_general definition

CREATE TABLE `medicina_consulta_general` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `is_deleted` tinyint(1) DEFAULT 0,
  `user_deleted` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_creator` varchar(255) DEFAULT NULL,
  `consulta_id` char(36) NOT NULL,
  `fecha_hora` varchar(40) DEFAULT NULL,
  `anamnesis` varchar(500) DEFAULT NULL,
  `examen_clinico` varchar(500) DEFAULT NULL,
  `indicaciones` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_medicina_atencion_integral_otros` (`consulta_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- u103412080_obu_db.odontologia_consulta definition

CREATE TABLE `odontologia_consulta` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `is_deleted` tinyint(1) DEFAULT 0,
  `user_deleted` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_creator` varchar(255) DEFAULT NULL,
  `consulta_odontologia_id` char(36) NOT NULL,
  `relato` char(36) NOT NULL,
  `diagnostico` varchar(255) DEFAULT NULL,
  `examen_auxiliar` varchar(255) DEFAULT NULL,
  `tratamiento` varchar(255) DEFAULT NULL,
  `indicaciones` varchar(255) DEFAULT NULL,
  `comentarios` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_odontologia_consulta` (`consulta_odontologia_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- u103412080_obu_db.odontologia_examen definition

CREATE TABLE `odontologia_examen` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `is_deleted` tinyint(1) DEFAULT 0,
  `user_deleted` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_creator` varchar(255) DEFAULT NULL,
  `consulta_odontologia_id` char(36) NOT NULL,
  `odontograma_img` varchar(255) NOT NULL,
  `cpod` varchar(255) DEFAULT NULL,
  `observacion` varchar(255) DEFAULT NULL,
  `ihos` varchar(255) DEFAULT NULL,
  `comentarios` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_odontologia_examen` (`consulta_odontologia_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- u103412080_obu_db.odontologia_examen_bucal definition

CREATE TABLE `odontologia_examen_bucal` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `is_deleted` tinyint(1) DEFAULT 0,
  `user_deleted` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_creator` varchar(255) DEFAULT NULL,
  `consulta_odontologia_id` char(36) NOT NULL,
  `capacidad_masticatoria` varchar(255) DEFAULT NULL,
  `encias` varchar(255) DEFAULT NULL,
  `caries_dentales` varchar(255) DEFAULT NULL,
  `edentulismo_parcial_total` varchar(255) DEFAULT NULL,
  `portador_protesis_dental` varchar(255) DEFAULT NULL,
  `estado_higiene_bucal` varchar(255) DEFAULT NULL,
  `urgencia_tratamiento` varchar(255) DEFAULT NULL,
  `fluorizacion` varchar(255) DEFAULT NULL,
  `destartraje` varchar(255) DEFAULT NULL,
  `comentarios` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_odontologia_examen_bucal` (`consulta_odontologia_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- u103412080_obu_db.odontologia_procedimiento definition

CREATE TABLE `odontologia_procedimiento` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `is_deleted` tinyint(1) DEFAULT 0,
  `user_deleted` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_creator` varchar(255) DEFAULT NULL,
  `consulta_odontologia_id` char(36) NOT NULL,
  `tipo_procedimiento` varchar(255) DEFAULT NULL,
  `recibo` varchar(255) DEFAULT NULL,
  `costo` varchar(255) DEFAULT NULL,
  `fecha_pago` varchar(255) DEFAULT NULL,
  `pieza_dental` varchar(255) DEFAULT NULL,
  `comentarios` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_odontologia_procedimiento` (`consulta_odontologia_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- u103412080_obu_db.odontologia_revision_odontograma definition

CREATE TABLE `odontologia_revision_odontograma` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `is_deleted` tinyint(1) DEFAULT 0,
  `user_deleted` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_creator` varchar(255) DEFAULT NULL,
  `consulta_odontologia_id` char(36) NOT NULL,
  `caries` varchar(255) DEFAULT NULL,
  `erupcionado` varchar(255) DEFAULT NULL,
  `perdido` varchar(255) DEFAULT NULL,
  `costo` varchar(255) DEFAULT NULL,
  `fecha_pago` varchar(255) DEFAULT NULL,
  `cpod` varchar(255) DEFAULT NULL,
  `diagnostico` varchar(255) DEFAULT NULL,
  `mes` varchar(255) DEFAULT NULL,
  `comentarios` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_odontologia_revision_odontograma` (`consulta_odontologia_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- u103412080_obu_db.tipos_vacunas definition

CREATE TABLE `tipos_vacunas` (
  `id` char(36) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `is_deleted` tinyint(1) DEFAULT 0,
  `user_deleted` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_creator` varchar(255) DEFAULT NULL,
  `nombre` char(36) NOT NULL,
  `estado` tinyint(1) DEFAULT 1,
  `duracion_meses` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


INSERT INTO tipos_vacunas (id, created_at, updated_at, is_deleted, user_deleted, deleted_at, user_creator, nombre, estado, duracion_meses) VALUES('403dd6ce-ca78-4cae-a5d4-6f7321f70215', '2025-03-17 20:24:18', '2025-03-18 04:08:07', 0, NULL, NULL, NULL, 'Covid 19', 1, '1');
INSERT INTO tipos_vacunas (id, created_at, updated_at, is_deleted, user_deleted, deleted_at, user_creator, nombre, estado, duracion_meses) VALUES('64acd168-cc37-4a59-b091-0644c731fbc0', '2025-03-17 20:24:25', '2025-03-18 04:08:07', 0, NULL, NULL, NULL, 'Influenza', 1, '1');
INSERT INTO tipos_vacunas (id, created_at, updated_at, is_deleted, user_deleted, deleted_at, user_creator, nombre, estado, duracion_meses) VALUES('9b810217-b872-4d1d-9bb4-8c831e201579', '2025-03-17 20:24:27', '2025-03-18 04:08:07', 0, NULL, NULL, NULL, 'Sarampión y Rubeola', 1, '1');
INSERT INTO tipos_vacunas (id, created_at, updated_at, is_deleted, user_deleted, deleted_at, user_creator, nombre, estado, duracion_meses) VALUES('c998a109-c573-467b-aee1-0daeb591fc62', '2025-03-17 20:24:23', '2025-03-18 04:08:07', 0, NULL, NULL, NULL, 'Antihepatitis B', 1, '1');
INSERT INTO tipos_vacunas (id, created_at, updated_at, is_deleted, user_deleted, deleted_at, user_creator, nombre, estado, duracion_meses) VALUES('d5689b53-5217-4ffa-ab7c-c555c93c41bd', '2025-03-17 20:24:16', '2025-03-18 04:08:08', 0, NULL, NULL, NULL, 'Antitetánica', 1, '1');
INSERT INTO tipos_vacunas (id, created_at, updated_at, is_deleted, user_deleted, deleted_at, user_creator, nombre, estado, duracion_meses) VALUES('ddc3e450-5cef-471b-94a4-05e56aae836b', '2025-03-17 20:24:21', '2025-03-18 04:08:08', 0, NULL, NULL, NULL, 'Antiamarílica', 1, '1');

ALTER TABLE odontologia_consulta
    MODIFY COLUMN relato VARCHAR(500);

ALTER TABLE odontologia_consulta
    ADD COLUMN examen_clinico VARCHAR(255);

CREATE TABLE servicios_medicos_config (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    area ENUM('medicina', 'enfermeria', 'odontologia', 'psicologia') NOT NULL,
    tipo_servicio VARCHAR(100) NOT NULL,
    nombre_servicio VARCHAR(255) NOT NULL,
    requiere_pago TINYINT(1) NOT NULL DEFAULT 1 COMMENT '1=requiere pago, 0=gratuito',
    codigo_concepto INT(6) UNSIGNED NULL COMMENT 'Requerido solo si requiere_pago=1',
    obligatorio_estudiante TINYINT(1) NOT NULL DEFAULT 0 COMMENT '1=obligatorio, 0=opcional',
    obligatorio_docente TINYINT(1) NOT NULL DEFAULT 0 COMMENT '1=obligatorio, 0=opcional',
    obligatorio_administrativo TINYINT(1) NOT NULL DEFAULT 0 COMMENT '1=obligatorio, 0=opcional',
    estado TINYINT(1) NOT NULL DEFAULT 1 COMMENT '1=activo, 0=inactivo',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_area_tipo (area, tipo_servicio),
    INDEX idx_codigo_concepto (codigo_concepto)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO servicios_medicos_config
(area, tipo_servicio, nombre_servicio, requiere_pago, codigo_concepto, obligatorio_estudiante, obligatorio_docente, obligatorio_administrativo)
VALUES
    ('odontologia', 'procedimiento', 'profilaxis dental', 1, 971, 0, 0, 0),
    ('odontologia', 'procedimiento', 'destartraje dental', 1, 971, 0, 0, 0),
    ('odontologia', 'procedimiento', 'aplicación de flúor gel', 1, 967, 0, 0, 0),
    ('odontologia', 'procedimiento', 'exodoncia simple', 1, 970, 0, 0, 0),
    ('odontologia', 'procedimiento', 'exodoncia compleja', 1, 969, 0, 0, 0),
    ('odontologia', 'procedimiento', 'curetaje alveolar', 1, 968, 0, 0, 0),
    ('odontologia', 'procedimiento', 'resina simple', 1, 954, 0, 0, 0),
    ('odontologia', 'procedimiento', 'resina compuesta', 1, 953, 0, 0, 0),
    ('odontologia', 'procedimiento', 'apertura cameral', 0, NULL, 0, 0, 0),
    ('odontologia', 'procedimiento', 'endodoncia anterior', 1, 965, 0, 0, 0),
    ('odontologia', 'procedimiento', 'endodoncia posterior', 1, 966, 0, 0, 0),
    ('odontologia', 'procedimiento', 'radiografía', 1, 1526, 0, 0, 0),
    ('odontologia', 'procedimiento', 'cementado de corona', 1, 1527, 0, 0, 0),
    ('odontologia', 'procedimiento', 'prevención colutorio', 0, NULL, 0, 0, 0),
    ('odontologia', 'procedimiento', 'prevención IHO - TCA cepillado', 0, NULL, 0, 0, 0);


INSERT INTO response_messages
(code, message, message_type, http_status)
VALUES
    (250, 'No se encontró el recibo', 'WARNING', 500),
    (251, 'No se pudo consultar pagos', 'ERROR', 500),
    (252, 'El recibo ya fue utilizado para este tipo de servicio', 'WARNING', 500),
    (253, 'EL recibo no corresponde al tipo de servicio', 'WARNING', 500),
    (254, 'El recibo es válido', 'SUCCESS', 200),
    (255, 'El tipo de servicio no tiene concepto pago asociado, consulte con el administrador', 'WARNING', 500),
    (256, 'El tipo de servicio no acepta pago', 'WARNING', 500),
    (257, 'Recibo válido', 'SUCESSS', 500);

alter table enfermeria_tratamiento_medicamentoso
    add column area_solicitante varchar(255) not null;

alter table enfermeria_tratamiento_medicamentoso
    add column especialista_solicitante varchar(255) not null;

alter table enfermeria_procedimientos_realizados
    add column area_solicitante varchar(255) not null;

alter table enfermeria_procedimientos_realizados
    add column especialista_solicitante varchar(255) not null;

delete from response_messages where code = 257;

INSERT INTO servicios_medicos_config
(area, tipo_servicio, nombre_servicio, requiere_pago, codigo_concepto, obligatorio_estudiante, obligatorio_docente, obligatorio_administrativo)
VALUES ('odontologia', 'procedimiento', 'otros', 1, 1313, 0, 0, 0);

/*
INSERT INTO servicios_medicos_config
(area, tipo_servicio, nombre_servicio, requiere_pago, codigo_concepto, obligatorio_estudiante, obligatorio_docente, obligatorio_administrativo)
VALUES
    ('enfermeria', 'procedimientos', 'INYECTABLES', 0, null, 0, 0, 0),
    ('enfermeria', 'procedimientos', 'CONTROL DE FUNCIONES VITALES', 0, null, 0, 0, 0),
    ('enfermeria', 'procedimientos', 'LAVADO DE OÍDO', 1, 710, 0, 0, 0),
    ('enfermeria', 'procedimientos', 'BAÑO RAYOS INFRAROJOS', 1, 955, 0, 0, 0),
    ('enfermeria', 'procedimientos', 'EXTRACCIÓN DE UÑERO', 1, 636, 0, 0, 0),
    ('enfermeria', 'procedimientos', 'VENOCLISIS', 0, null, 0, 0, 0),
    ('odontologia', 'procedimiento', 'OTROS', 1, 1313, 0, 0, 0);
*/