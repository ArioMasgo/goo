CREATE TABLE `registro_toxicologico` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `alumno_id` bigint(20) UNSIGNED NOT NULL,
  `convocatoria_id` bigint(20) UNSIGNED NOT NULL,
  `estado` enum('verificado','observado','pendiente') NOT NULL,
  `comentario` text DEFAULT NULL,
  `id_usuario` bigint(20) UNSIGNED DEFAULT NULL,
  `fecha_creacion` datetime NOT NULL DEFAULT current_timestamp(),
  `fecha_actualizacion` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `fk_rt_convocatoria_id` (`convocatoria_id`),
  KEY `fk_rt_id_usuario` (`id_usuario`),
  KEY `fk_rt_alumno_id` (`alumno_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

ALTER TABLE `registro_toxicologico`
  ADD CONSTRAINT `fk_rt_alumno_id` FOREIGN KEY (`alumno_id`) REFERENCES `alumnos` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `fk_rt_convocatoria_id` FOREIGN KEY (`convocatoria_id`) REFERENCES `convocatorias` (`id`),
  ADD CONSTRAINT `fk_rt_id_usuario` FOREIGN KEY (`id_usuario`) REFERENCES `users` (`id`);
