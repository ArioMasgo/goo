-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Servidor: 127.0.0.1:3306
-- Tiempo de generación: 18-03-2025 a las 15:50:18
-- Versión del servidor: 10.11.10-MariaDB
-- Versión de PHP: 7.2.34

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de datos: `u103412080_obu_db`
--

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `citas`
--

CREATE TABLE `citas` (
  `id` int(11) NOT NULL,
  `dni` varchar(20) NOT NULL,
  `nombre` varchar(100) NOT NULL,
  `apellido` varchar(100) NOT NULL,
  `facultad` varchar(100) NOT NULL,
  `fecha_inicio` datetime NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `fecha_fin` datetime NOT NULL,
  `estado` enum('Completado','Pendiente','Cancelado') NOT NULL DEFAULT 'Pendiente'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `diagnosticos`
--

CREATE TABLE `diagnosticos` (
  `id` int(11) NOT NULL,
  `codigo` varchar(50) NOT NULL,
  `nombre` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `encuestas`
--

CREATE TABLE `encuestas` (
  `id_encuesta` int(11) NOT NULL,
  `nombre_encuesta` varchar(255) NOT NULL,
  `descripcion` text DEFAULT NULL,
  `estado` enum('activa','inactiva') NOT NULL DEFAULT 'activa',
  `fecha_inicio` date DEFAULT NULL,
  `fecha_fin` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `historial`
--

CREATE TABLE `historial` (
  `id_historial` int(11) NOT NULL,
  `id_participante` int(11) NOT NULL,
  `id_encuesta` int(11) DEFAULT NULL,
  `fecha_respuesta` datetime NOT NULL DEFAULT current_timestamp(),
  `diagnostico` enum('Estable','Moderado','Grave','No Aplica') NOT NULL,
  `notes` text DEFAULT NULL,
  `num_telefono` varchar(20) DEFAULT NULL,
  `con_quienes_vive_actualmente` varchar(50) DEFAULT NULL,
  `estado_evaluacion` enum('Aprobado','Pendiente','Observada','No Aplica') NOT NULL DEFAULT 'Pendiente',
  `semestre_cursa` varchar(10) DEFAULT NULL,
  `direccion` varchar(255) NOT NULL,
  `quien_financia_carrera` varchar(255) NOT NULL,
  `motivo_consulta` varchar(255) NOT NULL,
  `situacion_actual` varchar(255) NOT NULL,
  `otros_procedimientos` varchar(200) NOT NULL,
  `created_date` datetime DEFAULT current_timestamp(),
  `es_srq` tinyint(1) DEFAULT 0,
  `diagnostico_id` int(11) DEFAULT NULL,
  `key_url` text DEFAULT NULL,
  `notas_atencion` text DEFAULT NULL,
  `instrumentos_utilizados` varchar(255) DEFAULT NULL,
  `resultados_obtenidos` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `participantes`
--

CREATE TABLE `participantes` (
  `id_participante` int(11) NOT NULL,
  `tipo_participante` enum('Alumno','Administrativo','Externo','Docente') NOT NULL,
  `nombre` varchar(100) NOT NULL,
  `apellido` varchar(100) NOT NULL,
  `dni` varchar(20) DEFAULT NULL,
  `estado` enum('Activo','Inactivo') NOT NULL DEFAULT 'Activo',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `colegio_procedencia` varchar(150) DEFAULT NULL,
  `anio_ingreso` year(4) DEFAULT NULL,
  `escuela` varchar(200) NOT NULL,
  `codigo_estudiante` varchar(100) NOT NULL,
  `fecha_nacimiento` varchar(30) NOT NULL,
  `edad` int(20) NOT NULL,
  `lugar_nacimiento` varchar(200) NOT NULL,
  `modalidad_ingreso` varchar(200) NOT NULL,
  `numero_atencion` int(11) NOT NULL,
  `sexo` varchar(20) NOT NULL,
  `num_telefono` varchar(20) DEFAULT NULL,
  `estado_evaluacion` varchar(20) NOT NULL,
  `diagnostico` varchar(20) NOT NULL,
  `con_quienes_vive_actualmente` varchar(255) DEFAULT NULL,
  `semestre_cursa` varchar(10) DEFAULT NULL,
  `direccion` varchar(255) DEFAULT NULL,
  `profesion` varchar(200) DEFAULT NULL,
  `estado_civil` varchar(200) DEFAULT NULL,
  `labora_en_unas` tinyint(1) DEFAULT 0,
  `grado_instruccion` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `questions`
--

CREATE TABLE `questions` (
  `id` int(11) NOT NULL,
  `texto_pregunta` text NOT NULL,
  `is_mandatory` tinyint(1) NOT NULL,
  `order` int(11) NOT NULL,
  `type` enum('open','multiple_choice','rating_scale','yes_no') NOT NULL DEFAULT 'open'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Volcado de datos para la tabla `questions`
--

INSERT INTO `questions` (`id`, `texto_pregunta`, `is_mandatory`, `order`, `type`) VALUES
(1, '¿Tiene frecuentes dolores de cabeza?', 1, 1, 'yes_no'),
(2, '¿Tiene mal apetito frecuentemente?', 1, 2, 'yes_no'),
(3, '¿Duerme mal?', 1, 3, 'yes_no'),
(4, '¿Se asusta con facilidad?', 1, 4, 'yes_no'),
(5, '¿Sufre de temblor de manos frecuentemente?', 1, 6, 'yes_no'),
(6, '¿Se siente nervioso o tenso con frecuencia?', 1, 7, 'yes_no'),
(7, '¿Sufre de mala digestión?', 1, 8, 'yes_no'),
(8, '¿Tiene dificultades para pensar con claridad?', 1, 9, 'yes_no'),
(9, '¿Se siente triste?', 1, 10, 'yes_no'),
(10, '¿Llora usted con mucha frecuencia?', 1, 11, 'yes_no'),
(11, '¿Tiene dificultad en disfrutar sus actividades diarias?', 1, 12, 'yes_no'),
(12, '¿Tiene dificultad para tomar decisiones?', 1, 13, 'yes_no'),
(13, '¿Tiene dificultad en hacer su trabajo y/o estudios? (¿Sufre usted con su trabajo y/o estudios?)', 1, 14, 'yes_no'),
(14, '¿Es incapaz de desempeñar un papel útil en su vida?', 1, 15, 'yes_no'),
(15, '¿Ha perdido interés en las cosas que usualmente hacía?', 1, 16, 'yes_no'),
(16, '¿Siente que usted es una persona inútil?', 1, 17, 'yes_no'),
(17, '¿Ha tenido la idea de acabar con su vida?', 1, 18, 'yes_no'),
(18, '¿Se siente cansado todo el tiempo?', 1, 19, 'yes_no'),
(19, '¿Tiene sensaciones desagradables en su estómago?', 1, 20, 'yes_no'),
(20, '¿Se cansa con facilidad?', 1, 21, 'yes_no'),
(21, '¿Siente usted que alguien ha tratado de herirlo en alguna forma? Ejm: pensar que alguien conspira contra usted, que alguien intenta dañarle, etc.', 1, 22, 'yes_no'),
(22, '¿Usted se considera una persona mucho más importante que los demás?', 1, 23, 'yes_no'),
(23, '¿Ha notado interferencias o algo raro en su pensamiento? Ejm: oír voces, ver cosas que otras personas no pueden ver ni oír, etc.', 1, 24, 'yes_no'),
(24, '¿Oye voces sin saber de dónde vienen o que otras personas no pueden oír?', 1, 25, 'yes_no'),
(25, '¿Ha tenido convulsiones, ataques o caídas al suelo, con movimientos de brazos y piernas; con mordedura de la lengua o pérdida del conocimiento?', 1, 26, 'yes_no'),
(26, '¿Alguna vez le ha parecido a su familia, sus amigos, su médico o a su sacerdote que usted estaba bebiendo demasiado licor?', 1, 27, 'yes_no'),
(27, '¿Alguna vez ha querido dejar de beber, pero no ha podido?', 1, 28, 'yes_no'),
(28, '¿Ha tenido alguna vez dificultades en el trabajo (o a causa de la bebida, como beber en el trabajo o en el colegio, o faltar a ellos)?', 1, 29, 'yes_no'),
(29, '¿Ha estado en riñas o lo han detenido por estar borracho?', 1, 30, 'yes_no'),
(30, '¿Le ha parecido alguna vez que usted bebía demasiado?', 1, 31, 'yes_no');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `question_options`
--

CREATE TABLE `question_options` (
  `id` int(11) NOT NULL,
  `question_id` int(11) NOT NULL,
  `option_text` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Volcado de datos para la tabla `question_options`
--

INSERT INTO `question_options` (`id`, `question_id`, `option_text`) VALUES
(5, 1, 'Sí'),
(6, 1, 'No'),
(7, 2, 'Sí'),
(8, 2, 'No'),
(9, 3, 'Sí'),
(10, 3, 'No'),
(11, 4, 'Sí'),
(12, 4, 'No'),
(13, 5, 'Sí'),
(14, 5, 'No'),
(15, 6, 'Sí'),
(16, 6, 'No'),
(17, 7, 'Sí'),
(18, 7, 'No'),
(19, 8, 'Sí'),
(20, 8, 'No'),
(21, 9, 'Sí'),
(22, 9, 'No'),
(23, 10, 'Sí'),
(24, 10, 'No'),
(25, 11, 'Sí'),
(26, 11, 'No'),
(27, 12, 'Sí'),
(28, 12, 'No'),
(29, 13, 'Sí'),
(30, 13, 'No'),
(31, 14, 'Sí'),
(32, 14, 'No'),
(33, 15, 'Sí'),
(34, 15, 'No'),
(35, 16, 'Sí'),
(36, 16, 'No'),
(37, 17, 'Sí'),
(38, 17, 'No'),
(39, 18, 'Sí'),
(40, 18, 'No'),
(41, 19, 'Sí'),
(42, 19, 'No'),
(43, 20, 'Sí'),
(44, 20, 'No'),
(45, 21, 'Sí'),
(46, 21, 'No'),
(47, 22, 'Sí'),
(48, 22, 'No'),
(49, 23, 'Sí'),
(50, 23, 'No'),
(51, 24, 'Sí'),
(52, 24, 'No'),
(53, 25, 'Sí'),
(54, 25, 'No'),
(55, 26, 'Sí'),
(56, 26, 'No'),
(57, 27, 'Sí'),
(58, 27, 'No'),
(59, 28, 'Sí'),
(60, 28, 'No'),
(61, 29, 'Sí'),
(62, 29, 'No'),
(63, 30, 'Sí'),
(64, 30, 'No');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `respuestas`
--

CREATE TABLE `respuestas` (
  `id_respuesta` int(11) NOT NULL,
  `id_participante` int(11) NOT NULL,
  `id_historial` int(11) NOT NULL,
  `id_pregunta` int(11) NOT NULL,
  `respuesta` text NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `numero_atencion` int(11) NOT NULL,
  `id_encuesta` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Índices para tablas volcadas
--

--
-- Indices de la tabla `citas`
--
ALTER TABLE `citas`
  ADD PRIMARY KEY (`id`);

--
-- Indices de la tabla `diagnosticos`
--
ALTER TABLE `diagnosticos`
  ADD PRIMARY KEY (`id`);

--
-- Indices de la tabla `encuestas`
--
ALTER TABLE `encuestas`
  ADD PRIMARY KEY (`id_encuesta`);

--
-- Indices de la tabla `historial`
--
ALTER TABLE `historial`
  ADD PRIMARY KEY (`id_historial`),
  ADD KEY `id_participante` (`id_participante`),
  ADD KEY `id_encuesta` (`id_encuesta`),
  ADD KEY `fk_diagnostico` (`diagnostico_id`);

--
-- Indices de la tabla `participantes`
--
ALTER TABLE `participantes`
  ADD PRIMARY KEY (`id_participante`),
  ADD UNIQUE KEY `dni` (`dni`);

--
-- Indices de la tabla `questions`
--
ALTER TABLE `questions`
  ADD PRIMARY KEY (`id`);

--
-- Indices de la tabla `question_options`
--
ALTER TABLE `question_options`
  ADD PRIMARY KEY (`id`),
  ADD KEY `question_id` (`question_id`);

--
-- Indices de la tabla `respuestas`
--
ALTER TABLE `respuestas`
  ADD PRIMARY KEY (`id_respuesta`),
  ADD KEY `id_participante` (`id_participante`),
  ADD KEY `id_pregunta` (`id_pregunta`),
  ADD KEY `id_historial` (`id_historial`),
  ADD KEY `id_encuesta` (`id_encuesta`);

--
-- AUTO_INCREMENT de las tablas volcadas
--

--
-- AUTO_INCREMENT de la tabla `citas`
--
ALTER TABLE `citas`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de la tabla `diagnosticos`
--
ALTER TABLE `diagnosticos`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de la tabla `encuestas`
--
ALTER TABLE `encuestas`
  MODIFY `id_encuesta` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de la tabla `historial`
--
ALTER TABLE `historial`
  MODIFY `id_historial` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de la tabla `participantes`
--
ALTER TABLE `participantes`
  MODIFY `id_participante` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de la tabla `questions`
--
ALTER TABLE `questions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=34;

--
-- AUTO_INCREMENT de la tabla `question_options`
--
ALTER TABLE `question_options`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=65;

--
-- AUTO_INCREMENT de la tabla `respuestas`
--
ALTER TABLE `respuestas`
  MODIFY `id_respuesta` int(11) NOT NULL AUTO_INCREMENT;

--
-- Restricciones para tablas volcadas
--

--
-- Filtros para la tabla `historial`
--
ALTER TABLE `historial`
  ADD CONSTRAINT `fk_diagnostico` FOREIGN KEY (`diagnostico_id`) REFERENCES `diagnosticos` (`id`),
  ADD CONSTRAINT `historial_ibfk_1` FOREIGN KEY (`id_participante`) REFERENCES `participantes` (`id_participante`) ON DELETE CASCADE,
  ADD CONSTRAINT `historial_ibfk_2` FOREIGN KEY (`id_encuesta`) REFERENCES `encuestas` (`id_encuesta`) ON DELETE CASCADE;

--
-- Filtros para la tabla `question_options`
--
ALTER TABLE `question_options`
  ADD CONSTRAINT `question_options_ibfk_1` FOREIGN KEY (`question_id`) REFERENCES `questions` (`id`) ON DELETE CASCADE;

--
-- Filtros para la tabla `respuestas`
--
ALTER TABLE `respuestas`
  ADD CONSTRAINT `respuestas_ibfk_1` FOREIGN KEY (`id_participante`) REFERENCES `participantes` (`id_participante`) ON DELETE CASCADE,
  ADD CONSTRAINT `respuestas_ibfk_3` FOREIGN KEY (`id_pregunta`) REFERENCES `questions` (`id`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
