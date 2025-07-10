CREATE TABLE `Visitas_Domiciliarias` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `alumno_id` bigint(20) UNSIGNED NOT NULL,
  `estado` enum('verificado','observado','pendiente') NOT NULL,
  `comentario` text DEFAULT NULL,
  `imagen_url` text DEFAULT NULL,
  `id_usuario` bigint(20) UNSIGNED DEFAULT NULL,
  `fecha_creacion` datetime NOT NULL DEFAULT current_timestamp(),
  `fecha_actualizacion` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `fk_rtvd_id_usuario` (`id_usuario`),
  KEY `fk_rtvd_alumno_id` (`alumno_id`),
  CONSTRAINT `fk_rtvd_alumno_id` FOREIGN KEY (`alumno_id`) REFERENCES `alumnos` (`id`),
  CONSTRAINT `fk_rtvd_id_usuario` FOREIGN KEY (`id_usuario`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;


